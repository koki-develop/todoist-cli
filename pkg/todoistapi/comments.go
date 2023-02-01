package todoistapi

import (
	"net/http"

	"github.com/koki-develop/todoist-cli/pkg/models"
)

type ListCommentsParameters struct {
	ProjectID *string `url:"project_id,omitempty"`
	TaskID    *string `url:"task_id,omitempty"`
}

func (cl *Client) ListComments(p *ListCommentsParameters) (models.Comments, error) {
	req, err := cl.newRequest(http.MethodGet, "comments", p, nil)
	if err != nil {
		return nil, err
	}

	var cs models.Comments
	if err := cl.doRequest(req, &cs); err != nil {
		return nil, err
	}

	return cs, nil
}
