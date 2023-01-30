package todoistapi

import (
	"fmt"
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

func (cl *Client) GetSection(id string) (*models.Section, error) {
	req, err := cl.newRequest(http.MethodGet, fmt.Sprintf("sections/%s", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var sec models.Section
	if err := cl.doRequest(req, &sec); err != nil {
		return nil, err
	}

	return &sec, nil
}

type CreateSectionParameters struct {
	Name      string `json:"name"`
	ProjectID string `json:"project_id"`
	Order     *int   `json:"order,omitempty"`
}

func (cl *Client) CreateSection(p *CreateSectionParameters) (*models.Section, error) {
	req, err := cl.newRequest(http.MethodPost, "sections", nil, p)
	if err != nil {
		return nil, err
	}

	var sec models.Section
	if err := cl.doRequest(req, &sec); err != nil {
		return nil, err
	}

	return &sec, nil
}

type UpdateSectionParameters struct {
	Name string `json:"name"`
}

func (cl *Client) UpdateSection(id string, p *UpdateSectionParameters) (*models.Section, error) {
	req, err := cl.newRequest(http.MethodPost, fmt.Sprintf("sections/%s", id), nil, p)
	if err != nil {
		return nil, err
	}

	var sec models.Section
	if err := cl.doRequest(req, &sec); err != nil {
		return nil, err
	}

	return &sec, nil
}
