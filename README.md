# G-Law

A simple League of Legends api wrapper.

## Setup

Register for the [Riot Api](https://developer.games.com/) to get a developer key.
Set the environmental variable `LEAGUE_API_KEY` to your new key where ever you're going
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
	"github.com/coltiebaby/glaw/api"
	"github.com/coltiebaby/glaw/league/champion"
)

func main() {
	ctx, cleanup := context.WithTimeout(context.Background(), time.Second*3)
	defer cleanup()

	opts := []glaw.Option{
		// glaw.WithRateLimiting(int, int, time.Duration)
		glaw.WithDevSettings(),
		glaw.WithAPIToken(os.Getenv("LEAGUE_API_KEY")),
	}

	client, _ := api.NewLeagueOfLegends(opts...)
	req := champion.FreeRotationRequest{
		Region: glaw.REGION_NA,
	}

	free, err := client.Champion.GetFreeRotation(ctx, req)
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
- [v1] clash
- [v4] league
- [v4] league-exp
- [v4] match
- [v4] summoner
- [v4] status
- [v4] third-party-code
- [v4] tournament-stub
- [v4] tournament
- [v4] spectator

### Valorant
- val-content-v1
- val-match-v1
- val-status-v1



## Not Covered

### Runeterra
- account-v1
- lor-match-v1
- lor-ranked-v1
- lor-status-v1

### TFT
- tft-league-v1
- tft-match-v1
- tft-summoner-v1
