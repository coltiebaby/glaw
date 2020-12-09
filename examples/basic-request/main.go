package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/coltiebaby/glaw"
)

func main() {
	ctx, cleanup := context.WithTimeout(context.Background(), time.Second*3)
	defer cleanup()

	opts := []glaw.Option{
		// glaw.WithRateLimiting(int, int, time.Duration)
		glaw.WithDevSettings(),
		glaw.WithAPIToken(os.Getenv("LEAGUE_API_KEY")),
	}

	client, _ := glaw.NewClient(opts...)
	req := glaw.ChampionRotationsRequest{
		Region: glaw.REGION_NA,
	}

	free, err := client.ChampionRotations(ctx, req)
	if err != nil {
		log.Fatalf("%s", err)
	}
	log.Printf("%+v\n", free)
}
