package models

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/koki-develop/todoist-cli/pkg/renderer"
)

var (
	_ renderer.Formattable = (*Section)(nil)
	_ renderer.Formattable = (Sections)(nil)
)

type Section struct {
	ID        *string `json:"id"`
	ProjectID *string `json:"project_id"`
	Order     *int    `json:"order"`
	Name      *string `json:"name"`
}

type Sections []*Section

var sectionTableHeader table.Row = table.Row{"ID", "NAME", "PROJECT_ID"}

func (*Section) TableHeader() table.Row {
	return sectionTableHeader
}

func (sec *Section) TableRows() []table.Row {
	return []table.Row{{*sec.ID, *sec.Name, *sec.ProjectID}}
}

func (Sections) TableHeader() table.Row {
	return sectionTableHeader
}

func (secs Sections) TableRows() []table.Row {
	rows := []table.Row{}
	for _, sec := range secs {
		rows = append(rows, sec.TableRows()...)
	}
	return rows
}
