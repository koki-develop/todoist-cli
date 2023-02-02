package todoistapi

import (
	"net/http"

	"github.com/koki-develop/todoist-cli/pkg/models"
)

func (cl *Client) ListLabels() (models.Labels, error) {
	req, err := cl.newRequest(http.MethodGet, "labels", nil, nil)
	if err != nil {
		return nil, err
	}

	var ls models.Labels
	if err := cl.doRequest(req, &ls); err != nil {
		return nil, err
	}

	return ls, nil
}
