package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/coltiebaby/glaw"
	"github.com/coltiebaby/glaw/league"
	"github.com/coltiebaby/glaw/league/core"
	"github.com/coltiebaby/glaw/league/tournament"
)

func main() {
	ctx, cleanup := context.WithTimeout(context.Background(), time.Second*3)
	defer cleanup()

	opts := []glaw.Option{
		// glaw.WithRateLimiting(int, int, time.Duration)
		glaw.WithDevSettings(),
		glaw.WithAPIToken(os.Getenv("LEAGUE_API_KEY")),
	}

	client, _ := league.NewClient(opts...)
	c := tournament.New(client)

	req := tournament.ProviderRequest{
		Region: glaw.REGION_AMERICAS,
		Registration: core.TournamentProviderRegistration{
			Url:    "http://test.com:80/callback",
			Region: glaw.Regions[glaw.REGION_NA],
		},
	}

	err := client.Wait(ctx, req.Region)
	if err != nil {
		log.Fatalf("%s", err)
	}

	providerId, err := c.GetProvider(ctx, req)
	if err != nil {
		log.Fatalf("%s", err)
	}

	log.Printf("created a new provider: %+v\n", providerId)
}
