package thirdpartycode

import (
	"context"
	"fmt"

	"github.com/coltiebaby/glaw"
	"github.com/coltiebaby/glaw/league"
)

type Client struct {
	client *league.Client
}

func New(c *league.Client) *Client {
	return &Client{
		client: c,
	}
}

func NewThirdPartyClient(opts ...glaw.Option) (*Client, error) {
	c, err := league.NewClient(opts...)
	client := &Client{
		client: c,
	}

	return client, err
}

type CodeRequest struct {
	SummonerId string
	Region     glaw.Region
}

func (c *Client) Get(ctx context.Context, cr CodeRequest) (code string, err error) {
	uri := fmt.Sprintf(`third-party-code/by-summoner/%s`, cr.SummonerId)
	req := league.NewRequest("GET", "platform", uri, cr.Region, glaw.V4)

	err = c.client.Do(ctx, req, &code)
	return code, err
}
