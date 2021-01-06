package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/coltiebaby/glaw"
	"github.com/coltiebaby/glaw/api"
	"github.com/coltiebaby/glaw/tft/core"
	"github.com/coltiebaby/glaw/tft/match"
	"github.com/coltiebaby/glaw/tft/summoner"
)

func main() {
	ctx, cleanup := context.WithTimeout(context.Background(), time.Second*5)
	defer cleanup()

	opts := []glaw.Option{
		// glaw.WithRateLimiting(int, int, time.Duration)
		glaw.WithDevSettings(),
		glaw.WithAPIToken(os.Getenv("LEAGUE_API_KEY")),
	}

	client, _ := api.NewTeamFightTactics(opts...)
	req := summoner.SummonerRequest{
		Type:   core.SummonerName,
		ID:     `SagePlaysYi`,
		Region: glaw.REGION_NA,
	}

	summoner, err := client.Summoner.Get(ctx, req)

	if err != nil {
		log.Fatalf("failed to get summoner %s", err)
	}

	matchreq := match.MatchRequest{
		ID:     summoner.PUUID,
		Region: glaw.REGION_AMERICAS,
	}

	matchIds, err := client.Match.GetMatchIds(ctx, matchreq)
	if err != nil {
		log.Fatalf("%s", err)
	}

	if len(matchIds) == 0 {
		log.Print("there were no matches")
		return
	}

	matchreq.ID = matchIds[0]
	match, err := client.Match.Get(ctx, matchreq)
	if err != nil {
		log.Fatalf("%s", err)
	}

	log.Printf("%+v\n", match)
}
