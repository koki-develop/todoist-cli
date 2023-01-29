package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "todoist",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		if err == ErrLoadConfig {
			fmt.Println("Run `todoist configure` to reconfigure.")
		}
		os.Exit(1)
	}
}
