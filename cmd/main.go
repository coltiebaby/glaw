package main

import (
	"fmt"

	"github.com/coltiebaby/glaw/riot"
	"github.com/coltiebaby/glaw/riot/v4/summoner"
)

func main() {
	summoner, err := summoner.ByName(riot.Client, `Oscillation`)
	if err != nil {
		fmt.Println("Got an err:", err.Error())
		return
	}

	fmt.Printf("Current Summoner: %s -- Level: %d\n", summoner.Name, summoner.SummonerLevel)
}
