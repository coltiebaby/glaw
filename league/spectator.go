package league

import (
	"context"
	"fmt"

	"github.com/coltiebaby/glaw"
	"github.com/coltiebaby/glaw/league/core"
)

type ActiveGameRequest struct {
	SummonerId string
	Region     glaw.Region
}

func (c *Client) ActiveGame(ctx context.Context, agr ActiveGameRequest) (game core.FeaturedGame, err error) {
	uri := fmt.Sprintf(`active-games/by-summoner/%s`, agr.SummonerId)
	req := NewRequest("GET", "spectator", uri, agr.Region, glaw.V4)

	err = c.Do(ctx, req, &game)
	return game, err
}

type FeaturedGamesRequest struct {
	Region glaw.Region
}

func (c *Client) FeaturedGames(ctx context.Context, fgr FeaturedGamesRequest) (game core.FeaturedGame, err error) {
	uri := `featured-games`
	req := NewRequest("GET", "spectator", uri, fgr.Region, glaw.V4)

	err = c.Do(ctx, req, &game)
	return game, err
}
