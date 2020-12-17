package glaw

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

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

// Wait will block until the either the ctx is done or a request is able to be sent
func (c *Client) Wait(ctx context.Context, region Region) error {
	if c.rl == nil {
		return nil
	}

	return c.rl.MustGet(ctx, int(region))
}

// WaitN will block for X amount of seconds before returning an error. If the context is canceled
// before the duration it will stop before then.
func (c *Client) WaitN(ctx context.Context, region Region, n time.Duration) error {
	if c.rl == nil {
		return nil
	}

	done := make(chan bool)
	timer := time.NewTimer(n)
	defer timer.Stop()

	go func() {
		<-timer.C
		close(done)
	}()

	for {
		select {
		case <-done:
			return ratelimit.EmptyErr
		default:
		}

		if err := c.rl.Get(ctx, int(region)); err != ratelimit.EmptyErr {
			break
		}
	}

	return nil
}

func (c *Client) Do(ctx context.Context, riotReq RiotRequest, to interface{}) error {
	req, err := riotReq.NewHttpRequest(ctx)
	req.Header.Add("X-Riot-Token", c.token)

	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return NewRequestError(resp)
	}

	return json.NewDecoder(resp.Body).Decode(to)
}

type Version string

const (
	V3 Version = `v3`
	V4 Version = `v4`
)
