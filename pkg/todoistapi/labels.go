package todoistapi

import (
	"fmt"
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

func (cl *Client) GetLabel(id string) (models.Label, error) {
	req, err := cl.newRequest(http.MethodGet, fmt.Sprintf("labels/%s", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var l models.Label
	if err := cl.doRequest(req, &l); err != nil {
		return nil, err
	}

	return l, nil
}

type CreateLabelParameters struct {
	Name       *string `json:"name,omitempty"`
	Order      *int    `json:"order,omitempty"`
	Color      *string `json:"color,omitempty"`
	IsFavorite *bool   `json:"is_favorite,omitempty"`
}

func (cl *Client) CreateLabel(p *CreateLabelParameters) (models.Label, error) {
	req, err := cl.newRequest(http.MethodPost, "labels", nil, p)
	if err != nil {
		return nil, err
	}

	var l models.Label
	if err := cl.doRequest(req, &l); err != nil {
		return nil, err
	}

	return l, nil
}
