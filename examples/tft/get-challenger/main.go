package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/coltiebaby/glaw"
	"github.com/coltiebaby/glaw/api"
	"github.com/coltiebaby/glaw/tft/core"
	"github.com/coltiebaby/glaw/tft/rank"
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
	req := rank.QueueRequest {
        Tier:   core.CHALLENGER,
		Region: glaw.REGION_NA,
	}

	rank, err := client.Rank.GetQueue(ctx, req)

	if err != nil {
		log.Fatalf("%s", err)
	}

	log.Printf("%+v\n", rank)
}
