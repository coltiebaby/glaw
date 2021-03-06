package glaw

import (
	"net/http"
	"time"

	"github.com/coltiebaby/glaw/ratelimit"
)

// Set a different HTTP client
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

// Set API token
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

// Set Rate Limiting
type RateLimitOption struct {
	value *ratelimit.RateLimiter
}

func (t RateLimitOption) apply(c *Client) (*Client, error) {
	c.rl = t.value
	return c, nil
}

func WithRateLimiting(burst, max int, dur time.Duration) RateLimitOption {
	return RateLimitOption{value: ratelimit.NewRateLimiter(burst, max, dur)}
}

func WithDevSettings() RateLimitOption {
	return WithRateLimiting(20, 100, time.Minute*2)
}
