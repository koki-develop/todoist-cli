package cmd

import (
	"github.com/koki-develop/todoist/pkg/renderer"
	"github.com/spf13/cobra"
)

var (
	format string
)

func init() {
	// output
	for _, cmd := range []*cobra.Command{
		projectsListCmd,
		projectsGetCmd,
	} {
		cmd.Flags().StringVarP(&format, "format", "f", string(renderer.FormatTable), "output format (table|json)")
	}
}
