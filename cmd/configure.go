package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/koki-develop/todoist-cli/pkg/renderer"
	"github.com/spf13/cobra"
)

func read(txt string) string {
	fmt.Printf("%s: ", txt)
	var in string
	fmt.Scanf("%s", &in)
	return in
}

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure CLI settings",
	Long:  "Configure CLI settings.",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := &Config{}
		if tkn := flagAPIToken.Get(cmd, true); tkn != nil {
			cfg.APIToken = *tkn
		} else {
			tkn := read("API Token")
			cfg.APIToken = tkn
		}

		if f := flagFormat.Get(cmd, true); f != nil {
			cfg.Format = *f
		} else {
			f := read(fmt.Sprintf("Default Output Format (optional) (%s)", strings.Join(renderer.Formats, "|")))
			cfg.Format = f
		}

		j, err := json.Marshal(cfg)
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
