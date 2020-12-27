package content

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

func NewContentClient(opts ...glaw.Option) (*Client, error) {
	c, err := valorant.NewClient(opts...)
	client := &Client{
		client: c,
	}

	return client, err
}

type ContentRequest struct {
	Region glaw.Region
	Locale string
}

func (c *Client) Get(ctx context.Context, cr ContentRequest) (value core.Content, err error) {
	uri := `contents`
	if cr.Locale != "" {
		uri = fmt.Sprintf(`%s?locale="%s"`, uri, cr.Locale)
	}

	req := valorant.NewRequest("GET", "content", uri, cr.Region, glaw.V1)

	err = c.client.Do(ctx, req, &value)
	return value, err
}
