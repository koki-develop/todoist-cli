package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure CLI settings",
	Long:  "Configure CLI settings.",
	RunE: func(cmd *cobra.Command, args []string) error {
		tkn := flagAPIToken.Get(cmd, true)
		if tkn == nil {
			var t string
			fmt.Print("API Token: ")
			fmt.Scanf("%s", &t)
			if t != "" {
				tkn = &t
			}
		}
		if tkn == nil {
			fmt.Println("Canceled.")
			return nil
		}

		j, err := json.Marshal(&Config{APIToken: *tkn})
		if err != nil {
			return err
		}

		cfgdir, err := configDir()
		if err != nil {
			return err
		}
		if err := os.MkdirAll(cfgdir, os.ModePerm); err != nil {
			return err
		}
		cfgfile := configFilename(cfgdir)
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
