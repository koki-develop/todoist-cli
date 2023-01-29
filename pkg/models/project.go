package models

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/koki-develop/todoist/pkg/renderer"
)

var (
	_ renderer.Formattable = (*Project)(nil)
	_ renderer.Formattable = (Projects)(nil)
)

type Project struct {
	ID             *string `json:"id"`
	ParentID       *string `json:"parent_id"`
	Order          *int    `json:"order"`
	Color          *string `json:"color"`
	Name           *string `json:"name"`
	CommentCount   *int    `json:"comment_count"`
	IsShared       *bool   `json:"is_shared"`
	IsFavorite     *bool   `json:"is_favorite"`
	IsInboxProject *bool   `json:"is_inbox_project"`
	IsTeamInbox    *bool   `json:"is_team_inbox"`
	URL            *string `json:"url"`
	ViewStyle      *string `json:"view_style"`
}

type Projects []*Project

var projectTableHeader table.Row = table.Row{"PROJECT_ID", "NAME", "STYLE"}

func (*Project) TableHeader() table.Row {
	return projectTableHeader
}

func (proj *Project) TableRows() []table.Row {
	return []table.Row{{*proj.ID, *proj.Name, *proj.ViewStyle}}
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
