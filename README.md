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

  "github.com/coltiebaby/g-law/riot/v4/summoner"
  "github.com/coltiebaby/g-law/riot/v4/matches"
)

func main() {
  summoner, err := summoner.ByName(`Oscillation`)
  fmt.Printf("Current Summoner: %s -- Level: %d", summoner.Name, summoner.SummonerLevel)

  match, err := matches.GetMatch(summoner.AccountID)
  // Do stuff with match...
}

```
