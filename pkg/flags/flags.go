package flags

import "github.com/spf13/cobra"

type Flag struct {
	Name        string
	ShortName   string
	Description string
	Nullable    bool
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

func (f *String) Add(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.value, f.Name, f.ShortName, f.Default, f.Description)
}

func (f *String) Set(v string) {
	f.value = v
}

func (f *String) Get(cmd *cobra.Command) *string {
	if f.Nullable && !f.Changed(cmd) {
		return nil
	}
	return &f.value
}

func (f *Bool) Add(cmd *cobra.Command) {
	cmd.Flags().BoolVarP(&f.value, f.Name, f.ShortName, f.Default, f.Description)
}

func (f *Bool) Set(v bool) {
	f.value = v
}

func (f *Bool) Get(cmd *cobra.Command) *bool {
	if f.Nullable && !f.Changed(cmd) {
		return nil
	}
	return &f.value
}
