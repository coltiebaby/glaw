package glaw

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/coltiebaby/glaw/ratelimit"
)

type RiotRequest interface {
	GetRegion() Region
	NewHttpRequest(context.Context) (*http.Request, error)
}

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

func (c *Client) Verify(ctx context.Context, region Region) error {
	if c.rl == nil {
		return nil
	}

	return c.rl.MustGet(ctx, int(region))
}

func (c *Client) Do(ctx context.Context, riotReq RiotRequest) (resp *http.Response, err error) {
	err = c.Verify(ctx, riotReq.GetRegion())
	if err != nil {
		return nil, err
	}

	req, err := riotReq.NewHttpRequest(ctx)
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

type Version string

const (
	V3 Version = `v3`
	V4 Version = `v4`
)
