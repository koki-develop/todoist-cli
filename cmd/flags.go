package cmd

import (
	"fmt"
	"strings"

	"github.com/koki-develop/todoist-cli/pkg/renderer"
	"github.com/spf13/cobra"
)

type Flag struct {
	Name        string
	ShortName   string
	Description string
	Nullable    bool
}

func (f *Flag) Changed(cmd *cobra.Command) bool {
	return cmd.Flag(f.Name).Changed
}

type FlagString struct {
	*Flag
	Default string
	value   string
}

type FlagBool struct {
	*Flag
	Default bool
	value   bool
}

func (f *FlagString) Add(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.value, f.Name, f.ShortName, f.Default, f.Description)
}

func (f *FlagString) Set(v string) {
	f.value = v
}

func (f *FlagString) Get(cmd *cobra.Command) *string {
	if f.Nullable && !f.Changed(cmd) {
		return nil
	}
	return &f.value
}

func (f *FlagBool) Add(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&f.value, f.Name, f.ShortName, f.Default, f.Description)
}

func (f *FlagBool) Set(v bool) {
	f.value = v
}

func (f *FlagBool) Get(cmd *cobra.Command) *bool {
	if f.Nullable && !f.Changed(cmd) {
		return nil
	}
	return &f.value
}

var (
	flagAPIToken = &FlagString{Flag: &Flag{Name: "api-token", Description: "todoist api token", Nullable: true}}
	flagFormat   = &FlagString{Flag: &Flag{Name: "format", ShortName: "f", Description: fmt.Sprintf("output format (%s)", strings.Join(renderer.Formats, "|"))}, Default: "table"}
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
