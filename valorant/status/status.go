package status

import (
	"context"

	"github.com/coltiebaby/glaw"
	"github.com/coltiebaby/glaw/league/core"
	"github.com/coltiebaby/glaw/valorant"
)

type Client struct {
	client *valorant.Client
}

func New(c *valorant.Client) *Client {
	return &Client{
		client: c,
	}
}

func NewChampionClient(opts ...glaw.Option) (*Client, error) {
	c, err := valorant.NewClient(opts...)
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

// Uses the league of legends core.Platform
func (c *Client) Status(ctx context.Context, sr StatusRequest) (platform core.Platform, err error) {
	uri := sr.String()
	req := valorant.NewRequest("GET", "status", uri, sr.Region, glaw.V4)

	err = c.client.Do(ctx, req, &platform)
	return platform, err
}
