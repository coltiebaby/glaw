// Please look at their suggested tips before running this.
// https://developer.riotgames.com/docs/lol#riot-games-api_tournament-api

package league

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/coltiebaby/glaw"
	"github.com/coltiebaby/glaw/league/core"
)

type ProviderRequest struct {
	Region       glaw.Region
	Registration core.TournamentProviderRegistration
}

func (c *Client) Provider(ctx context.Context, pr ProviderRequest) (id int, err error) {
	uri := `providers`
	req := NewRequest("POST", "tournament-stub", uri, pr.Region, glaw.V4)

	b, err := json.Marshal(pr.Registration)
	if err != nil {
		return id, err
	}

	buf := bytes.NewBuffer(b)
	req.Body = buf

	err = c.Do(ctx, req, &id)
	return id, err
}

type TournamentRequest struct {
	Region       glaw.Region
	Registration core.TournamentRegistration
}

func (c *Client) CreateTournament(ctx context.Context, tr TournamentRequest) (id int, err error) {
	uri := `tournaments`
	req := NewRequest("POST", "tournament-stub", uri, pr.Region, glaw.V4)

	b, err := json.Marshal(pr.Registration)
	if err != nil {
		return id, err
	}

	buf := bytes.NewBuffer(b)
	req.Body = buf

	err = c.Do(ctx, req, &id)
	return id, err
}

type TournamentCodeRequest struct {
	TournamentId  int
	CodesToCreate int // Defaults to 1
	Registration  core.TournamentCodeRegistration
}

func (c *Client) CreateTournamentCode(ctx context.Context, tr TournamentRequest) (codes []string, err error) {
	if tr.CodesToCreate == 0 {
		tr.CodesToCreate = 1
	}

	uri := `codes`
	req := NewRequest("POST", "tournament-stub", uri, tr.Region, glaw.V4)

	b, err := json.Marshal(tr.Registration)
	if err != nil {
		return codes, err
	}

	buf := bytes.NewBuffer(b)
	req.Body = buf

	err = c.Do(ctx, req, &codes)
	return codes, err
}

// Helper Function that will create 1 code for a 5 man team tournament draft game.
// Use the metadata to add information in like team vs team and so on.
func StandardCodeCreation(ctx context.Context, c *Client, tournamentId int, metadata string) ([]string, error) {
	reg := core.TournamentCodeRegistration{
		TeamSize:       5,
		MAP:            core.SUMMONERS_RIFT,
		GameType:       core.TOURNAMENT_DRAFT,
		SpecatatorType: core.ALL,
		Metadata:       metadata,
	}

	req := TournamentCodeRequest{
		TournamentId:  tournamentId,
		CodesToCreate: 1,
		Registartion:  reg,
	}

	return c.CreateTournamentCode(ctx, req)
}
