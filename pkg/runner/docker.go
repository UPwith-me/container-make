package runner

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/container-make/cm/pkg/config"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/docker/go-connections/nat"
	"golang.org/x/term"
)

type Runner struct {
	Client *client.Client
	Config *config.DevContainerConfig
}

func NewRunner(cfg *config.DevContainerConfig) (*Runner, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	return &Runner{Client: cli, Config: cfg}, nil
}

func (r *Runner) Run(ctx context.Context, command []string) error {
	var imageTag string
	var err error

	// 1. Determine Image Source
	if r.Config.Build != nil {
		// Build from Dockerfile
		imageTag, err = r.Build(ctx)
		if err != nil {
			return fmt.Errorf("build failed: %w", err)
		}
		r.Config.Image = imageTag // Update config to use built image
	} else if r.Config.Image != "" {
		// Use existing image
		imageTag = r.Config.Image
		_, _, err := r.Client.ImageInspectWithRaw(ctx, imageTag)
		if err != nil {
			if client.IsErrNotFound(err) {
				fmt.Printf("Image %s not found locally, pulling...\n", imageTag)
				reader, err := r.Client.ImagePull(ctx, imageTag, image.PullOptions{})
				if err != nil {
					return fmt.Errorf("failed to pull image: %w", err)
				}
				io.Copy(io.Discard, reader)
				reader.Close()
				fmt.Println("Image pulled.")
			} else {
				return fmt.Errorf("failed to inspect image: %w", err)
			}
		} else {
			fmt.Printf("Image %s found locally, skipping pull.\n", imageTag)
		}
	} else {
		return fmt.Errorf("no image or build configuration found")
	}

	// 2. Create Container
	fmt.Println("Creating container...")

	// Check if we are in a terminal
	isTerminal := term.IsTerminal(int(os.Stdin.Fd()))

	// Basic HostConfig
	hostConfig := &container.HostConfig{
		AutoRemove: true,             // --rm
		Init:       &[]bool{true}[0], // --init
		Binds:      r.Config.Mounts,
	}

	// Port Forwarding
	exposedPorts := nat.PortSet{}
	portBindings := nat.PortMap{}

	for _, p := range r.Config.ForwardPorts {
		var portStr string
		switch v := p.(type) {
		case float64: // JSON numbers are floats
			portStr = fmt.Sprintf("%d", int(v))
		case int:
			portStr = fmt.Sprintf("%d", v)
		case string:
			portStr = v
		default:
			fmt.Printf("Warning: invalid port format: %v\n", p)
			continue
		}

		// Assume TCP for now
		port := nat.Port(portStr + "/tcp")
		exposedPorts[port] = struct{}{}
		portBindings[port] = []nat.PortBinding{
			{
				HostIP:   "127.0.0.1", // Bind to localhost
				HostPort: portStr,     // Map 1:1
			},
		}
		fmt.Printf("Forwarding port %s\n", portStr)
	}

	hostConfig.PortBindings = portBindings

	// Entrypoint setup
	// We inject a script to handle UID mapping
	entrypointPath := "/tmp/cm-entrypoint.sh"

	// ContainerConfig
	containerConfig := &container.Config{
		Image:        r.Config.Image,
		Cmd:          command,
		Env:          mapToEnvList(r.Config.ContainerEnv),
		User:         "root", // Always start as root to allow user creation, script will drop privileges
		Tty:          isTerminal,
		OpenStdin:    true,
		Entrypoint:   []string{"/bin/sh", entrypointPath},
		ExposedPorts: exposedPorts,
	}

	resp, err := r.Client.ContainerCreate(ctx, containerConfig, hostConfig, nil, nil, "")
	if err != nil {
		return fmt.Errorf("failed to create container: %w", err)
	}
	fmt.Printf("Container created: %s\n", resp.ID)

	// 2.5 Inject Entrypoint Script
	if err := r.copyEntrypointToContainer(ctx, resp.ID, entrypointPath); err != nil {
		// Clean up
		r.Client.ContainerRemove(ctx, resp.ID, container.RemoveOptions{Force: true})
		return fmt.Errorf("failed to inject entrypoint: %w", err)
	}

	// 3. Start Container
	fmt.Println("Starting container...")
	if err := r.Client.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		return fmt.Errorf("failed to start container: %w", err)
	}

	// 3.1 Lifecycle Hooks: PostCreateCommand & PostStartCommand
	// Since we are ephemeral, we run both here.
	if err := r.executeLifecycleHook(ctx, resp.ID, "postCreateCommand", r.Config.PostCreateCommand); err != nil {
		fmt.Printf("Warning: postCreateCommand failed: %v\n", err)
	}
	if err := r.executeLifecycleHook(ctx, resp.ID, "postStartCommand", r.Config.PostStartCommand); err != nil {
		fmt.Printf("Warning: postStartCommand failed: %v\n", err)
	}

	// 3.2 Features Warning
	if len(r.Config.Features) > 0 {
		fmt.Println("Warning: 'features' are detected in devcontainer.json but are not yet supported by Container-Make. They will be ignored.")
	}

	// 4. Handle Signals & TTY
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Resize TTY if applicable
	if isTerminal {
		// Set initial size
		width, height, _ := term.GetSize(int(os.Stdin.Fd()))
		r.Client.ContainerResize(ctx, resp.ID, container.ResizeOptions{
			Height: uint(height),
			Width:  uint(width),
		})

		// Put terminal in raw mode
		oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
		if err != nil {
			fmt.Printf("Warning: failed to set raw mode: %v\n", err)
		} else {
			defer term.Restore(int(os.Stdin.Fd()), oldState)
		}
	}

	go func() {
		<-sigChan
		// Restore terminal before printing (if in raw mode)
		// Note: defer handles restoration on return, but here we might want to ensure clean output
		// For now, just stop container.
		timeout := 10 // seconds
		r.Client.ContainerStop(ctx, resp.ID, container.StopOptions{Timeout: &timeout})
	}()

	// 5. Attach / Logs
	attachResp, err := r.Client.ContainerAttach(ctx, resp.ID, container.AttachOptions{
		Stream: true,
		Stdin:  true,
		Stdout: true,
		Stderr: true,
		Logs:   true,
	})
	if err != nil {
		return fmt.Errorf("failed to attach: %w", err)
	}
	defer attachResp.Close()

	// Stream IO
	go io.Copy(attachResp.Conn, os.Stdin)

	if isTerminal {
		// In TTY mode, stdout and stderr are merged
		_, err = io.Copy(os.Stdout, attachResp.Reader)
	} else {
		// In non-TTY mode, use StdCopy to demultiplex
		_, err = stdcopy.StdCopy(os.Stdout, os.Stderr, attachResp.Reader)
	}

	// 6. Wait for exit
	statusCh, errCh := r.Client.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			return fmt.Errorf("error waiting for container: %w", err)
		}
	case <-statusCh:
	}

	return nil
}

func (r *Runner) copyEntrypointToContainer(ctx context.Context, containerID, path string) error {
	script := GetEntrypoint()

	// Create a tar archive in memory
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)

	hdr := &tar.Header{
		Name: "cm-entrypoint.sh", // Filename in tar, will be extracted to path's directory? No, CopyToContainer extracts to the destination path which must be a directory.
		// Wait, CopyToContainer path is the destination directory.
		// So if I want /tmp/cm-entrypoint.sh, I should copy to /tmp.
		Mode: 0755,
		Size: int64(len(script)),
	}

	if err := tw.WriteHeader(hdr); err != nil {
		return err
	}
	if _, err := tw.Write([]byte(script)); err != nil {
		return err
	}
	if err := tw.Close(); err != nil {
		return err
	}

	// Copy to container
	// Path must be a directory
	return r.Client.CopyToContainer(ctx, containerID, "/tmp", buf, container.CopyToContainerOptions{})
}

func (r *Runner) Build(ctx context.Context) (string, error) {
	if r.Config.Build == nil {
		return "", fmt.Errorf("no build configuration")
	}

	// Determine context and dockerfile
	buildContext := r.Config.Build.Context
	if buildContext == "" {
		buildContext = "."
	}
	dockerfile := r.Config.Build.Dockerfile
	if dockerfile == "" {
		dockerfile = "Dockerfile"
	}

	// Generate a tag based on the config hash or project name
	// For simplicity, let's use "cm-dev-env" for now, or maybe hash the path
	tag := "cm-dev-env:latest"

	fmt.Printf("Building image %s from %s...\n", tag, dockerfile)

	// Construct docker build command
	args := []string{"build", "-t", tag, "-f", dockerfile}

	// Add build args
	for k, v := range r.Config.Build.Args {
		args = append(args, "--build-arg", fmt.Sprintf("%s=%s", k, v))
	}

	args = append(args, buildContext)

	cmd := exec.CommandContext(ctx, "docker", args...)
	cmd.Env = append(os.Environ(), "DOCKER_BUILDKIT=1")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return "", err
	}

	return tag, nil
}

func (r *Runner) executeLifecycleHook(ctx context.Context, containerID, name string, cmd interface{}) error {
	if cmd == nil {
		return nil
	}

	var commands []string
	switch v := cmd.(type) {
	case string:
		commands = []string{v}
	case []interface{}:
		for _, item := range v {
			if s, ok := item.(string); ok {
				commands = append(commands, s)
			}
		}
	}

	if len(commands) == 0 {
		return nil
	}

	fmt.Printf("Executing %s...\n", name)
	for _, c := range commands {
		// Create Exec
		execConfig := container.ExecOptions{
			Cmd:          []string{"/bin/sh", "-c", c},
			AttachStdout: true,
			AttachStderr: true,
		}
		execIDResp, err := r.Client.ContainerExecCreate(ctx, containerID, execConfig)
		if err != nil {
			return fmt.Errorf("failed to create exec for %s: %w", name, err)
		}

		// Attach to Exec
		resp, err := r.Client.ContainerExecAttach(ctx, execIDResp.ID, container.ExecStartOptions{})
		if err != nil {
			return fmt.Errorf("failed to attach exec for %s: %w", name, err)
		}
		defer resp.Close()

		// Stream output
		// We just dump to stdout/stderr for now
		stdcopy.StdCopy(os.Stdout, os.Stderr, resp.Reader)
	}

	return nil
}

func mapToEnvList(m map[string]string) []string {
	var env []string
	for k, v := range m {
		env = append(env, fmt.Sprintf("%s=%s", k, v))
	}
	return env
}
