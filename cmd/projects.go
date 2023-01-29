package cmd

import (
	"fmt"
	"os"

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
		cl := rest.New(&rest.Config{Token: os.Getenv("TODOIST_API_TOKEN")})
		rdr := renderer.New(renderer.FormatTable)

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

var projectGetCmd = &cobra.Command{
	Use:  "get <PROJECT_ID>",
	Args: cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		cl := rest.New(&rest.Config{Token: os.Getenv("TODOIST_API_TOKEN")})
		rdr := renderer.New(renderer.FormatTable)

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
		projectGetCmd,
	)
}
