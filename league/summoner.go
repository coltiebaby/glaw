package league

import (
	"context"
	"fmt"

	"github.com/coltiebaby/glaw"
	"github.com/coltiebaby/glaw/league/core"
)

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

func (c *Client) Summoner(ctx context.Context, sr SummonerRequest) (summoner core.Summoner, err error) {
	req := glaw.Request{
		Method:  `GET`,
		Domain:  `summoner`,
		Version: glaw.V4,
		Region:  sr.Region,
		Uri:     sr.String(),
	}

	resp, err := c.Do(ctx, req)
	if err != nil {
		return summoner, err
	}

	err = glaw.ProcessResponse(resp, &summoner)
	return summoner, err
}
