package todoistapi

import (
	"net/http"

	"github.com/koki-develop/todoist-cli/pkg/models"
)

type ListSectionsParameters struct {
	ProjectID *string `url:"project_id,omitempty"`
}

func (cl *Client) ListSections(p *ListSectionsParameters) (models.Sections, error) {
	req, err := cl.newRequest(http.MethodGet, "sections", p, nil)
	if err != nil {
		return nil, err
	}

	var secs models.Sections
	if err := cl.doRequest(req, &secs); err != nil {
		return nil, err
	}

	return secs, nil
}
