package cmd

import (
	"fmt"

	"github.com/koki-develop/todoist-cli/pkg/todoistapi"
	"github.com/spf13/cobra"
)

var labelsCmd = &cobra.Command{
	Use:     "labels",
	Aliases: []string{"label", "l"},
	Short:   "Manage labels",
	Long:    "Manage labels.",
}

var labelsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all labels",
	Long:  "List all labels",
	Args:  cobra.MaximumNArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := load(cmd); err != nil {
			return err
		}

		ls, err := client.ListLabels()
		if err != nil {
			return err
		}

		o, err := rdr.Render(ls, *flagColumnsLabel.Get(cmd, false))
		if err != nil {
			return err
		}
		fmt.Println(o)
		return nil
	},
}

var labelsGetCmd = &cobra.Command{
	Use:   "get LABEL_ID",
	Short: "Get a label",
	Long:  "Get a label.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		if err := load(cmd); err != nil {
			return err
		}

		l, err := client.GetLabel(id)
		if err != nil {
			return err
		}

		o, err := rdr.Render(l, *flagColumnsLabel.Get(cmd, false))
		if err != nil {
			return err
		}
		fmt.Println(o)
		return nil
	},
}

var labelsCreateCmd = &cobra.Command{
	Use:   "create LABEL_NAME",
	Short: "Create a label",
	Long:  "Create a label.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		if err := load(cmd); err != nil {
			return err
		}

		p := &todoistapi.CreateLabelParameters{
			Name:       &name,
			Order:      flagLabelsCreateOrder.Get(cmd, true),
			Color:      flagLabelsCreateColor.Get(cmd, true),
			IsFavorite: flagLabelsCreateFavorite.Get(cmd, true),
		}
		l, err := client.CreateLabel(p)
		if err != nil {
			return err
		}

		o, err := rdr.Render(l, *flagColumnsLabel.Get(cmd, false))
		if err != nil {
			return err
		}
		fmt.Println(o)
		return nil
	},
}

var labelsUpdateCmd = &cobra.Command{
	Use:   "update LABEL_ID",
	Short: "Update a label",
	Long:  "Update a label.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		if err := load(cmd); err != nil {
			return err
		}

		p := &todoistapi.UpdateLabelParameters{
			Name:       flagLabelsUpdateName.Get(cmd, true),
			Order:      flagLabelsUpdateOrder.Get(cmd, true),
			Color:      flagLabelsUpdateColor.Get(cmd, true),
			IsFavorite: flagLabelsUpdateFavorite.Get(cmd, true),
		}
		l, err := client.UpdateLabel(id, p)
		if err != nil {
			return err
		}

		o, err := rdr.Render(l, *flagColumnsLabel.Get(cmd, false))
		if err != nil {
			return err
		}
		fmt.Println(o)
		return nil
	},
}

var labelsDeleteCmd = &cobra.Command{
	Use:   "delete LABEL_ID",
	Short: "Delete a label",
	Long:  "Delete a label.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		if err := load(cmd); err != nil {
			return err
		}

		if err := client.DeleteLabel(id); err != nil {
			return err
		}

		return nil
	},
}

var sharedLabelsCmd = &cobra.Command{
	Use:     "shared-labels",
	Aliases: []string{"shared-label"},
	Short:   "Manage shared labels",
	Long:    "Manage shared labels.",
}

var sharedLabelsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all shared labels currently assigned to tasks",
	Long:  "List all shared labels currently assigned to tasks.",
	Args:  cobra.MaximumNArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := load(cmd); err != nil {
			return err
		}

		p := &todoistapi.ListSharedLabelsParameters{
			OmitPersonal: flagSharedLabelsListOmitPersonal.Get(cmd, true),
		}
		ls, err := client.ListSharedLabels(p)
		if err != nil {
			return err
		}

		o, err := rdr.Render(ls, *flagColumnsSharedLabel.Get(cmd, false))
		if err != nil {
			return err
		}
		fmt.Println(o)
		return nil
	},
}

var sharedLabelsRenameCmd = &cobra.Command{
	Use:   "rename LABEL_NAME",
	Short: "Rename all instances of a shared label",
	Long:  "Rename all instances of a shared label.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		if err := load(cmd); err != nil {
			return err
		}

		p := &todoistapi.RenameSharedLabelParameters{
			Name:    &name,
			NewName: flagSharedLabelsRenameNewName.Get(cmd, true),
		}
		if err := client.RenameSharedLabel(p); err != nil {
			return err
		}

		return nil
	},
}

var sharedLabelsRemoveCmd = &cobra.Command{
	Use:   "remove LABEL_NAME",
	Short: "Remove all instances of a shared label from the tasks where it is applied",
	Long:  "Remove all instances of a shared label from the tasks where it is applied.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		if err := load(cmd); err != nil {
			return err
		}

		p := &todoistapi.RemoveSharedLabelParameters{
			Name: &name,
		}
		if err := client.RemoveSharedLabel(p); err != nil {
			return err
		}

		return nil
	},
}
