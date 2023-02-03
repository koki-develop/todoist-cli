package todoistapi

import "net/http"

type httpAPI interface {
	Do(req *http.Request) (*http.Response, error)
}
