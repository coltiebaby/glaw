package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/coltiebaby/glaw"
)

func main() {
	opts := []glaw.Option{
		glaw.WithRateLimiting(),
		glaw.WithAPIToken(os.Getenv("LEAGUE_API_KEY")),
	}

	client, _ := glaw.NewClient(opts...)
	req := glaw.ChampionRotationsRequest{
		Region: glaw.REGION_NA,
	}

	for i := 0; i < 30; i++ {
		ctx, cleanup := context.WithTimeout(context.Background(), time.Second*3)
		free, err := client.ChampionRotations(ctx, req)
		cleanup()

		if err != nil {
			log.Printf("%s", err)
			time.Sleep(time.Second)
		}
		log.Printf("%+v\n", free)
	}
}
