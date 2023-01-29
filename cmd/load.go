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
	ErrLoadConfig            = errors.New("failed to load config")
	ErrAPITokenNotConfigured = errors.New("api token is not configured")
)

func load(cmd *cobra.Command) error {
	cfg, err := config.Load(&config.Config{
		APIToken: flagAPIToken.Get(cmd, true),
		Format:   (*renderer.Format)(flagFormat.Get(cmd, true)),
	})
	if err != nil {
		return ErrLoadConfig
	}
	if cfg.APIToken == nil {
		return ErrAPITokenNotConfigured
	}

	client = todoistapi.New(&todoistapi.Config{Token: *cfg.APIToken})
	rdr = renderer.New(renderer.Format(*cfg.Format))

	return nil
}
