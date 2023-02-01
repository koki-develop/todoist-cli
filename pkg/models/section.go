package models

import (
	"github.com/koki-develop/todoist-cli/pkg/renderer"
)

var (
	_ renderer.Formattable = (Section)(nil)
	_ renderer.Formattable = (Sections)(nil)
)

type Section map[string]interface{}
type Sections []Section

func (sec Section) Maps() []map[string]interface{} {
	return []map[string]interface{}{sec}
}

func (secs Sections) Maps() []map[string]interface{} {
	maps := make([]map[string]interface{}, len(secs))
	for i, s := range secs {
		maps[i] = s.Maps()[0]
	}
	return maps
}
