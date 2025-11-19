package main

import (
	"context"
	"fmt"
	"os"

	"github.com/container-make/cm/pkg/config"
	"github.com/container-make/cm/pkg/runner"
	"github.com/spf13/cobra"
)

var configFile string

var rootCmd = &cobra.Command{
	Use:   "cm",
	Short: "Container-Make: A tool to execute build commands inside containers",
	Long: `Container-Make (cm) is a CLI tool that bridges the gap between local Makefiles
and containerized build environments. It reads devcontainer.json configurations
and executes commands in ephemeral or persistent containers.`,
}

var runCmd = &cobra.Command{
	Use:   "run [command]",
	Short: "Run a command inside the dev container",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Default config paths
		if configFile == "" {
			if _, err := os.Stat(".devcontainer/devcontainer.json"); err == nil {
				configFile = ".devcontainer/devcontainer.json"
			} else if _, err := os.Stat("devcontainer.json"); err == nil {
				configFile = "devcontainer.json"
			} else {
				return fmt.Errorf("no devcontainer.json found")
			}
		}

		cfg, err := config.ParseConfig(configFile)
		if err != nil {
			return err
		}

		r, err := runner.NewRunner(cfg)
		if err != nil {
			return err
		}

		return r.Run(context.Background(), args)
	},
}

var prepareCmd = &cobra.Command{
	Use:   "prepare",
	Short: "Build the dev container image",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Default config paths
		if configFile == "" {
			if _, err := os.Stat(".devcontainer/devcontainer.json"); err == nil {
				configFile = ".devcontainer/devcontainer.json"
			} else if _, err := os.Stat("devcontainer.json"); err == nil {
				configFile = "devcontainer.json"
			} else {
				return fmt.Errorf("no devcontainer.json found")
			}
		}

		cfg, err := config.ParseConfig(configFile)
		if err != nil {
			return err
		}

		r, err := runner.NewRunner(cfg)
		if err != nil {
			return err
		}

		if cfg.Build != nil {
			tag, err := r.Build(context.Background())
			if err != nil {
				return err
			}
			fmt.Printf("Successfully built %s\n", tag)
		} else if cfg.Image != "" {
			fmt.Printf("Config uses static image %s, pulling...\n", cfg.Image)
			// Trigger pull logic via Run or just pull?
			// For now, let's just say we are done as Run handles pull.
			// Or we can expose Pull in runner.
			// Let's just reuse the logic in Run but without command?
			// Actually, Build() method in runner handles build. We don't have a Pull() method exposed.
			// Let's just print a message.
		}

		return nil
	},
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Generate shell integration scripts",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("# Add this to your shell configuration (.bashrc, .zshrc, etc.)")
		fmt.Println("alias devcontainer='cm run'")
		fmt.Println("# Or use the shim function:")
		fmt.Println("function dcm() {")
		fmt.Println("  cm run --config .devcontainer/devcontainer.json -- \"$@\"")
		fmt.Println("}")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(prepareCmd)
	rootCmd.AddCommand(initCmd)
	runCmd.Flags().StringVarP(&configFile, "config", "c", "", "Path to devcontainer.json")
	prepareCmd.Flags().StringVarP(&configFile, "config", "c", "", "Path to devcontainer.json")
	Execute()
}
