package summoner

import (
	"context"
	"fmt"

	"github.com/coltiebaby/glaw"
	"github.com/coltiebaby/glaw/tft"
	"github.com/coltiebaby/glaw/tft/core"
)

type Client struct {
	client *tft.Client
}

func New(c *tft.Client) *Client {
	return &Client{
		client: c,
	}
}

func NewSummonerClient(opts ...glaw.Option) (*Client, error) {
	c, err := tft.NewClient(opts...)
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
	var uri string

	switch sr.Type {
	case core.SummonerID:
		uri = fmt.Sprintf("summoners/%s", sr.ID)
	case core.SummonerName:
		uri = fmt.Sprintf("summoners/by-name/%s", sr.ID)
	case core.SummonerPUUID:
		uri = fmt.Sprintf("summoners/by-puuid/%s", sr.ID)
	case core.SummonerAccountID:
		uri = fmt.Sprintf("summoners/by-account/%s", sr.ID)
	}

	return uri
}

func (c *Client) Get(ctx context.Context, sr SummonerRequest) (summoner core.Summoner, err error) {
	uri := sr.String()
	req := tft.NewRequest("GET", "summoner", uri, sr.Region, glaw.V1)

	err = c.client.Do(ctx, req, &summoner)
	return summoner, err
}
