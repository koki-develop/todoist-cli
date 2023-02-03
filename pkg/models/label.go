package models

import (
	"github.com/koki-develop/todoist-cli/pkg/renderer"
)

var (
	_ renderer.Formattable = (Label)(nil)
	_ renderer.Formattable = (Labels)(nil)
	_ renderer.Formattable = (*SharedLabel)(nil)
	_ renderer.Formattable = (SharedLabels)(nil)
)

type Label map[string]interface{}
type Labels []Label

type SharedLabel string
type SharedLabels []SharedLabel

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

func (l SharedLabel) Maps() []map[string]interface{} {
	return []map[string]interface{}{{"name": l}}
}

func (ls SharedLabels) Maps() []map[string]interface{} {
	maps := make([]map[string]interface{}, len(ls))
	for i, p := range ls {
		maps[i] = p.Maps()[0]
	}
	return maps
}
