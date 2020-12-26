package summoner

import (
	"context"
	"fmt"

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

func NewSummonerClient(opts ...glaw.Option) (*Client, error) {
	c, err := league.NewClient(opts...)
	client := &Client{
		client: c,
	}

	return client, err
}

type SummonerRequest struct {
	Type   int
	ID     string
	Region glaw.Region
}

func (sr SummonerRequest) String() string {
	uri := sr.ID

	switch sr.Type {
	case core.SummonerID:
	case core.SummonerName:
		uri = fmt.Sprintf("by-name/%s", uri)
	case core.SummonerPUUID:
		uri = fmt.Sprintf("by-puuid/%s", uri)
	case core.SummonerAccountID:
		uri = fmt.Sprintf("by-account/%s", uri)
	}

	return uri
}

func (c *Client) Get(ctx context.Context, sr SummonerRequest) (summoner core.Summoner, err error) {
	uri := sr.String()
	req := league.NewRequest("GET", "summoner", uri, sr.Region, glaw.V4)

	err = c.client.Do(ctx, req, &summoner)
	return summoner, err
}
