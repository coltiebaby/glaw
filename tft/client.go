package tft

import (
	"github.com/coltiebaby/glaw"
)

type Client struct {
	*glaw.Client
}

func NewClient(opts ...glaw.Option) (*Client, error) {
	client, err := glaw.NewClient(opts...)
	if err != nil {
		return nil, err
	}

	return &Client{
		Client: client,
	}, nil
}
