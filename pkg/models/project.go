package models

import (
	"encoding/json"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/koki-develop/todoist/pkg/renderer"
	"gopkg.in/yaml.v3"
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

var projectHeader table.Row = table.Row{"PROJECT_ID", "NAME", "STYLE"}

func (proj *Project) tableRow() table.Row {
	return table.Row{*proj.ID, *proj.Name, *proj.ViewStyle}
}

func (projs Projects) tableRows() []table.Row {
	rows := []table.Row{}
	for _, proj := range projs {
		rows = append(rows, proj.tableRow())
	}
	return rows
}

func newProjectsTableWriter() table.Writer {
	t := table.NewWriter()
	t.AppendHeader(projectHeader)
	return t
}

func (proj *Project) Table() (string, error) {
	t := newProjectsTableWriter()
	t.AppendRow(proj.tableRow())
	return t.Render(), nil
}

func (proj *Project) JSON() (string, error) {
	j, err := json.MarshalIndent(proj, "", "  ")
	if err != nil {
		return "", err
	}
	return string(j), nil
}

func (proj *Project) YAML() (string, error) {
	y, err := yaml.Marshal(proj)
	if err != nil {
		return "", err
	}
	return string(y), nil
}

func (projs Projects) Table() (string, error) {
	t := newProjectsTableWriter()
	t.AppendRows(projs.tableRows())
	return t.Render(), nil
}

func (projs Projects) JSON() (string, error) {
	j, err := json.MarshalIndent(projs, "", "  ")
	if err != nil {
		return "", err
	}
	return string(j), nil
}

func (projs Projects) YAML() (string, error) {
	y, err := yaml.Marshal(projs)
	if err != nil {
		return "", err
	}
	return string(y), nil
}
