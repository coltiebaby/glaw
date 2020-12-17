// This example shows if you receive a bunch of requests at the same time what will happen
// Dev settings currently say you can only handle a burst of requests every two seconds
// We should only be able to process 20 and 10 others will error out.
// ---
// You can increase the context.WithTimeout to have all the users complete if you want.

package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/coltiebaby/glaw"
	"github.com/coltiebaby/glaw/league"
)

func user(ctx context.Context, client *league.Client, number int, start <-chan bool) {
	req := league.ChampionRotationsRequest{
		Region: glaw.REGION_NA,
	}

	<-start
	err := client.Wait(ctx, req.Region)
	if err != nil {
		log.Printf("worker %d: %s", number, err)
		return
	}

	free, err := client.ChampionRotations(ctx, req)
	if err != nil {
		log.Printf("worker %d: %s", number, err)
		return
	}

	log.Printf("worker %d: %+v\n", number, free)
}

func main() {
	opts := []glaw.Option{
		// glaw.WithRateLimiting(int, int, time.Duration)
		glaw.WithDevSettings(),
		glaw.WithAPIToken(os.Getenv("LEAGUE_API_KEY")),
	}

	client, _ := league.NewClient(opts...)

	start := make(chan bool)

	ctx, cleanup := context.WithTimeout(context.Background(), 1500*time.Millisecond)
	defer cleanup()

	for i := 1; i < 31; i++ {
		go user(ctx, client, i, start)
	}

	close(start)

	// Wait until the context times out
	<-ctx.Done()
}
