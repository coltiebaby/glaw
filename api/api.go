package api

import (
	"github.com/coltiebaby/glaw"
	"github.com/coltiebaby/glaw/league"
	"github.com/coltiebaby/glaw/league/champion"
	"github.com/coltiebaby/glaw/league/clash"
	"github.com/coltiebaby/glaw/league/match"
	"github.com/coltiebaby/glaw/league/rank"
	"github.com/coltiebaby/glaw/league/spectator"
	"github.com/coltiebaby/glaw/league/status"
	"github.com/coltiebaby/glaw/league/summoner"
	"github.com/coltiebaby/glaw/league/thirdpartycode"
	"github.com/coltiebaby/glaw/league/tournament"
)

type LeagueOfLegends struct {
	Clash      *clash.Client
	Champion   *champion.Client
	Match      *match.Client
	Rank       *rank.Client
	Spectator  *spectator.Client
	Status     *status.Client
	Summoner   *summoner.Client
	ThirdParty *thirdpartycode.Client
	Tournament *tournament.Client
}

func NewLeagueOfLegends(opts ...glaw.Option) (*LeagueOfLegends, error) {
	client, err := league.NewClient(opts...)
	if err != nil {
		return nil, err
	}

	leagueclient := &LeagueOfLegends{
		Clash:      clash.New(client),
		Champion:   champion.New(client),
		Match:      match.New(client),
		Rank:       rank.New(client),
		Spectator:  spectator.New(client),
		Status:     status.New(client),
		Summoner:   summoner.New(client),
		ThirdParty: thirdpartycode.New(client),
		Tournament: tournament.New(client),
	}

	return leagueclient, nil
}
