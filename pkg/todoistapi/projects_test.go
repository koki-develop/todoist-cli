package todoistapi

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/koki-develop/todoist-cli/pkg/models"
	"github.com/koki-develop/todoist-cli/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestClient_ListProjects(t *testing.T) {
	tests := []struct {
		resp    string
		status  int
		want    models.Projects
		wantErr bool
	}{
		{
			resp:   `[{"id": "1", "name": "PROJECT_1"}, {"id": "2", "name": "PROJECT_2"}, {"id": "3", "name": "PROJECT_3"}]`,
			status: 200,
			want: models.Projects{
				{"id": "1", "name": "PROJECT_1"},
				{"id": "2", "name": "PROJECT_2"},
				{"id": "3", "name": "PROJECT_3"},
			},
			wantErr: false,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			cl, m := newClientWithMock(t)

			m.mockHTTP(t, &mockHTTPConfig{
				Request: &mockHTTPConfigRequest{
					URL:    "https://api.todoist.com/rest/v2/projects",
					Method: http.MethodGet,
					Headers: map[string]string{
						"Authorization": "Bearer TODOIST_API_TOKEN",
						"Content-Type":  "application/json",
					},
				},
				Response: &mockHTTPConfigResponse{
					Status: 200,
					Body:   tt.resp,
				},
			})

			got, err := cl.ListProjects()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestClient_GetProject(t *testing.T) {
	tests := []struct {
		id      string
		resp    string
		status  int
		want    models.Project
		wantErr bool
	}{
		{
			id:      "1",
			resp:    `{"id": "1", "name": "PROJECT"}`,
			status:  200,
			want:    models.Project{"id": "1", "name": "PROJECT"},
			wantErr: false,
		},
		{
			id:      "1",
			resp:    "ERROR_RESPONSE",
			status:  400,
			want:    nil,
			wantErr: true,
		},
		{
			id:      "1",
			resp:    "ERROR_RESPONSE",
			status:  500,
			want:    nil,
			wantErr: true,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			cl, m := newClientWithMock(t)

			m.mockHTTP(t, &mockHTTPConfig{
				Request: &mockHTTPConfigRequest{
					URL:    fmt.Sprintf("https://api.todoist.com/rest/v2/projects/%s", tt.id),
					Method: http.MethodGet,
					Headers: map[string]string{
						"Authorization": "Bearer TODOIST_API_TOKEN",
						"Content-Type":  "application/json",
					},
				},
				Response: &mockHTTPConfigResponse{
					Status: tt.status,
					Body:   tt.resp,
				},
			})

			proj, err := cl.GetProject(tt.id)
			if tt.wantErr {
				assert.EqualError(t, err, tt.resp)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, proj)
		})
	}
}

func TestClient_CreateProject(t *testing.T) {
	tests := []struct {
		p       *CreateProjectParameters
		req     string
		resp    string
		status  int
		want    models.Project
		wantErr bool
	}{
		{
			p:       &CreateProjectParameters{Name: util.Ptr("PROJECT")},
			req:     `{"name":"PROJECT"}`,
			resp:    `{"id": "1", "name": "PROJECT"}`,
			status:  201,
			want:    models.Project{"id": "1", "name": "PROJECT"},
			wantErr: false,
		},
		{
			p: &CreateProjectParameters{
				Name:       util.Ptr("PROJECT"),
				ParentID:   util.Ptr("PARENT_ID"),
				Color:      util.Ptr("COLOR"),
				IsFavorite: util.Ptr(true),
				ViewStyle:  util.Ptr("VIEW_STYLE"),
			},
			req:     `{"name":"PROJECT","parent_id":"PARENT_ID","color":"COLOR","is_favorite":true,"view_style":"VIEW_STYLE"}`,
			resp:    `{"id": "1", "name": "PROJECT"}`,
			status:  201,
			want:    models.Project{"id": "1", "name": "PROJECT"},
			wantErr: false,
		},
		{
			p:       &CreateProjectParameters{Name: util.Ptr("PROJECT")},
			req:     `{"name":"PROJECT"}`,
			resp:    "ERROR_RESPONSE",
			status:  400,
			want:    nil,
			wantErr: true,
		},
		{
			p:       &CreateProjectParameters{Name: util.Ptr("PROJECT")},
			req:     `{"name":"PROJECT"}`,
			resp:    "ERROR_RESPONSE",
			status:  500,
			want:    nil,
			wantErr: true,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			cl, m := newClientWithMock(t)

			m.mockHTTP(t, &mockHTTPConfig{
				Request: &mockHTTPConfigRequest{
					URL:    "https://api.todoist.com/rest/v2/projects",
					Method: http.MethodPost,
					Body:   tt.req,
					Headers: map[string]string{
						"Authorization": "Bearer TODOIST_API_TOKEN",
						"Content-Type":  "application/json",
					},
				},
				Response: &mockHTTPConfigResponse{Status: tt.status, Body: tt.resp},
			})

			got, err := cl.CreateProject(tt.p)
			if tt.wantErr {
				assert.EqualError(t, err, tt.resp)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
