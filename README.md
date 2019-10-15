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
	"fmt"

	"github.com/coltiebaby/glaw"
	"github.com/coltiebaby/glaw/v4/summoner"
)

func main() {
	summoner, err := summoner.ByName(glaw.Client, `Oscillation`)
	if err != nil {
		fmt.Println("Got an err:", err)
		return
	}

	fmt.Printf("Current Summoner: %s -- Level: %d\n", summoner.Name, summoner.SummonerLevel)
}
```
