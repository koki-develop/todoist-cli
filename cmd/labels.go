package cmd

import (
	"fmt"

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
