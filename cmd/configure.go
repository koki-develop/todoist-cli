package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/koki-develop/todoist-cli/pkg/config"
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
		ipt := &config.Config{
			APIToken: flagAPIToken.Get(cmd, true),
			Format:   (*renderer.Format)(flagFormat.Get(cmd, true)),
		}
		it := ipt.APIToken == nil && ipt.Format == nil

		def := &config.Config{
			APIToken: ipt.APIToken,
			Format:   ipt.Format,
		}
		cfg, _ := config.Load(def)
		if cfg == nil {
			cfg = def
		}

		if cfg.APIToken == nil || it {
			tkn := read("API Token")
			if tkn != "" {
				cfg.APIToken = &tkn
			}
		}

		if ipt.Format == nil && it {
			f := read(fmt.Sprintf("Default Output Format (optional) (%s)", strings.Join(renderer.Formats, "|")))
			if f != "" {
				cfg.Format = (*renderer.Format)(&f)
			}
		}
		if cfg.Format == nil {
			f := renderer.FormatTable
			cfg.Format = &f
		}

		j, err := json.MarshalIndent(cfg, "", "  ")
		if err != nil {
			return err
		}

		cfgdir, err := config.Dir()
		if err != nil {
			return err
		}
		if err := os.MkdirAll(cfgdir, os.ModePerm); err != nil {
			return err
		}
		cfgfile := config.Filename(cfgdir)
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
