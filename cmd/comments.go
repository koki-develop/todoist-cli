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

var commentsCreateCmd = &cobra.Command{
	Use:   "create CONTENT",
	Short: "Create a comment",
	Long:  "Create a comment.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
	RunE: func(cmd *cobra.Command, args []string) error {
		content := args[0]

		if err := load(cmd); err != nil {
			return err
		}

		var a *todoistapi.CreateCommentAttachment = nil
		fn := flagCommentsCreateFileName.Get(cmd, true)
		fu := flagCommentsCreateFileURL.Get(cmd, true)
		ft := flagCommentsCreateFileType.Get(cmd, true)
		if fn != nil || fu != nil || ft != nil {
			a = &todoistapi.CreateCommentAttachment{
				FileName: fn,
				FileURL:  fu,
				FileType: ft,
			}
		}
		p := &todoistapi.CreateCommentParameters{
			TaskID:     flagCommentsCreateTaskID.Get(cmd, true),
			ProjectID:  flagCommentsCreateProjectID.Get(cmd, true),
			Content:    &content,
			Attachment: a,
		}
		c, err := client.CreateComment(p)
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

var commentsUpdateCmd = &cobra.Command{
	Use:   "update COMMENT_ID",
	Short: "Update a comment",
	Long:  "Update a comment.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		if err := load(cmd); err != nil {
			return err
		}

		p := &todoistapi.UpdateCommentParameters{
			Content: flagCommentsUpdateContent.Get(cmd, false),
		}
		c, err := client.UpdateComment(id, p)
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
