package cmd

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/spf13/cobra"
)

var (
	version string
)

var rootCmd = &cobra.Command{
	Use:  "todoist-cli",
	Long: "CLI Client for Todoist.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		switch err {
		case ErrLoadConfig, ErrAPITokenNotConfigured:
			fmt.Fprintln(os.Stderr, "Run `todoist-cli configure` to reconfigure.")
		}
		os.Exit(1)
	}
}

func init() {
	/*
	 * version
	 */
	if version == "" {
		if info, ok := debug.ReadBuildInfo(); ok {
			version = info.Main.Version
		}
	}
	rootCmd.Version = version
	_ = notifyNewRelease(os.Stderr)

	/*
	 * add commands
	 */

	// configure
	rootCmd.AddCommand(configureCmd)

	// projects
	rootCmd.AddCommand(projectsCmd)
	projectsCmd.AddCommand(
		projectsListCmd,
		projectsGetCmd,
		projectsCreateCmd,
		projectsUpdateCmd,
		projectsDeleteCmd,
	)

	// sections
	rootCmd.AddCommand(sectionsCmd)
	sectionsCmd.AddCommand(
		sectionsListCmd,
		sectionsGetCmd,
		sectionsCreateCmd,
		sectionsUpdateCmd,
	)

	/*
	 * flags
	 */

	// --api-token
	flagAPIToken.Add(configureCmd)
	flagAPIToken.Add(projectsCmd.Commands()...)
	flagAPIToken.Add(sectionsCmd.Commands()...)

	// -f, --format
	flagFormat.Add(configureCmd)
	flagFormat.Add(projectsCmd.Commands()...)
	flagFormat.Add(sectionsCmd.Commands()...)

	// --parent-id
	flagProjectParentID.Add(
		projectsCreateCmd,
	)

	// --color
	flagProjectColor.Add(
		projectsCreateCmd,
		projectsUpdateCmd,
	)

	// --favorite
	flagProjectFavorite.Add(
		projectsCreateCmd,
		projectsUpdateCmd,
	)

	// --name
	flagProjectName.Add(
		projectsUpdateCmd,
	)
	flagSectionNameForUpdate.Add(
		sectionsUpdateCmd,
	)

	// --project-id
	flagSectionProjectID.Add(
		sectionsListCmd,
	)
	flagSectionProjectIDForCreate.Add(
		sectionsCreateCmd,
	)

	// --order
	flagSectionOrder.Add(
		sectionsCreateCmd,
	)
}
