package todoistapi

import (
	"fmt"
	"net/http"

	"github.com/koki-develop/todoist-cli/pkg/models"
)

type ListTasksParameters struct {
	ProjectID *string  `url:"project_id,omitempty"`
	SectionID *string  `url:"section_id,omitempty"`
	Label     *string  `url:"label,omitempty"`
	Filter    *string  `url:"filter,omitempty"`
	Lang      *string  `url:"lang,omitempty"`
	IDs       []string `url:"ids,comma,omitempty"`
}

func (cl *Client) ListTasks(p *ListTasksParameters) (models.Tasks, error) {
	req, err := cl.newRequest(http.MethodGet, "tasks", p, nil)
	if err != nil {
		return nil, err
	}
	fmt.Println(req.URL.String())

	var ts models.Tasks
	if err := cl.doRequest(req, &ts); err != nil {
		return nil, err
	}

	return ts, nil
}
