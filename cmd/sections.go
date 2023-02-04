package cmd

import (
	"fmt"

	"github.com/koki-develop/todoist-cli/pkg/todoistapi"
	"github.com/spf13/cobra"
)

var sectionsCmd = &cobra.Command{
	Use:     "sections",
	Aliases: []string{"sections", "s"},
	Short:   "Manage sections",
	Long:    "Manage sections.",
}

var sectionsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all sections",
	Long:  "List all sections.",
	Args:  cobra.MaximumNArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := load(cmd); err != nil {
			return err
		}

		p := &todoistapi.ListSectionsParameters{
			ProjectID: flagSectionsListProjectID.Get(cmd, true),
		}
		secs, err := client.ListSections(p)
		if err != nil {
			return err
		}

		o, err := rdr.Render(secs, *flagColumnsSection.Get(cmd, false))
		if err != nil {
			return err
		}
		fmt.Println(o)
		return nil
	},
}

var sectionsGetCmd = &cobra.Command{
	Use:   "get SECTION_ID",
	Short: "Get a section",
	Long:  "Get a section.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		if err := load(cmd); err != nil {
			return err
		}

		sec, err := client.GetSection(id)
		if err != nil {
			return err
		}

		o, err := rdr.Render(sec, *flagColumnsSection.Get(cmd, false))
		if err != nil {
			return err
		}
		fmt.Println(o)
		return nil
	},
}

var sectionsCreateCmd = &cobra.Command{
	Use:   "create SECTION_NAME",
	Short: "Create a section",
	Long:  "Create a section.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		if err := load(cmd); err != nil {
			return err
		}

		p := &todoistapi.CreateSectionParameters{
			Name:      &name,
			ProjectID: flagSectionsCreateProjectID.Get(cmd, false),
			Order:     flagSectionsCreateOrder.Get(cmd, true),
		}
		sec, err := client.CreateSection(p)
		if err != nil {
			return err
		}

		o, err := rdr.Render(sec, *flagColumnsSection.Get(cmd, false))
		if err != nil {
			return err
		}
		fmt.Println(o)
		return nil
	},
}

var sectionsUpdateCmd = &cobra.Command{
	Use:   "update SECTION_ID",
	Short: "Update a section",
	Long:  "Update a section.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		if err := load(cmd); err != nil {
			return err
		}

		p := &todoistapi.UpdateSectionParameters{
			Name: flagSectionsUpdateName.Get(cmd, false),
		}
		sec, err := client.UpdateSection(id, p)
		if err != nil {
			return err
		}

		o, err := rdr.Render(sec, *flagColumnsSection.Get(cmd, false))
		if err != nil {
			return err
		}
		fmt.Println(o)
		return nil
	},
}

var sectionsDeleteCmd = &cobra.Command{
	Use:   "delete SECTION_ID",
	Short: "Delete a section",
	Long:  "Delete a section.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		if err := load(cmd); err != nil {
			return err
		}

		if err := client.DeleteSection(id); err != nil {
			return err
		}

		return nil
	},
}
