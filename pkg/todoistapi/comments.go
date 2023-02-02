package todoistapi

import (
	"fmt"
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

func (cl *Client) GetComment(id string) (models.Comment, error) {
	req, err := cl.newRequest(http.MethodGet, fmt.Sprintf("comments/%s", id), nil, nil)
	if err != nil {
		return nil, err
	}

	var c models.Comment
	if err := cl.doRequest(req, &c); err != nil {
		return nil, err
	}

	return c, nil
}

type CreateCommentParameters struct {
	TaskID     *string                  `json:"task_id,omitempty"`
	ProjectID  *string                  `json:"project_id,omitempty"`
	Content    *string                  `json:"content,omitempty"`
	Attachment *CreateCommentAttachment `json:"attachment,omitempty"`
}

type CreateCommentAttachment struct {
	FileName *string `json:"file_name,omitempty"`
	FileURL  *string `json:"file_url,omitempty"`
	FileType *string `json:"file_type,omitempty"`
}

func (cl *Client) CreateComment(p *CreateCommentParameters) (models.Comment, error) {
	req, err := cl.newRequest(http.MethodPost, "comments", nil, p)
	if err != nil {
		return nil, err
	}

	var c models.Comment
	if err := cl.doRequest(req, &c); err != nil {
		return nil, err
	}

	return c, nil
}
