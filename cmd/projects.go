package cmd

import (
	"fmt"

	"github.com/koki-develop/todoist/pkg/renderer"
	"github.com/koki-develop/todoist/pkg/rest"
	"github.com/spf13/cobra"
)

var (
	projectParentID   string
	projectColor      string
	projectIsFavorite bool
	projectViewStyle  string
)

var projectsCmd = &cobra.Command{
	Use: "projects",
}

var projectsListCmd = &cobra.Command{
	Use: "list",
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
	Use:  "get <PROJECT_ID>",
	Args: cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
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
	Use:  "create <PROJECT_NAME>",
	Args: cobra.MatchAll(cobra.MinimumNArgs(1), cobra.MaximumNArgs(1)),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]

		cfg, err := loadConfig()
		if err != nil {
			return ErrLoadConfig
		}

		cl := rest.New(&rest.Config{Token: cfg.APIToken})
		rdr := renderer.New(renderer.Format(format))

		p := &rest.CreateProjectPayload{
			Name:       name,
			IsFavorite: &projectIsFavorite,
		}
		if projectParentID != "" {
			p.ParentID = &projectParentID
		}
		if projectColor != "" {
			p.Color = &projectColor
		}
		if projectViewStyle != "" {
			p.ViewStyle = &projectViewStyle
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

func init() {
	rootCmd.AddCommand(projectsCmd)
	projectsCmd.AddCommand(
		projectsListCmd,
		projectsGetCmd,
		projectsCreateCmd,
	)

	projectsCreateCmd.Flags().StringVar(&projectParentID, "parent-id", "", "parent project id")
	projectsCreateCmd.Flags().StringVar(&projectColor, "color", "", "the color of the project icon")
	projectsCreateCmd.Flags().BoolVar(&projectIsFavorite, "favorite", false, "whether the project is a favorite")
	projectsCreateCmd.Flags().StringVar(&projectViewStyle, "view-style", "list", "this determines the way the project is displayed within the todoist clients")
}
