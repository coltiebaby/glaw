package status

import (
	"context"

	"github.com/coltiebaby/glaw"
	"github.com/coltiebaby/glaw/league"
	"github.com/coltiebaby/glaw/league/core"
)

type Client struct {
	client *league.Client
}

func New(c *league.Client) *Client {
	return &Client{
		client: c,
	}
}

func NewChampionClient(opts ...glaw.Option) (*Client, error) {
	c, err := league.NewClient(opts...)
	client := &Client{
		client: c,
	}

	return client, err
}

type StatusRequest struct {
	Region glaw.Region
}

func (sr StatusRequest) String() string {
	return "platform-data"
}

func (c *Client) Status(ctx context.Context, sr StatusRequest) (platform core.Platform, err error) {
	uri := sr.String()
	req := league.NewRequest("GET", "status", uri, sr.Region, glaw.V4)

	err = c.client.Do(ctx, req, &platform)
	return platform, err
}
