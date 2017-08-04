package mastery

import (
    "fmt"
    "log"
    "vs/riot"
)

var mastery = fmt.Sprintf("/champion-mastery/%s/", riot.Version)

func MasteryAllChampions() ([]byte, error) {
    // All Champion Mastery
    // Ask Riot for all the summoners champions based on mastery score.

    uri := mastery + "champion-masteries/by-summoner/%s"
    uri = fmt.Sprintf(uri, "28747969")

    log.Println("Requesting MasteryAllChampions...")
    return riot.GetData("GET", uri)
}

func MasteryGetChampion() ([]byte, error) {
    // Get Champion Score
    // Find summoner score for a single champion

    uri := mastery + "champion-masteries/by-summoner/%s/by-champion/%s"
    uri = fmt.Sprintf(uri, "28747969", "12")

    log.Println("Requesting MasteryGetChampion...")
    return riot.GetData("GET", uri)
}

func MasterySummonerScore() ([]byte, error) {
    // SummonerScore
    // Find the total score for all champions across the board

    uri := mastery + "scores/by-summoner/%s"
    uri = fmt.Sprintf(uri, "28747969")

    log.Println("Requesting MasterySummonerScore...")
    return riot.GetData("GET", uri)
}
