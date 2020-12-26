package match

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

func NewMatchClient(opts ...glaw.Option) (*Client, error) {
	c, err := league.NewClient(opts...)
	client := &Client{
		client: c,
	}

	return client, err
}

type MatchRequest struct {
	ID     string
	Region glaw.Region
}

func (mr MatchRequest) String() string {
	return fmt.Sprintf("matches/%s", mr.ID)
}

func (c *Client) Get(ctx context.Context, mr MatchRequest) (matches core.MatchStorage, err error) {
	uri := mr.String()
	req := league.NewRequest("GET", "match", uri, mr.Region, glaw.V4)

	err = c.client.Do(ctx, req, &matches)
	return matches, err

}

type MatchesRequest struct {
	AccountID string
	Region    glaw.Region
}

func (mr MatchesRequest) String() string {
	return fmt.Sprintf("matchlists/by-account/%s", mr.AccountID)
}

func (c *Client) GetAll(ctx context.Context, mr MatchRequest) (matches core.MatchStorage, err error) {
	uri := mr.String()
	req := league.NewRequest("GET", "match", uri, mr.Region, glaw.V4)

	err = c.client.Do(ctx, req, &matches)
	return matches, err
}

type TimelineRequest struct {
	ID     string
	Region glaw.Region
}

func (tr TimelineRequest) String() string {
	return fmt.Sprintf("timelines/by-match/%s", tr.ID)
}

func (c *Client) GetTimeline(ctx context.Context, mr MatchRequest) (matches core.MatchStorage, err error) {
	uri := mr.String()
	req := league.NewRequest("GET", "match", uri, mr.Region, glaw.V4)

	err = c.client.Do(ctx, req, &matches)
	return matches, err
}
