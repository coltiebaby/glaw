package glaw

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/coltiebaby/glaw/ratelimit"
)

type Client struct {
	client *http.Client
	rl     *ratelimit.RateLimiter
	token  string
}

type Option interface {
	apply(*Client) (*Client, error)
}

func NewClient(opts ...Option) (c *Client, err error) {
	c = &Client{
		client: http.DefaultClient,
	}

	for _, opt := range opts {
		if c, err := opt.apply(c); err != nil {
			return c, err
		}
	}

	return c, err
}

func (c *Client) Do(req *http.Request) (resp *http.Response, err error) {
	req.Header.Add("X-Riot-Token", c.token)

	resp, err = c.client.Do(req)
	if err != nil {
		return resp, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		err = NewRequestError(resp)
	}

	return resp, err
}

func ProcessResponse(resp *http.Response, to interface{}) error {
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(to)
}

type Request struct {
	Method  string
	Domain  string
	Version Version
	Region  Region
	Uri     string
	Body    io.Reader
}

func (r Request) URL() string {
	return fmt.Sprintf(partial, r.Region.Base(), r.Domain, r.Version, r.Uri)
}

func (r Request) NewHttpRequest(ctx context.Context) (*http.Request, error) {
	// return http.NewRequest(r.Method, r.URL(), r.Body)
	return http.NewRequestWithContext(ctx, r.Method, r.URL(), r.Body)
}

// func (r Request) NewHttpRequest() (*http.Request, error) {
// 	ctx := context.Background()
//
// 	return r.NewHttpRequestWithCtx(ctx)
// }

type Version string

const (
	V3 Version = `v3`
	V4 Version = `v4`
)

const partial = "https://%s/lol/%s/%s/%s"
