package todoistapi

import (
	"fmt"
	"net/http"

	"github.com/koki-develop/todoist-cli/pkg/models"
)

func (cl *Client) ListProjects() (models.Projects, error) {
	req, err := cl.newRequest(http.MethodGet, "projects", nil, nil)
	if err != nil {
		return nil, err
	}

	var projs models.Projects
	if err := cl.doRequest(req, &projs); err != nil {
		return nil, err
	}

	return projs, nil
}

func (cl *Client) GetProject(id string) (models.Project, error) {
	req, err := cl.newRequest(http.MethodGet, fmt.Sprintf("projects/%s", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var proj models.Project
	if err := cl.doRequest(req, &proj); err != nil {
		return nil, err
	}

	return proj, nil
}

type CreateProjectParameters struct {
	Name       *string `json:"name,omitempty"`
	ParentID   *string `json:"parent_id,omitempty"`
	Color      *string `json:"color,omitempty"`
	IsFavorite *bool   `json:"is_favorite,omitempty"`
	ViewStyle  *string `json:"view_style,omitempty"`
}

func (cl *Client) CreateProject(p *CreateProjectParameters) (models.Project, error) {
	req, err := cl.newRequest(http.MethodPost, "projects", nil, p)
	if err != nil {
		return nil, err
	}

	var proj models.Project
	if err := cl.doRequest(req, &proj); err != nil {
		return nil, err
	}

	return proj, nil
}

type UpdateProjectParameters struct {
	Name       *string `json:"name,omitempty"`
	Color      *string `json:"color,omitempty"`
	IsFavorite *bool   `json:"is_favorite,omitempty"`
	ViewStyle  *string `json:"view_style,omitempty"`
}

func (cl *Client) UpdateProject(id string, p *UpdateProjectParameters) (models.Project, error) {
	req, err := cl.newRequest(http.MethodPost, fmt.Sprintf("projects/%s", id), nil, p)
	if err != nil {
		return nil, err
	}

	var proj models.Project
	if err := cl.doRequest(req, &proj); err != nil {
		return nil, err
	}

	return proj, nil
}

func (cl *Client) DeleteProject(id string) error {
	req, err := cl.newRequest(http.MethodDelete, fmt.Sprintf("projects/%s", id), nil, nil)
	if err != nil {
		return err
	}

	if err := cl.doRequest(req, nil); err != nil {
		return err
	}

	return nil
}
