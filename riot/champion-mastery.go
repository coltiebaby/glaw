package riot

import (
    "fmt"
    "log"

    "github.com/julienschmidt/httprouter"
)

var mastery = fmt.Sprintf("/champion-mastery/%s/", Version)

func mastery_init(router *httprouter.Router) {
    var cm = "/champ-masteries/"
    router.GET(cm + ":summoner_name/all", paramsWithSummoner(MasteryAllChampions))
    router.GET(cm + ":summoner_name/sum", paramsWithSummoner(MasterySummonerScore))
    router.GET(cm + ":summoner_name/champion/:champion_id", paramsWithSummoner(MasteryGetChampion))
}

func MasteryAllChampions(ps *httprouter.Params, summoner *Summoner) ([]byte, error) {
    // All Champion Mastery
    // Ask Riot for all the summoners champions based on mastery score.
    // Summoner ID = 28747969

    uri := mastery + "champion-masteries/by-summoner/%v"
    uri = fmt.Sprintf(uri, summoner.ID)

    return GetData("GET", uri)
}

func MasteryGetChampion(ps *httprouter.Params, summoner *Summoner) ([]byte, error) {
    // Get Champion Score
    // Find summoner score for a single champion

    uri := mastery + "champion-masteries/by-summoner/%v/by-champion/%s"
    uri = fmt.Sprintf(uri, summoner.ID, ps.ByName("champion_id"))

    return GetData("GET", uri)
}

func MasterySummonerScore(ps *httprouter.Params, summoner *Summoner) ([]byte, error) {
    // SummonerScore
    // Find the total score for all champions across the board

    uri := mastery + "scores/by-summoner/%v"
    uri = fmt.Sprintf(uri, summoner.ID)

    return GetData("GET", uri)
}
