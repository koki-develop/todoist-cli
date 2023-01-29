package cmd

import (
	"fmt"

	"github.com/koki-develop/todoist/pkg/renderer"
	"github.com/koki-develop/todoist/pkg/rest"
	"github.com/spf13/cobra"
)

var (
	projectParentID   string
	projectName       string
	projectColor      string
	projectIsFavorite bool
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
		cfg, err := loadConfig()
		if err != nil {
			return ErrLoadConfig
		}

		cl := rest.New(&rest.Config{Token: cfg.APIToken})
		rdr := renderer.New(renderer.Format(format))

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

var projectsGetCmd = &cobra.Command{
	Use:   "get <PROJECT_ID>",
	Short: "Get a project",
	Long:  "Get a project.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		cfg, err := loadConfig()
		if err != nil {
			return ErrLoadConfig
		}

		cl := rest.New(&rest.Config{Token: cfg.APIToken})
		rdr := renderer.New(renderer.Format(format))

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

var projectsCreateCmd = &cobra.Command{
	Use:   "create <PROJECT_NAME>",
	Short: "Create a project",
	Long:  "Create a project.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		cfg, err := loadConfig()
		if err != nil {
			return ErrLoadConfig
		}

		cl := rest.New(&rest.Config{Token: cfg.APIToken})
		rdr := renderer.New(renderer.Format(format))

		p := &rest.CreateProjectPayload{Name: name}
		if cmd.Flag("parent-id").Changed {
			p.ParentID = &projectParentID
		}
		if cmd.Flag("favorite").Changed {
			p.IsFavorite = &projectIsFavorite
		}
		if cmd.Flag("color").Changed {
			p.Color = &projectColor
		}

		proj, err := cl.CreateProject(p)
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
	Use:   "update <PROJECT_ID>",
	Short: "Update a project",
	Long:  "Update a project.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		cfg, err := loadConfig()
		if err != nil {
			return ErrLoadConfig
		}

		cl := rest.New(&rest.Config{Token: cfg.APIToken})
		rdr := renderer.New(renderer.Format(format))

		p := &rest.UpdateProjectPayload{}
		if cmd.Flag("name").Changed {
			p.Name = &projectName
		}
		if cmd.Flag("color").Changed {
			p.Color = &projectColor
		}
		if cmd.Flag("favorite").Changed {
			p.IsFavorite = &projectIsFavorite
		}

		proj, err := cl.UpdateProject(id, p)
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
	Use:   "delete <PROJECT_ID>",
	Short: "Delete a project",
	Long:  "Delete a project.",
	Args:  cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
	RunE: func(cmd *cobra.Command, args []string) error {
		id := args[0]

		cfg, err := loadConfig()
		if err != nil {
			return ErrLoadConfig
		}

		cl := rest.New(&rest.Config{Token: cfg.APIToken})

		if err := cl.DeleteProject(id); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(projectsCmd)
	projectsCmd.AddCommand(
		projectsListCmd,
		projectsGetCmd,
		projectsCreateCmd,
		projectsUpdateCmd,
		projectsDeleteCmd,
	)

	// create
	projectsCreateCmd.Flags().StringVar(&projectParentID, "parent-id", "", "parent project id")
	projectsCreateCmd.Flags().StringVar(&projectColor, "color", "", "the color of the project icon")
	projectsCreateCmd.Flags().BoolVar(&projectIsFavorite, "favorite", false, "whether the project is a favorite")

	// update
	projectsUpdateCmd.Flags().StringVar(&projectName, "name", "", "name of the project")
	projectsUpdateCmd.Flags().StringVar(&projectColor, "color", "", "the color of the project icon")
	projectsUpdateCmd.Flags().BoolVar(&projectIsFavorite, "favorite", false, "whether the project is a favorite")
}
