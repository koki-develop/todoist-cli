package cmd

import (
	"github.com/koki-develop/todoist-cli/pkg/renderer"
	"github.com/koki-develop/todoist-cli/pkg/todoistapi"
	"github.com/spf13/cobra"
)

var (
	config *Config            = nil
	client *todoistapi.Client = nil
	rdr    *renderer.Renderer = nil
)

func load(cmd *cobra.Command) error {
	cfg, err := loadConfig(cmd)
	if err != nil {
		return ErrLoadConfig
	}

	config = cfg
	client = todoistapi.New(&todoistapi.Config{Token: config.APIToken})
	rdr = renderer.New(renderer.Format(*flagFormat.Get(cmd, false)))

	return nil
}
