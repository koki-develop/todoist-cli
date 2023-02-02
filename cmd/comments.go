package cmd

import (
	"fmt"

	"github.com/koki-develop/todoist-cli/pkg/todoistapi"
	"github.com/spf13/cobra"
)

var commentsCmd = &cobra.Command{
	Use:     "comments",
	Aliases: []string{"comment", "c"},
	Short:   "Manage comments",
	Long:    "Manage comments.",
}

var commentsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all comments for a task or project",
	Long:  "List all comments for a task or project.",
	Args:  cobra.MaximumNArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := load(cmd); err != nil {
			return err
		}

		p := &todoistapi.ListCommentsParameters{
			ProjectID: flagCommentsListProjectID.Get(cmd, true),
			TaskID:    flagCommentsListTaskID.Get(cmd, true),
		}
		cs, err := client.ListComments(p)
		if err != nil {
			return err
		}

		o, err := rdr.Render(cs, *flagColumnsComment.Get(cmd, false))
		if err != nil {
			return err
		}
		fmt.Println(o)
		return nil
	},
}

var commentsGetCmd = &cobra.Command{
	Use:   "get COMMENT_ID",
	Short: "Get a comment",
	Long:  "Get a comment.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		if err := load(cmd); err != nil {
			return err
		}

		c, err := client.GetComment(id)
		if err != nil {
			return err
		}

		o, err := rdr.Render(c, *flagColumnsComment.Get(cmd, false))
		if err != nil {
			return err
		}
		fmt.Println(o)
		return nil
	},
}
