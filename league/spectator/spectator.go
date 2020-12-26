package spectator

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

func NewSpectatorClient(opts ...glaw.Option) (*Client, error) {
	c, err := league.NewClient(opts...)
	client := &Client{
		client: c,
	}

	return client, err
}

// Get Active Game
//
// get an active game based on the summoner id given. Returns a zero valued struct if nothing exist.

type ActiveGameRequest struct {
	SummonerId string
	Region     glaw.Region
}

func (c *Client) GetActiveGame(ctx context.Context, agr ActiveGameRequest) (game core.FeaturedGame, err error) {
	uri := fmt.Sprintf(`active-games/by-summoner/%s`, agr.SummonerId)
	req := league.NewRequest("GET", "spectator", uri, agr.Region, glaw.V4)

	err = c.client.Do(ctx, req, &game)
	return game, err
}

// Get Feature Games
//
// fetch a slice of featured games

type FeaturedGamesRequest struct {
	Region glaw.Region
}

func (c *Client) GetFeaturedGames(ctx context.Context, fgr FeaturedGamesRequest) (game core.FeaturedGame, err error) {
	uri := `featured-games`
	req := league.NewRequest("GET", "spectator", uri, fgr.Region, glaw.V4)

	err = c.client.Do(ctx, req, &game)
	return game, err
}
