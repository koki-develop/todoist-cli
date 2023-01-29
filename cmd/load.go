package cmd

import (
	"errors"

	"github.com/koki-develop/todoist-cli/pkg/config"
	"github.com/koki-develop/todoist-cli/pkg/renderer"
	"github.com/koki-develop/todoist-cli/pkg/todoistapi"
	"github.com/spf13/cobra"
)

var (
	client *todoistapi.Client = nil
	rdr    *renderer.Renderer = nil
)

var (
	ErrLoadConfig = errors.New("failed to load config")
)

func load(cmd *cobra.Command) error {
	cfg, err := config.Load(cmd, &config.Config{
		APIToken: flagAPIToken.Get(cmd, true),
		Format:   (*renderer.Format)(flagFormat.Get(cmd, true)),
	})
	if err != nil {
		return ErrLoadConfig
	}

	client = todoistapi.New(&todoistapi.Config{Token: *cfg.APIToken})
	rdr = renderer.New(renderer.Format(*cfg.Format))

	return nil
}
