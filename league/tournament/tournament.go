// Please look at their suggested tips before running this.
// https://developer.riotgames.com/docs/lol#riot-games-api_tournament-api

package tournament

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/coltiebaby/glaw"
	"github.com/coltiebaby/glaw/league"
	"github.com/coltiebaby/glaw/league/core"
)

type Client struct {
	client  *league.Client
	Enabled bool
}

func New(c *league.Client) *Client {
	return &Client{
		client: c,
	}
}

func NewTournamentClient(opts ...glaw.Option) (*Client, error) {
	c, err := league.NewClient(opts...)
	client := &Client{
		client: c,
	}

	return client, err
}

type ProviderRequest struct {
	Region       glaw.Region
	Registration core.TournamentProviderRegistration
}

func (c *Client) GetProvider(ctx context.Context, pr ProviderRequest) (id int, err error) {
	uri := `providers`
	req := league.NewRequest("POST", "tournament-stub", uri, pr.Region, glaw.V4)

	b, err := json.Marshal(pr.Registration)
	if err != nil {
		return id, err
	}

	buf := bytes.NewBuffer(b)
	req.Body = buf

	err = c.client.Do(ctx, req, &id)
	return id, err
}

type CreateRequest struct {
	Region       glaw.Region
	Registration core.TournamentRegistration
}

func (c *Client) Create(ctx context.Context, tr CreateRequest) (id int, err error) {
	uri := `tournaments`
	req := league.NewRequest("POST", "tournament-stub", uri, tr.Region, glaw.V4)

	b, err := json.Marshal(tr.Registration)
	if err != nil {
		return id, err
	}

	buf := bytes.NewBuffer(b)
	req.Body = buf

	err = c.client.Do(ctx, req, &id)
	return id, err
}

type CodeRequest struct {
	Region        glaw.Region
	TournamentId  int
	CodesToCreate int // Defaults to 1
	Registration  core.TournamentCodeRegistration
}

func (c *Client) CreateCode(ctx context.Context, tr CodeRequest) (codes []string, err error) {
	if tr.CodesToCreate == 0 {
		tr.CodesToCreate = 1
	}

	uri := `codes`
	req := league.NewRequest("POST", "tournament-stub", uri, tr.Region, glaw.V4)

	b, err := json.Marshal(tr.Registration)
	if err != nil {
		return codes, err
	}

	buf := bytes.NewBuffer(b)
	req.Body = buf

	err = c.client.Do(ctx, req, &codes)
	return codes, err
}

// Helper Function that will create 1 code for a 5 man team tournament draft game.
// Use the metadata to add information in like team vs team and so on.
func StandardCodeCreation(ctx context.Context, c *Client, tournamentId int, metadata string) ([]string, error) {
	reg := core.TournamentCodeRegistration{
		TeamSize:      5,
		Map:           core.SUMMONERS_RIFT,
		GameType:      core.TOURNAMENT_DRAFT,
		SpectatorType: core.ALL,
		Metadata:      metadata,
	}

	req := CodeRequest{
		TournamentId:  tournamentId,
		CodesToCreate: 1,
		Registration:  reg,
	}

	return c.CreateCode(ctx, req)
}

type LobbyEventRequest struct {
	Code   string
	Region glaw.Region
}

func (c *Client) GetLobbyEvents(ctx context.Context, tr LobbyEventRequest) (events []core.TournamentEvent, err error) {
	uri := fmt.Sprintf(`lobby-events/by-code/%s`, tr.Code)
	req := league.NewRequest("GET", "tournament-stub", uri, tr.Region, glaw.V4)

	var e core.TournamentEvents
	err = c.client.Do(ctx, req, &e)

	events = e.Events
	return events, err
}
