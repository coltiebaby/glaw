package match

import (
	"context"
	"fmt"

	"github.com/coltiebaby/glaw"
	"github.com/coltiebaby/glaw/valorant"
	"github.com/coltiebaby/glaw/valorant/core"
)

type Client struct {
	client *valorant.Client
}

func New(c *valorant.Client) *Client {
	return &Client{
		client: c,
	}
}

func NewMatchClient(opts ...glaw.Option) (*Client, error) {
	c, err := valorant.NewClient(opts...)
	client := &Client{
		client: c,
	}

	return client, err
}

type MatchRequest struct {
	Region glaw.Region
	Queue  core.Queue
	PUUID  string
	Id     string
}

func (c *Client) GetRecent(ctx context.Context, mr MatchRequest) (value core.RecentMatches, err error) {
	uri := fmt.Sprintf(`recent-matches/by-queue/%s`, mr.Queue)

	req := league.NewRequest("GET", "match", uri, mr.Region, glaw.V1)

	err = c.client.Do(ctx, req, &value)
	return value, err
}

func (c *Client) Get(ctx context.Context, mr MatchRequest) (value core.Match, err error) {
	uri := fmt.Sprintf(`matches/%s`, mr.Id)

	req := league.NewRequest("GET", "match", uri, mr.Region, glaw.V1)

	err = c.client.Do(ctx, req, &value)
	return value, err
}

func (c *Client) GetMatchlist(ctx context.Context, mr MatchRequest) (value core.Matchlist, err error) {
	uri := fmt.Sprintf(`matchlists/by-puuid/%s`, mr.Queue)

	req := league.NewRequest("GET", "match", uri, mr.Region, glaw.V1)

	err = c.client.Do(ctx, req, &value)
	return value, err
}
