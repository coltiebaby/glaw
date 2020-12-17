package league

import (
	"context"

	"github.com/coltiebaby/glaw"
	"github.com/coltiebaby/glaw/league/core"
)

type ChampionRotationsRequest struct {
	Region glaw.Region
}

func (c *Client) ChampionRotations(ctx context.Context, fcr ChampionRotationsRequest) (ci core.ChampionInfo, err error) {
	uri := `champion-rotations`
	req := NewRequest("GET", "platform", uri, fcr.Region, glaw.V3)

	resp, err := c.Do(ctx, req)
	if err != nil {
		return ci, err
	}

	err = glaw.ProcessResponse(resp, &ci)
	return ci, err
}
