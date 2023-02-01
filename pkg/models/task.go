package models

import (
	"github.com/koki-develop/todoist-cli/pkg/renderer"
	"github.com/koki-develop/todoist-cli/pkg/util"
)

var (
	_ renderer.Formattable = (Task)(nil)
	_ renderer.Formattable = (Tasks)(nil)
)

type Task map[string]interface{}
type Tasks []Task

func (t Task) Maps() []map[string]interface{} {
	m := util.CloneMap(t)
	m["due"] = t.due()
	return []map[string]interface{}{m}
}

func (t Task) due() string {
	d, ok := t["due"].(map[string]interface{})
	if !ok || d == nil {
		return ""
	}

	if dt := d["datetime"]; dt != nil {
		return dt.(string)
	} else {
		return d["date"].(string)
	}
}

func (ts Tasks) Maps() []map[string]interface{} {
	maps := make([]map[string]interface{}, len(ts))
	for i, t := range ts {
		maps[i] = t.Maps()[0]
	}
	return maps
}
