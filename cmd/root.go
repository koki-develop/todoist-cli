package cmd

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/spf13/cobra"
)

var (
	version string
)

var rootCmd = &cobra.Command{
	Use:  "todoist-cli",
	Long: "CLI Client for Todoist.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		if err == ErrLoadConfig {
			fmt.Println("Run `todoist-cli configure` to reconfigure.")
		}
		os.Exit(1)
	}
}

func init() {
	if version == "" {
		if info, ok := debug.ReadBuildInfo(); ok {
			version = info.Main.Version
		}
	}

	rootCmd.Version = version
}
