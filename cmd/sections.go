package cmd

import (
	"fmt"

	"github.com/koki-develop/todoist-cli/pkg/todoistapi"
	"github.com/spf13/cobra"
)

var sectionsCmd = &cobra.Command{
	Use:   "sections",
	Short: "Manage sections",
	Long:  "Manage sections.",
}

var sectionsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all sections",
	Long:  "List all sections.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := load(cmd); err != nil {
			return err
		}

		p := &todoistapi.ListSectionsParameters{
			ProjectID: flagSectionProjectID.Get(cmd, true),
		}
		secs, err := client.ListSections(p)
		if err != nil {
			return err
		}

		o, err := rdr.Render(secs)
		if err != nil {
			return err
		}
		fmt.Println(o)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(sectionsCmd)
	sectionsCmd.AddCommand(
		sectionsListCmd,
	)
}
