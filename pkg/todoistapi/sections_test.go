package todoistapi

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/koki-develop/todoist-cli/pkg/models"
	"github.com/koki-develop/todoist-cli/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestClient_ListSections(t *testing.T) {
	tests := []struct {
		p           *ListSectionsParameters
		expectedURL string
		resp        string
		status      int
		want        models.Sections
		wantErr     bool
	}{
		{
			p:           &ListSectionsParameters{},
			expectedURL: "https://api.todoist.com/rest/v2/sections",
			resp:        `[{"id": "1", "name": "SECTION_1"}, {"id": "2", "name": "SECTION_2"}, {"id": "3", "name": "SECTION_3"}]`,
			status:      http.StatusOK,
			want:        models.Sections{{"id": "1", "name": "SECTION_1"}, {"id": "2", "name": "SECTION_2"}, {"id": "3", "name": "SECTION_3"}},
			wantErr:     false,
		},
		{
			p:           &ListSectionsParameters{ProjectID: util.Ptr("PROJECT_ID")},
			expectedURL: "https://api.todoist.com/rest/v2/sections?project_id=PROJECT_ID",
			resp:        `[{"id": "1", "name": "SECTION_1"}, {"id": "2", "name": "SECTION_2"}, {"id": "3", "name": "SECTION_3"}]`,
			status:      http.StatusOK,
			want:        models.Sections{{"id": "1", "name": "SECTION_1"}, {"id": "2", "name": "SECTION_2"}, {"id": "3", "name": "SECTION_3"}},
			wantErr:     false,
		},
		{
			p:           &ListSectionsParameters{},
			expectedURL: "https://api.todoist.com/rest/v2/sections",
			resp:        "ERROR_RESPONSE",
			status:      http.StatusBadRequest,
			want:        nil,
			wantErr:     true,
		},
		{
			p:           &ListSectionsParameters{},
			expectedURL: "https://api.todoist.com/rest/v2/sections",
			resp:        "ERROR_RESPONSE",
			status:      http.StatusInternalServerError,
			want:        nil,
			wantErr:     true,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			cl, m := newClientWithMock(t)

			m.mockHTTP(t, &mockHTTPConfig{
				Request: &mockHTTPConfigRequest{
					URL:    tt.expectedURL,
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

			got, err := cl.ListSections(tt.p)
			if tt.wantErr {
				assert.EqualError(t, err, tt.resp)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestClient_GetSection(t *testing.T) {
	tests := []struct {
		id      string
		resp    string
		status  int
		want    models.Section
		wantErr bool
	}{
		{
			id:      "1",
			resp:    `{"id": "1", "name": "SECTION"}`,
			status:  http.StatusOK,
			want:    models.Section{"id": "1", "name": "SECTION"},
			wantErr: false,
		},
		{
			id:      "1",
			resp:    "ERROR_RESPONSE",
			status:  http.StatusBadRequest,
			want:    nil,
			wantErr: true,
		},
		{
			id:      "1",
			resp:    "ERROR_RESPONSE",
			status:  http.StatusInternalServerError,
			want:    nil,
			wantErr: true,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			cl, m := newClientWithMock(t)

			m.mockHTTP(t, &mockHTTPConfig{
				Request: &mockHTTPConfigRequest{
					URL:    fmt.Sprintf("https://api.todoist.com/rest/v2/sections/%s", tt.id),
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

			got, err := cl.GetSection(tt.id)
			if tt.wantErr {
				assert.EqualError(t, err, tt.resp)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestClient_CreateSection(t *testing.T) {
	tests := []struct {
		p           *CreateSectionParameters
		expectedReq string
		resp        string
		status      int
		want        models.Section
		wantErr     bool
	}{
		{
			p:           &CreateSectionParameters{Name: util.Ptr("SECTION")},
			expectedReq: `{"name":"SECTION"}`,
			resp:        `{"id": "1", "name": "SECTION"}`,
			status:      http.StatusCreated,
			want:        models.Section{"id": "1", "name": "SECTION"},
			wantErr:     false,
		},
		{
			p: &CreateSectionParameters{
				Name:      util.Ptr("SECTION"),
				ProjectID: util.Ptr("PROJECT_ID"),
				Order:     util.Ptr(0),
			},
			expectedReq: `{"name":"SECTION","project_id":"PROJECT_ID","order":0}`,
			resp:        `{"id": "1", "name": "SECTION"}`,
			status:      http.StatusCreated,
			want:        models.Section{"id": "1", "name": "SECTION"},
			wantErr:     false,
		},
		{
			p:           &CreateSectionParameters{Name: util.Ptr("SECTION")},
			expectedReq: `{"name":"SECTION"}`,
			resp:        "ERROR_RESPONSE",
			status:      http.StatusBadRequest,
			want:        nil,
			wantErr:     true,
		},
		{
			p:           &CreateSectionParameters{Name: util.Ptr("SECTION")},
			expectedReq: `{"name":"SECTION"}`,
			resp:        "ERROR_RESPONSE",
			status:      http.StatusInternalServerError,
			want:        nil,
			wantErr:     true,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			cl, m := newClientWithMock(t)

			m.mockHTTP(t, &mockHTTPConfig{
				Request: &mockHTTPConfigRequest{
					URL:    "https://api.todoist.com/rest/v2/sections",
					Method: http.MethodPost,
					Body:   tt.expectedReq,
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

			got, err := cl.CreateSection(tt.p)
			if tt.wantErr {
				assert.EqualError(t, err, tt.resp)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestClient_UpdateSection(t *testing.T) {
	tests := []struct {
		id          string
		p           *UpdateSectionParameters
		expectedReq string
		resp        string
		status      int
		want        models.Section
		wantErr     bool
	}{
		{
			id:          "1",
			p:           &UpdateSectionParameters{Name: util.Ptr("SECTION")},
			expectedReq: `{"name":"SECTION"}`,
			resp:        `{"id": "1", "name": "SECTION"}`,
			status:      http.StatusOK,
			want:        models.Section{"id": "1", "name": "SECTION"},
			wantErr:     false,
		},
		{
			id:          "1",
			p:           &UpdateSectionParameters{Name: util.Ptr("SECTION")},
			expectedReq: `{"name":"SECTION"}`,
			resp:        "ERROR_RESPONSE",
			status:      http.StatusBadRequest,
			want:        nil,
			wantErr:     true,
		},
		{
			id:          "1",
			p:           &UpdateSectionParameters{Name: util.Ptr("SECTION")},
			expectedReq: `{"name":"SECTION"}`,
			resp:        "ERROR_RESPONSE",
			status:      http.StatusInternalServerError,
			want:        nil,
			wantErr:     true,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			cl, m := newClientWithMock(t)

			m.mockHTTP(t, &mockHTTPConfig{
				Request: &mockHTTPConfigRequest{
					URL:    fmt.Sprintf("https://api.todoist.com/rest/v2/sections/%s", tt.id),
					Method: http.MethodPost,
					Body:   tt.expectedReq,
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

			got, err := cl.UpdateSection(tt.id, tt.p)
			if tt.wantErr {
				assert.EqualError(t, err, tt.resp)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestClient_DeleteSection(t *testing.T) {
	tests := []struct {
		id      string
		resp    string
		status  int
		wantErr bool
	}{
		{
			id:      "1",
			status:  http.StatusNoContent,
			wantErr: false,
		},
		{
			id:      "1",
			resp:    "ERROR_RESPONSE",
			status:  http.StatusBadRequest,
			wantErr: true,
		},
		{
			id:      "1",
			resp:    "ERROR_RESPONSE",
			status:  http.StatusInternalServerError,
			wantErr: true,
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			cl, m := newClientWithMock(t)

			m.mockHTTP(t, &mockHTTPConfig{
				Request: &mockHTTPConfigRequest{
					URL:    fmt.Sprintf("https://api.todoist.com/rest/v2/sections/%s", tt.id),
					Method: http.MethodDelete,
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

			err := cl.DeleteSection(tt.id)
			if tt.wantErr {
				assert.EqualError(t, err, tt.resp)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
