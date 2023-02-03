package todoistapi

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func newClientWithMock(t *testing.T) (*Client, *mockHttpAPI) {
	m := newMockHttpAPI(t)
	cl := &Client{token: "TODOIST_API_TOKEN", httpAPI: m}

	return cl, m
}

type mockHTTPConfig struct {
	Request  *mockHTTPConfigRequest
	Response *mockHTTPConfigResponse
}

type mockHTTPConfigRequest struct {
	URL     string
	Method  string
	Body    string
	Headers map[string]string
}
type mockHTTPConfigResponse struct {
	Status int
	Body   string
}

func (m *mockHttpAPI) mockHTTP(t *testing.T, cfg *mockHTTPConfig) {
	var b io.Reader = nil
	if cfg.Request.Body != "" {
		b = bytes.NewReader([]byte(cfg.Request.Body))
	}

	req, _ := http.NewRequest(cfg.Request.Method, cfg.Request.URL, b)
	for k, v := range cfg.Request.Headers {
		req.Header.Set(k, v)
	}

	resp := &http.Response{
		StatusCode: cfg.Response.Status,
		Body:       io.NopCloser(strings.NewReader(cfg.Response.Body)),
	}

	m.On("Do", mock.Anything).Return(resp, nil).Run(func(args mock.Arguments) {
		req := args.Get(0).(*http.Request)
		assert.Equal(t, cfg.Request.URL, req.URL.String())
		assert.Equal(t, cfg.Request.Method, req.Method)
		if cfg.Request.Body != "" {
			b, _ := req.GetBody()
			buf, _ := io.ReadAll(b)
			assert.Equal(t, cfg.Request.Body, string(buf))
		}
	})
}

func TestNew(t *testing.T) {
	cfg := &Config{Token: "TODOIST_API_TOKEN"}
	cl := New(cfg)

	assert.NotNil(t, cl)
	assert.Equal(t, cfg.Token, cl.token)
	assert.Equal(t, new(http.Client), cl.httpAPI)
}
