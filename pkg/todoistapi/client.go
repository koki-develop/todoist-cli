package todoistapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/google/go-querystring/query"
)

type Client struct {
	token   string
	httpAPI httpAPI
}

type Config struct {
	Token string
}

func New(cfg *Config) *Client {
	return &Client{
		token:   cfg.Token,
		httpAPI: new(http.Client),
	}
}

func (cl *Client) newRequest(method, pathname string, params, body interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: "https", Path: "api.todoist.com/rest/v2"}
	u.Path = path.Join(u.Path, pathname)
	if params != nil {
		v, err := query.Values(params)
		if err != nil {
			return nil, err
		}
		u.RawQuery = v.Encode()
	}

	var p io.Reader = nil
	if body != nil {
		j, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		p = bytes.NewReader(j)
	}

	req, err := http.NewRequest(method, u.String(), p)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cl.token))
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (cl *Client) doRequest(req *http.Request, dst interface{}) error {
	resp, err := cl.httpAPI.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if !(200 <= resp.StatusCode && resp.StatusCode < 300) {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(b))
	}

	if dst != nil {
		if err := json.NewDecoder(resp.Body).Decode(dst); err != nil {
			return err
		}
	}

	return nil
}
