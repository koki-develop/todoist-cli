package models

import (
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/koki-develop/todoist-cli/pkg/renderer"
)

var (
	_ renderer.Formattable = (Task)(nil)
	_ renderer.Formattable = (Tasks)(nil)
)

type Task map[string]interface{}

type TaskDue struct {
	Date        *string `json:"date"`
	IsRecurring *bool   `json:"is_recurring"`
	Datetime    *string `json:"datetime"`
	String      *string `json:"string"`
	Timezone    *string `json:"timezone"`
}

type Tasks []Task

var taskTableHeader table.Row = table.Row{"ID", "CONTENT", "DUE", "LABELS", "URL"}

func (Task) TableHeader() table.Row {
	return taskTableHeader
}

func (t Task) due() string {
	if due, ok := t["due"].(map[string]interface{}); ok && due != nil {
		if dt := due["datetime"]; dt != nil {
			return dt.(string)
		} else {
			return due["date"].(string)
		}
	}

	return ""
}

func (t Task) labels() []string {
	labels := t["labels"].([]interface{})
	rtn := make([]string, len(labels))
	for i, l := range labels {
		rtn[i] = l.(string)
	}

	return rtn
}

func (t Task) TableRows() []table.Row {
	return []table.Row{{t["id"], t["content"], t.due(), strings.Join(t.labels(), ", "), t["url"]}}
}

func (Tasks) TableHeader() table.Row {
	return taskTableHeader
}

func (ts Tasks) TableRows() []table.Row {
	rows := []table.Row{}
	for _, t := range ts {
		rows = append(rows, t.TableRows()...)
	}
	return rows
}
