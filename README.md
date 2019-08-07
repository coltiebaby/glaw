# G-Law

A simple League of Legends api wrapper.

## Setup

Register for the [Riot Api](https://developer.riotgames.com/) to get a developer key.
Set the environmental variable `RIOT_API_TOKEN` to your new key where ever you're going
to run your stuff.

## Example
```go
package main

import (
	"fmt"

	"github.com/coltiebaby/g-law/riot"
	"github.com/coltiebaby/g-law/riot/v4/summoner"
)

func main() {
	summoner, err := summoner.ByName(riot.Client, `Oscillation`)
	if err != nil {
		fmt.Println("Got an err:", err)
		return
	}

	fmt.Printf("Current Summoner: %s -- Level: %d\n", summoner.Name, summoner.SummonerLevel)
}
```
