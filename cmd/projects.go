package cmd

import (
	"fmt"

	"github.com/koki-develop/todoist/pkg/renderer"
	"github.com/koki-develop/todoist/pkg/rest"
	"github.com/spf13/cobra"
)

var projectsCmd = &cobra.Command{
	Use: "projects",
}

var projectsListCmd = &cobra.Command{
	Use: "list",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := loadConfig()
		if err != nil {
			return ErrLoadConfig
		}

		cl := rest.New(&rest.Config{Token: cfg.APIToken})
		rdr := renderer.New(renderer.Format(format))

		projs, err := cl.ListProjects()
		if err != nil {
			return err
		}

		o, err := rdr.Render(projs)
		if err != nil {
			return err
		}

		fmt.Println(o)

		return nil
	},
}

var projectsGetCmd = &cobra.Command{
	Use:  "get <PROJECT_ID>",
	Args: cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		cfg, err := loadConfig()
		if err != nil {
			return ErrLoadConfig
		}

		cl := rest.New(&rest.Config{Token: cfg.APIToken})
		rdr := renderer.New(renderer.Format(format))

		proj, err := cl.GetProject(id)
		if err != nil {
			return err
		}

		o, err := rdr.Render(proj)
		if err != nil {
			return err
		}

		fmt.Println(o)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(projectsCmd)
	projectsCmd.AddCommand(
		projectsListCmd,
		projectsGetCmd,
	)
}
