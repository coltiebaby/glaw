package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/coltiebaby/glaw"
	"github.com/coltiebaby/glaw/api"
	"github.com/coltiebaby/glaw/tft/core"
	"github.com/coltiebaby/glaw/tft/summoner"
)

func main() {
	ctx, cleanup := context.WithTimeout(context.Background(), time.Second*3)
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
		log.Fatalf("%s", err)
	}

	log.Printf("%+v\n", summoner)
}
