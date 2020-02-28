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
```
