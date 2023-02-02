package models

import (
	"github.com/koki-develop/todoist-cli/pkg/renderer"
	"github.com/koki-develop/todoist-cli/pkg/util"
)

type Comment map[string]interface{}
type Comments []Comment

var (
	_ renderer.Formattable = (Comment)(nil)
	_ renderer.Formattable = (Comments)(nil)
)

func (c Comment) Maps() []map[string]interface{} {
	m := util.CloneMap(c)
	m["attachment"] = c.attachment()
	return []map[string]interface{}{m}
}

func (c Comment) attachment() string {
	a, ok := c["attachment"].(map[string]interface{})
	if !ok || a == nil {
		return ""
	}

	fn, ok := a["file_name"].(string)
	if !ok {
		return ""
	}

	return fn
}

func (cs Comments) Maps() []map[string]interface{} {
	maps := make([]map[string]interface{}, len(cs))
	for i, c := range cs {
		maps[i] = c.Maps()[0]
	}
	return maps
}
