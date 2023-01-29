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

func init() {
	rootCmd.AddCommand(projectsCmd)
	projectsCmd.AddCommand(projectsListCmd)
}
