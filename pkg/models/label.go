package models

import (
	"github.com/koki-develop/todoist-cli/pkg/renderer"
)

var (
	_ renderer.Formattable = (Label)(nil)
	_ renderer.Formattable = (Labels)(nil)
)

type Label map[string]interface{}
type Labels []Label

func (l Label) Maps() []map[string]interface{} {
	return []map[string]interface{}{l}
}

func (ls Labels) Maps() []map[string]interface{} {
	maps := make([]map[string]interface{}, len(ls))
	for i, p := range ls {
		maps[i] = p.Maps()[0]
	}
	return maps
}
