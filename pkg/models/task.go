package models

import (
	"github.com/koki-develop/todoist-cli/pkg/renderer"
)

var (
	_ renderer.Formattable = (Task)(nil)
	_ renderer.Formattable = (Tasks)(nil)
)

type Task map[string]interface{}
type Tasks []Task

func (t Task) Maps() []map[string]interface{} {
	return []map[string]interface{}{t}
}

func (ts Tasks) Maps() []map[string]interface{} {
	maps := make([]map[string]interface{}, len(ts))
	for i, t := range ts {
		maps[i] = t.Maps()[0]
	}
	return maps
}
