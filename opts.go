package glaw

import (
	"net/http"
)

type HTTPClientOption struct {
	client *http.Client
}

func (hco HTTPClientOption) apply(c *Client) (*Client, error) {
	c.client = hco.client
	return c, nil
}

func WithHTTPClient(client *http.Client) HTTPClientOption {
	return HTTPClientOption{client: client}
}

type TokenOption struct {
	value string
}

func (t TokenOption) apply(c *Client) (*Client, error) {
	c.token = t.value
	return c, nil
}

func WithAPIToken(token string) TokenOption {
	return TokenOption{value: token}
}
