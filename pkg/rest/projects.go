package rest

import (
	"fmt"
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

func (cl *Client) GetProject(id string) (*models.Project, error) {
	req, err := cl.newRequest(http.MethodGet, fmt.Sprintf("projects/%s", id), nil)
	if err != nil {
		return nil, err
	}

	var proj models.Project
	if err := cl.doRequest(req, &proj); err != nil {
		return nil, err
	}

	return &proj, nil
}

type CreateProjectPayload struct {
	Name       string  `json:"name"`
	ParentID   *string `json:"parent_id"`
	Color      *string `json:"color"`
	IsFavorite *bool   `json:"is_favorite"`
	ViewStyle  *string `json:"view_style"`
}

func (cl *Client) CreateProject(p *CreateProjectPayload) (*models.Project, error) {
	req, err := cl.newRequest(http.MethodPost, "projects", p)
	if err != nil {
		return nil, err
	}

	var proj models.Project
	if err := cl.doRequest(req, &proj); err != nil {
		return nil, err
	}

	return &proj, nil
}
