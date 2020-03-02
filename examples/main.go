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

	client, _ := glaw.NewClient(glaw.WithAPIToken(os.Getenv("LEAGUE_API_KEY")))
	req := glaw.ChampionRotationsRequest{
		Region: glaw.REGION_NA,
	}

	free, err := client.ChampionRotations(ctx, req)
	if err != nil {
		log.Fatalf("%s", err)
	}
	log.Printf("%+v\n", free)
}
