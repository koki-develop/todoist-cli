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

type UpdateLabelParameters struct {
	Name       *string `json:"name,omitempty"`
	Order      *int    `json:"order,omitempty"`
	Color      *string `json:"color,omitempty"`
	IsFavorite *bool   `json:"is_favorite,omitempty"`
}

func (cl *Client) UpdateLabel(id string, p *UpdateLabelParameters) (models.Label, error) {
	req, err := cl.newRequest(http.MethodPost, fmt.Sprintf("labels/%s", id), nil, p)
	if err != nil {
		return nil, err
	}

	var l models.Label
	if err := cl.doRequest(req, &l); err != nil {
		return nil, err
	}

	return l, nil
}

func (cl *Client) DeleteLabel(id string) error {
	req, err := cl.newRequest(http.MethodDelete, fmt.Sprintf("labels/%s", id), nil, nil)
	if err != nil {
		return err
	}

	if err := cl.doRequest(req, nil); err != nil {
		return err
	}

	return nil
}

type ListSharedLabelsParameters struct {
	OmitPersonal *bool `url:"omit_personal,omitempty"`
}

func (cl *Client) ListSharedLabels(p *ListSharedLabelsParameters) (models.SharedLabels, error) {
	req, err := cl.newRequest(http.MethodGet, "labels/shared", p, nil)
	if err != nil {
		return nil, err
	}

	var ls models.SharedLabels
	if err := cl.doRequest(req, &ls); err != nil {
		return nil, err
	}

	return ls, nil
}

type RenameSharedLabelParameters struct {
	Name    *string `json:"name,omitempty"`
	NewName *string `json:"new_name,omitempty"`
}

func (cl *Client) RenameSharedLabel(p *RenameSharedLabelParameters) error {
	req, err := cl.newRequest(http.MethodPost, "labels/shared/rename", nil, p)
	if err != nil {
		return err
	}

	if err := cl.doRequest(req, nil); err != nil {
		return err
	}

	return nil
}

type RemoveSharedLabelParameters struct {
	Name *string `json:"name,omitempty"`
}

func (cl *Client) RemoveSharedLabel(p *RemoveSharedLabelParameters) error {
	req, err := cl.newRequest(http.MethodPost, "labels/shared/remove", nil, p)
	if err != nil {
		return err
	}

	if err := cl.doRequest(req, nil); err != nil {
		return err
	}

	return nil
}
