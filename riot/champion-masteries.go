package riot

import (
    "fmt"
    "encoding/json"
    "strconv"
)

type ChampionMastery struct {
	ChampionLevel                int   `json:"championLevel"`
	ChestGranted                 bool  `json:"chestGranted"`
	ChampionPoints               int   `json:"championPoints"`
	ChampionPointsSinceLastLevel int   `json:"championPointsSinceLastLevel"`
	PlayerID                     int   `json:"playerId"`
	ChampionPointsUntilNextLevel int   `json:"championPointsUntilNextLevel"`
	TokensEarned                 int   `json:"tokensEarned"`
	ChampionID                   int   `json:"championId"`
	LastPlayTime                 int64 `json:"lastPlayTime"`
}

func GetChampionMasteries(summoner Summoner) ([]ChampionMastery) {
    masteries := make([]ChampionMastery, 0)

    uri := "champion-masteries/by-summoner/%d"
    rr := &RiotRequest {
        Type: "champion-mastery",
        Uri:  fmt.Sprintf(uri, summoner.ID),
    }

    resp := rr.GetData()
    err := json.Unmarshal(resp, &masteries)
    if err != nil {
        panic(err)
    }

    return masteries
}

func GetMasteriesByChampion(summoner Summoner, champion_id int) (ChampionMastery) {
    // /lol/champion-mastery/v3/champion-masteries/by-summoner/{summonerId}/by-champion/{championId}
    var mastery ChampionMastery

    uri := "champion-masteries/by-summoner/%d/by-champion/%d"
    rr := &RiotRequest {
        Type: "champion-mastery",
        Uri:  fmt.Sprintf(uri, summoner.ID, champion_id),
    }

    resp := rr.GetData()
    err := json.Unmarshal(resp, &mastery)
    if err != nil {
        panic(err)
    }

    return mastery
}

func GetMasteryScore(summoner Summoner) (int64) {
    // /lol/champion-mastery/v3/scores/by-summoner/{summonerId}

    uri := "scores/by-summoner/%d"
    rr := &RiotRequest {
        Type: "champion-mastery",
        Uri:  fmt.Sprintf(uri, summoner.ID),
    }

    resp := rr.GetData()
    score, err := strconv.ParseInt(string(resp), 10, 64)
    if err != nil {
        panic(err)
    }

    return score
}
