package api

import (
	"github.com/coltiebaby/glaw"
	"github.com/coltiebaby/glaw/tft"
	"github.com/coltiebaby/glaw/tft/rank"
	"github.com/coltiebaby/glaw/tft/summoner"
)

type TeamFightTactics struct {
    Rank     *rank.Client
	Summoner *summoner.Client
}

func NewTeamFightTactics(opts ...glaw.Option) (*TeamFightTactics, error) {
	client, err := tft.NewClient(opts...)
	if err != nil {
		return nil, err
	}

	tftclient := &TeamFightTactics{
		Summoner: summoner.New(client),
        Rank:     rank.New(client),
	}

	return tftclient, nil
}
