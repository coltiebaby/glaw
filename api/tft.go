package api

import (
	"github.com/coltiebaby/glaw"
	"github.com/coltiebaby/glaw/tft"
	"github.com/coltiebaby/glaw/tft/summoner"
)

type TeamFightTactics struct {
	Summoner *summoner.Client
}

func NewTeamFightTactics(opts ...glaw.Option) (*TeamFightTactics, error) {
	client, err := tft.NewClient(opts...)
	if err != nil {
		return nil, err
	}

	tftclient := &TeamFightTactics{
		Summoner: summoner.New(client),
	}

	return tftclient, nil
}
