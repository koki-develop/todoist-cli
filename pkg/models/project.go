package models

import (
	"github.com/koki-develop/todoist-cli/pkg/renderer"
)

var (
	_ renderer.Formattable = (Project)(nil)
	_ renderer.Formattable = (Projects)(nil)
)

type Project map[string]interface{}
type Projects []Project

func (proj Project) Maps() []map[string]interface{} {
	return []map[string]interface{}{proj}
}

func (projs Projects) Maps() []map[string]interface{} {
	maps := make([]map[string]interface{}, len(projs))
	for i, p := range projs {
		maps[i] = p.Maps()[0]
	}
	return maps
}
