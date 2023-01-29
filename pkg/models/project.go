package models

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/koki-develop/todoist/pkg/renderer"
)

var (
	_ renderer.Formattable = (*Projects)(nil)
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

func (projs Projects) Table() string {
	t := table.NewWriter()

	t.AppendHeader(table.Row{"PROJECT_ID", "NAME", "STYLE"})
	for _, proj := range projs {
		t.AppendRow(table.Row{*proj.ID, *proj.Name, *proj.ViewStyle})
	}

	return t.Render()
}
