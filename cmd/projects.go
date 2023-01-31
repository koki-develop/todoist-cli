package cmd

import (
	"fmt"

	"github.com/koki-develop/todoist-cli/pkg/todoistapi"
	"github.com/spf13/cobra"
)

var projectsCmd = &cobra.Command{
	Use:   "projects",
	Short: "Manage projects",
	Long:  "Manage projects.",
}

var projectsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all projects",
	Long:  "List all projects.",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := load(cmd); err != nil {
			return err
		}

		projs, err := client.ListProjects()
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
	Use:   "get PROJECT_ID",
	Short: "Get a project",
	Long:  "Get a project.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]
		if err := load(cmd); err != nil {
			return err
		}

		proj, err := client.GetProject(id)
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

var projectsCreateCmd = &cobra.Command{
	Use:   "create PROJECT_NAME",
	Short: "Create a project",
	Long:  "Create a project.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		if err := load(cmd); err != nil {
			return err
		}

		p := &todoistapi.CreateProjectParameters{
			Name:       name,
			ParentID:   flagProjectsCreateParentID.Get(cmd, true),
			IsFavorite: flagProjectsCreateFavorite.Get(cmd, true),
			Color:      flagProjectsCreateColor.Get(cmd, true),
		}

		proj, err := client.CreateProject(p)
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

var projectsUpdateCmd = &cobra.Command{
	Use:   "update PROJECT_ID",
	Short: "Update a project",
	Long:  "Update a project.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		if err := load(cmd); err != nil {
			return err
		}

		p := &todoistapi.UpdateProjectParameters{
			Name:       flagProjectsUpdateName.Get(cmd, true),
			Color:      flagProjectsCreateColor.Get(cmd, true),
			IsFavorite: flagProjectsCreateFavorite.Get(cmd, true),
		}

		proj, err := client.UpdateProject(id, p)
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

var projectsDeleteCmd = &cobra.Command{
	Use:   "delete PROJECT_ID",
	Short: "Delete a project",
	Long:  "Delete a project.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		if err := load(cmd); err != nil {
			return err
		}

		if err := client.DeleteProject(id); err != nil {
			return err
		}

		return nil
	},
}
