package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/tailscale/hujson"
)

// DevContainerConfig represents the structure of devcontainer.json
// Reference: https://containers.dev/implementors/json_reference/
type DevContainerConfig struct {
	Image             string                 `json:"image,omitempty"`
	Build             *BuildConfig           `json:"build,omitempty"`
	RunArgs           []string               `json:"runArgs,omitempty"`
	Mounts            []string               `json:"mounts,omitempty"`
	ContainerEnv      map[string]string      `json:"containerEnv,omitempty"`
	RemoteEnv         map[string]string      `json:"remoteEnv,omitempty"`
	PostCreateCommand interface{}            `json:"postCreateCommand,omitempty"` // string or []string
	PostStartCommand  interface{}            `json:"postStartCommand,omitempty"`  // string or []string
	PostAttachCommand interface{}            `json:"postAttachCommand,omitempty"` // string or []string
	Features          map[string]interface{} `json:"features,omitempty"`
	ForwardPorts      []interface{}          `json:"forwardPorts,omitempty"` // number or string (we'll parse to int later if needed, or just handle int/string)
	User              string                 `json:"user,omitempty"`
	WorkspaceMount    string                 `json:"workspaceMount,omitempty"`
	WorkspaceFolder   string                 `json:"workspaceFolder,omitempty"`
}

type BuildConfig struct {
	Dockerfile string            `json:"dockerfile,omitempty"`
	Context    string            `json:"context,omitempty"`
	Args       map[string]string `json:"args,omitempty"`
}

// ParseConfig reads and parses a devcontainer.json file
func ParseConfig(path string) (*DevContainerConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	// Use hujson to standardize the JSON (remove comments, trailing commas)
	stdData, err := hujson.Standardize(data)
	if err != nil {
		return nil, fmt.Errorf("failed to standardize jsonc: %w", err)
	}

	var config DevContainerConfig
	if err := json.Unmarshal(stdData, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}
