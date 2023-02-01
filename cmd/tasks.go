package cmd

import (
	"fmt"

	"github.com/koki-develop/todoist-cli/pkg/todoistapi"
	"github.com/spf13/cobra"
)

var tasksCmd = &cobra.Command{
	Use:     "tasks",
	Aliases: []string{"task", "t"},
	Short:   "Manage tasks",
	Long:    "Manage tasks.",
}

var tasksListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all active tasks",
	Long:  "List all active tasks.",
	Args:  cobra.MaximumNArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := load(cmd); err != nil {
			return err
		}

		p := &todoistapi.ListTasksParameters{
			ProjectID: flagTasksListProjectID.Get(cmd, true),
			SectionID: flagTasksListSectionID.Get(cmd, true),
			Label:     flagTasksListLabel.Get(cmd, true),
			Filter:    flagTasksListFilter.Get(cmd, true),
			Lang:      flagTasksListLang.Get(cmd, true),
			IDs:       flagTasksListIDs.Get(cmd, true),
		}
		ts, err := client.ListTasks(p)
		if err != nil {
			return err
		}

		o, err := rdr.Render(ts, *flagColumnsTask.Get(cmd, false))
		if err != nil {
			return err
		}
		fmt.Println(o)
		return nil
	},
}

var tasksGetCmd = &cobra.Command{
	Use:   "get TASK_ID",
	Short: "Get an active task",
	Long:  "Get an active task.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		if err := load(cmd); err != nil {
			return err
		}

		t, err := client.GetTask(id)
		if err != nil {
			return err
		}

		o, err := rdr.Render(t, *flagColumnsTask.Get(cmd, false))
		if err != nil {
			return err
		}
		fmt.Println(o)
		return nil
	},
}

var tasksCreateCmd = &cobra.Command{
	Use:   "create TASK_CONTENT",
	Short: "Create a task",
	Long:  "Create a task.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
	RunE: func(cmd *cobra.Command, args []string) error {
		content := args[0]

		if err := load(cmd); err != nil {
			return err
		}

		p := &todoistapi.CreateTaskParameters{
			Content:     &content,
			Description: flagTasksCreateDescription.Get(cmd, true),
			ProjectID:   flagTasksCreateProjectID.Get(cmd, true),
			SectionID:   flagTasksCreateSectionID.Get(cmd, true),
			ParentID:    flagTasksCreateParentID.Get(cmd, true),
			Order:       flagTasksCreateOrder.Get(cmd, true),
			Labels:      flagTasksCreateLabels.Get(cmd, true),
			Priority:    flagTasksCreatePriority.Get(cmd, true),
			DueString:   flagTasksCreateDueString.Get(cmd, true),
			DueDate:     flagTasksCreateDueDate.Get(cmd, true),
			DueDatetime: flagTasksCreateDueDatetime.Get(cmd, true),
			DueLang:     flagTasksCreateDueLang.Get(cmd, true),
			AssigneeID:  flagTasksCreateAssigneeID.Get(cmd, true),
		}
		t, err := client.CreateTask(p)
		if err != nil {
			return err
		}

		o, err := rdr.Render(t, *flagColumnsTask.Get(cmd, false))
		if err != nil {
			return err
		}
		fmt.Println(o)
		return nil
	},
}

var tasksUpdateCmd = &cobra.Command{
	Use:   "update TASK_ID",
	Short: "Update a task",
	Long:  "Update a task.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		if err := load(cmd); err != nil {
			return err
		}

		p := &todoistapi.UpdateTaskParameters{
			Content:     flagTasksUpdateContent.Get(cmd, true),
			Description: flagTasksUpdateDescription.Get(cmd, true),
			Labels:      flagTasksUpdateLabels.Get(cmd, true),
			Priority:    flagTasksUpdatePriority.Get(cmd, true),
			DueString:   flagTasksUpdateDueString.Get(cmd, true),
			DueDate:     flagTasksUpdateDueDate.Get(cmd, true),
			DueDatetime: flagTasksUpdateDueDatetime.Get(cmd, true),
			DueLang:     flagTasksUpdateDueLang.Get(cmd, true),
			AssigneeID:  flagTasksUpdateAssigneeID.Get(cmd, true),
		}
		t, err := client.UpdateTask(id, p)
		if err != nil {
			return err
		}

		o, err := rdr.Render(t, *flagColumnsTask.Get(cmd, false))
		if err != nil {
			return err
		}
		fmt.Println(o)
		return nil
	},
}

var tasksDeleteCmd = &cobra.Command{
	Use:   "delete TASK_ID",
	Short: "Delete a task",
	Long:  "Delete a task.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		if err := load(cmd); err != nil {
			return err
		}

		if err := client.DeleteTask(id); err != nil {
			return err
		}

		return nil
	},
}

var tasksCloseCmd = &cobra.Command{
	Use:   "close TASK_ID",
	Short: "Close a task",
	Long:  "Close a task.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		if err := load(cmd); err != nil {
			return err
		}

		if err := client.CloseTask(id); err != nil {
			return err
		}

		return nil
	},
}

var tasksReopenCmd = &cobra.Command{
	Use:   "reopen TASK_ID",
	Short: "reopen a task",
	Long:  "reopen a task.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		if err := load(cmd); err != nil {
			return err
		}

		if err := client.ReopenTask(id); err != nil {
			return err
		}

		return nil
	},
}
