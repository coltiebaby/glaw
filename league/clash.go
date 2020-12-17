package league

import (
	"context"
	"fmt"

	"github.com/coltiebaby/glaw"
	"github.com/coltiebaby/glaw/league/core"
)

type ClashRegistrationRequest struct {
	Region     glaw.Region
	Summoner   core.Summoner
	SummonerID string
}

func (c *Client) ClashRegistration(ctx context.Context, crr ClashRegistrationRequest) (cp []core.ClashPlayer, err error) {
	var id string
	switch {
	case crr.SummonerID != "":
		id = crr.SummonerID
	case crr.Summoner.ID != "":
		id = crr.Summoner.ID
	}

	uri := fmt.Sprintf("players/by-summoner/%s", id)
	req := NewRequest("GET", "clash", uri, crr.Region, glaw.V1)

	err = c.Do(ctx, req, &cp)
	return cp, err
}

type ClashTeamRequest struct {
	Region glaw.Region
	TeamID string
}

func (c *Client) ClashTeam(ctx context.Context, ctr ClashTeamRequest) (ct []core.ClashTeam, err error) {
	uri := `teams`
	req := NewRequest("GET", "clash", uri, ctr.Region, glaw.V1)

	err = c.Do(ctx, req, &ct)
	return ct, err
}

type ClashTournamentsRequest struct {
	Region glaw.Region
}

func (c *Client) ClashTournaments(ctx context.Context, ctr ClashTournamentsRequest) (ct []core.ClashTournament, err error) {
	uri := `tournaments`
	req := NewRequest("GET", "clash", uri, ctr.Region, glaw.V1)

	err = c.Do(ctx, req, &ct)
	return ct, err
}

type ClashTournamentRequest struct {
	Region       glaw.Region
	TeamID       string
	TournamentID string
}

func (c *Client) ClashTournament(ctx context.Context, ctr ClashTournamentRequest) (ct []core.ClashTournament, err error) {
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

	req := NewRequest("GET", "clash", uri, ctr.Region, glaw.V1)
	err = c.Do(ctx, req, &ct)
	return ct, err
}
