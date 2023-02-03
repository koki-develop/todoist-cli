package todoistapi

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	cfg := &Config{Token: "TODOIST_API_TOKEN"}
	cl := New(cfg)

	assert.NotNil(t, cl)
	assert.Equal(t, cfg.Token, cl.token)
	assert.Equal(t, new(http.Client), cl.httpAPI)
}
