package match

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

func NewMatchClient(opts ...glaw.Option) (*Client, error) {
	c, err := tft.NewClient(opts...)
	client := &Client{
		client: c,
	}

	return client, err
}

type MatchRequest struct {
	ID     string
	Region glaw.Region
	Count  int
}

func (c *Client) Get(ctx context.Context, mr MatchRequest) (matches core.MatchStorage, err error) {
	uri := fmt.Sprintf("matches/%s", mr.ID)
	if mr.Count != 0 {
		uri = fmt.Sprintf("%s?count=\"%d\"", uri, mr.Count)
	}

	req := tft.NewRequest("GET", "match", uri, mr.Region, glaw.V1)

	err = c.client.Do(ctx, req, &matches)
	return matches, err

}

func (c *Client) GetMatchIds(ctx context.Context, mr MatchRequest) (matchIds []string, err error) {
	uri := fmt.Sprintf("matches/by-puuid/%s/ids", mr.ID)
	req := tft.NewRequest("GET", "match", uri, mr.Region, glaw.V1)

	err = c.client.Do(ctx, req, &matchIds)

	return matchIds, err
}
