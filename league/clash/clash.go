package clash

import (
	"context"
	"fmt"

	"github.com/coltiebaby/glaw"
	"github.com/coltiebaby/glaw/league"
	"github.com/coltiebaby/glaw/league/core"
)

type Client struct {
	client *glaw.Client
}

func New(c *glaw.Client) *Client {
	return &Client{
		client: c,
	}
}

func NewClashClient(opts ...glaw.Option) (*Client, error) {
	c, err := glaw.NewClient(opts...)
	client := &Client{
		client: c,
	}

	return client, err
}

// Get Registration
//
// Returns a slice of information if a summoner is registered for a clash day(s)

type RegistrationRequest struct {
	Region     glaw.Region
	Summoner   core.Summoner
	SummonerID string
}

func (c *Client) GetRegistration(ctx context.Context, crr RegistrationRequest) (cp []core.ClashPlayer, err error) {
	var id string
	switch {
	case crr.SummonerID != "":
		id = crr.SummonerID
	case crr.Summoner.ID != "":
		id = crr.Summoner.ID
	}

	uri := fmt.Sprintf("players/by-summoner/%s", id)
	req := league.NewRequest("GET", "clash", uri, crr.Region, glaw.V1)

	err = c.Do(ctx, req, &cp)
	return cp, err
}

// Get Team
//
// Returns information based on a clash team

type TeamRequest struct {
	Region glaw.Region
	TeamID string
}

func (c *Client) GetTeam(ctx context.Context, ctr TeamRequest) (ct []core.ClashTeam, err error) {
	uri := `teams`
	req := league.NewRequest("GET", "clash", uri, ctr.Region, glaw.V1)

	err = c.Do(ctx, req, &ct)
	return ct, err
}

// Get Tournaments
//
// Get some clash tournaments info

type TournamentsRequest struct {
	Region glaw.Region
}

func (c *Client) GetTournaments(ctx context.Context, ctr TournamentsRequest) (ct []core.ClashTournament, err error) {
	uri := `tournaments`
	req := league.NewRequest("GET", "clash", uri, ctr.Region, glaw.V1)

	err = c.Do(ctx, req, &ct)
	return ct, err
}

// Get Tournaments
//
// Get some clash tournament info

type TournamentRequest struct {
	Region       glaw.Region
	TeamID       string
	TournamentID string
}

func (c *Client) GetTournament(ctx context.Context, ctr TournamentRequest) (ct []core.ClashTournament, err error) {
	var id string
	uri := `tournaments`

	switch {
	case ctr.TeamID != "":
		id = ctr.TeamID
		uri = fmt.Sprintf("%s/by-team/%s", uri, id)
	case ctr.TournamentID != "":
		id = ctr.TournamentID
		uri = fmt.Sprintf("%s/%s", uri, id)
	}

	req := league.NewRequest("GET", "clash", uri, ctr.Region, glaw.V1)
	err = c.Do(ctx, req, &ct)
	return ct, err
}
