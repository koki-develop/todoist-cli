package models

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/koki-develop/todoist-cli/pkg/renderer"
)

var (
	_ renderer.Formattable = (Project)(nil)
	_ renderer.Formattable = (Projects)(nil)
)

type Project map[string]interface{}
type Projects []Project

var projectTableHeader table.Row = table.Row{"ID", "NAME", "VIEW_STYLE", "URL"}

func (Project) TableHeader() table.Row {
	return projectTableHeader
}

func (proj Project) TableRows() []table.Row {
	return []table.Row{{proj["id"], proj["name"], proj["view_style"], proj["url"]}}
}

func (Projects) TableHeader() table.Row {
	return projectTableHeader
}

func (projs Projects) TableRows() []table.Row {
	rows := []table.Row{}
	for _, proj := range projs {
		rows = append(rows, proj.TableRows()...)
	}
	return rows
}
