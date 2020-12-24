package league

import (
	"context"
	"fmt"

	"github.com/coltiebaby/glaw"
	"github.com/coltiebaby/glaw/league/core"
)

type StatusRequest struct {
	Region glaw.Region
}

func (sr StatusRequest) String() string {
	return "platform-data"
}

func (c *Client) Status(ctx context.Context, sr StatusRequest) (platform core.Platform, err error) {
	uri := sr.String()
	req := NewRequest("GET", "status", uri, sr.Region, glaw.V4)

	err = c.Do(ctx, req, &platform)
	return platform, err
}
