package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
)

type Config struct {
	APIToken string `json:"api_token"`
}

var configureCmd = &cobra.Command{
	Use: "configure",
	RunE: func(cmd *cobra.Command, args []string) error {
		if apiToken == "" {
			fmt.Print("API Token: ")
			fmt.Scanf("%s", &apiToken)
		}

		cfg := &Config{APIToken: apiToken}
		j, err := json.Marshal(cfg)
		if err != nil {
			return err
		}

		hmd, err := os.UserHomeDir()
		if err != nil {
			return err
		}

		cfgdir := path.Join(hmd, ".todoist")
		if err := os.MkdirAll(cfgdir, os.ModePerm); err != nil {
			return err
		}
		cfgfile := path.Join(cfgdir, "config.json")
		f, err := os.Create(cfgfile)
		if err != nil {
			return err
		}
		defer f.Close()

		if _, err := f.Write(j); err != nil {
			return err
		}

		fmt.Println("Configured!")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
}
