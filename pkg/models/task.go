package models

import (
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/koki-develop/todoist-cli/pkg/renderer"
)

var (
	_ renderer.Formattable = (*Task)(nil)
	_ renderer.Formattable = (Tasks)(nil)
)

type Task struct {
	ID           *string  `json:"id"`
	AssignerID   *string  `json:"assigner_id"`
	AssigneeID   *string  `json:"assignee_id"`
	ProjectID    *string  `json:"project_id"`
	SectionID    *string  `json:"section_id"`
	ParentID     *string  `json:"parent_id"`
	Order        *int     `json:"order"`
	Content      *string  `json:"content"`
	Description  *string  `json:"description"`
	IsCompleted  *bool    `json:"is_completed"`
	Labels       []string `json:"labels"`
	Priority     *int     `json:"priority"`
	CommentCount *int     `json:"comment_count"`
	CreatorID    *string  `json:"creator_id"`
	CreatedAt    *string  `json:"created_at"`
	Due          *TaskDue `json:"due"`
	URL          *string  `json:"url"`
}

type TaskDue struct {
	Date        *string `json:"date"`
	IsRecurring *bool   `json:"is_recurring"`
	Datetime    *string `json:"datetime"`
	String      *string `json:"string"`
	Timezone    *string `json:"timezone"`
}

type Tasks []*Task

var taskTableHeader table.Row = table.Row{"ID", "CONTENT", "DUE", "LABELS", "URL"}

func (*Task) TableHeader() table.Row {
	return taskTableHeader
}

func (t *Task) TableRows() []table.Row {
	var d string = ""
	if t.Due != nil {
		if t.Due.Datetime != nil {
			d = *t.Due.Datetime
		} else {
			d = *t.Due.Date
		}
	}

	return []table.Row{{*t.ID, *t.Content, d, strings.Join(t.Labels, ", "), *t.URL}}
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
