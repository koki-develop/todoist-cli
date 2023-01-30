package flags

import (
	"github.com/spf13/cobra"
)

type Flag struct {
	Name        string
	ShortName   string
	Description string
	Required    bool
}

func (f *Flag) Changed(cmd *cobra.Command) bool {
	return cmd.Flag(f.Name).Changed
}

type String struct {
	*Flag
	Default string
	value   string
}

type Bool struct {
	*Flag
	Default bool
	value   bool
}

type Int struct {
	*Flag
	Default int
	value   int
}

func (f *String) Add(cmds ...*cobra.Command) {
	for _, cmd := range cmds {
		cmd.Flags().StringVarP(&f.value, f.Name, f.ShortName, f.Default, f.Description)
		if f.Required {
			_ = cmd.MarkFlagRequired(f.Name)
		}
	}
}

func (f *String) Get(cmd *cobra.Command, nullable bool) *string {
	if nullable && !f.Changed(cmd) {
		return nil
	}
	return &f.value
}

func (f *Bool) Add(cmds ...*cobra.Command) {
	for _, cmd := range cmds {
		cmd.Flags().BoolVarP(&f.value, f.Name, f.ShortName, f.Default, f.Description)
		if f.Required {
			_ = cmd.MarkFlagRequired(f.Name)
		}
	}
}

func (f *Bool) Get(cmd *cobra.Command, nullable bool) *bool {
	if nullable && !f.Changed(cmd) {
		return nil
	}
	return &f.value
}

func (f *Int) Add(cmds ...*cobra.Command) {
	for _, cmd := range cmds {
		cmd.Flags().IntVarP(&f.value, f.Name, f.ShortName, f.Default, f.Description)
		if f.Required {
			_ = cmd.MarkFlagRequired(f.Name)
		}
	}
}

func (f *Int) Get(cmd *cobra.Command, nullable bool) *int {
	if nullable && !f.Changed(cmd) {
		return nil
	}
	return &f.value
}
