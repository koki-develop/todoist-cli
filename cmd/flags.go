package cmd

import (
	"fmt"
	"strings"

	"github.com/koki-develop/todoist-cli/pkg/flags"
	"github.com/koki-develop/todoist-cli/pkg/renderer"
	"github.com/spf13/cobra"
)

var (
	flagAPIToken = &flags.String{Flag: &flags.Flag{Name: "api-token", Description: "todoist api token", Nullable: true}}
	flagFormat   = &flags.String{Flag: &flags.Flag{Name: "format", ShortName: "f", Description: fmt.Sprintf("output format (%s)", strings.Join(renderer.Formats, "|"))}, Default: "table"}
)

func init() {
	// api token
	for _, cmd := range []*cobra.Command{
		configureCmd,
		// projects
		projectsListCmd,
		projectsGetCmd,
		projectsCreateCmd,
		projectsUpdateCmd,
	} {
		flagAPIToken.Add(cmd)
	}

	// format
	for _, cmd := range []*cobra.Command{
		// projects
		projectsListCmd,
		projectsGetCmd,
		projectsCreateCmd,
		projectsUpdateCmd,
	} {
		flagFormat.Add(cmd)
	}
}
