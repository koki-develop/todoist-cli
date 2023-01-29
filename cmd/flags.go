package cmd

import (
	"fmt"
	"strings"

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
		cmd.Flags().StringVarP(&format, "format", "f", string(renderer.FormatTable), fmt.Sprintf("output format (%s)", strings.Join(renderer.Formats, "|")))
	}
}
