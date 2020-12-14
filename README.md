# G-Law

A simple League of Legends api wrapper.

## Setup

Register for the [Riot Api](https://developer.games.com/) to get a developer key.
Set the environmental variable `RIOT_API_TOKEN` to your new key where ever you're going
to run your stuff.

## Example
```go
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
```

## Covered

### General

### League Of Legends
- [v4] champion-mastery
- [v3] champion (free-champion-rotation)
- [v4] league
- [v4] match
- [v4] summoner


## Not Covered

### League of Legends
- clash-v1
- tournament-stub-v4
- tournament-v4
- third-party-code-v4
- spectator-v4
- lol-status-v4
- league-exp-v4

### Runeterra
- account-v1
- lor-match-v1
- lor-ranked-v1
- lor-status-v1

### TFT
- tft-league-v1
- tft-match-v1
- tft-summoner-v1

### Valorant
- val-content-v1
- val-match-v1
- val-status-v1
