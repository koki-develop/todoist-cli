package rest

import (
	"net/http"

	"github.com/koki-develop/todoist/pkg/models"
)

func (cl *Client) ListProjects() (models.Projects, error) {
	req, err := cl.newRequest(http.MethodGet, "projects", nil)
	if err != nil {
		return nil, err
	}

	var projs models.Projects
	if err := cl.doRequest(req, &projs); err != nil {
		return nil, err
	}

	return projs, nil
}
