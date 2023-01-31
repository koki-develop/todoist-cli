package todoistapi

import (
	"fmt"
	"net/http"

	"github.com/koki-develop/todoist-cli/pkg/models"
)

type ListTasksParameters struct {
	ProjectID *string `url:"project_id,omitempty"`
	SectionID *string `url:"section_id,omitempty"`
	Label     *string `url:"label,omitempty"`
	Filter    *string `url:"filter,omitempty"`
	Lang      *string `url:"lang,omitempty"`
	IDs       *[]int  `url:"ids,comma,omitempty"`
}

func (cl *Client) ListTasks(p *ListTasksParameters) (models.Tasks, error) {
	req, err := cl.newRequest(http.MethodGet, "tasks", p, nil)
	if err != nil {
		return nil, err
	}

	var ts models.Tasks
	if err := cl.doRequest(req, &ts); err != nil {
		return nil, err
	}

	return ts, nil
}

func (cl *Client) GetTask(id string) (*models.Task, error) {
	req, err := cl.newRequest(http.MethodGet, fmt.Sprintf("tasks/%s", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var t models.Task
	if err := cl.doRequest(req, &t); err != nil {
		return nil, err
	}

	return &t, nil
}

type CreateTaskParameters struct {
	Content     *string   `json:"content,omitempty"`
	Description *string   `json:"description,omitempty"`
	ProjectID   *string   `json:"project_id,omitempty"`
	SectionID   *string   `json:"section_id,omitempty"`
	ParentID    *string   `json:"parent_id,omitempty"`
	Order       *int      `json:"order,omitempty"`
	Labels      *[]string `json:"labels,omitempty"`
	Priority    *int      `json:"priority,omitempty"`
	DueString   *string   `json:"due_string,omitempty"`
	DueDate     *string   `json:"due_date,omitempty"`
	DueDatetime *string   `json:"due_datetime,omitempty"`
	DueLang     *string   `json:"due_lang,omitempty"`
	AssigneeID  *string   `json:"assignee_id,omitempty"`
}

func (cl *Client) CreateTask(p *CreateTaskParameters) (*models.Task, error) {
	req, err := cl.newRequest(http.MethodPost, "tasks", nil, p)
	if err != nil {
		return nil, err
	}

	var t models.Task
	if err := cl.doRequest(req, &t); err != nil {
		return nil, err
	}

	return &t, nil
}

type UpdateTaskParameters struct {
	Content     *string   `json:"content,omitempty"`
	Description *string   `json:"description,omitempty"`
	Labels      *[]string `json:"labels,omitempty"`
	Priority    *int      `json:"priority,omitempty"`
	DueString   *string   `json:"due_string,omitempty"`
	DueDate     *string   `json:"due_date,omitempty"`
	DueDatetime *string   `json:"due_datetime,omitempty"`
	DueLang     *string   `json:"due_lang,omitempty"`
	AssigneeID  *string   `json:"assignee_id,omitempty"`
}

func (cl *Client) UpdateTask(id string, p *UpdateTaskParameters) (*models.Task, error) {
	req, err := cl.newRequest(http.MethodPost, fmt.Sprintf("tasks/%s", id), nil, p)
	if err != nil {
		return nil, err
	}

	var t models.Task
	if err := cl.doRequest(req, &t); err != nil {
		return nil, err
	}

	return &t, nil
}
