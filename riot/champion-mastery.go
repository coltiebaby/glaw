package riot

import (
    "fmt"
    "log"

    "github.com/julienschmidt/httprouter"
)

func mastery_init(router *httprouter.Router) {
    router.GET("/summoner/:summoner_id/champion/masteries", hasParams(MasteryAllChampions))
    //router.GET("/summoner/:summoner_id/champion/:champion_id/mastery", hasParams(m.MasteryGetChampion))
    //router.GET("/summoner/:summoner_id/champion/mastery/sum", hasParams(m.MasterySummonerScore))
}

var mastery = fmt.Sprintf("/champion-mastery/%s/", Version)

func MasteryAllChampions(ps *httprouter.Params) ([]byte, error) {
    // All Champion Mastery
    // Ask Riot for all the summoners champions based on mastery score.
    // Summoner ID = 28747969

    uri := mastery + "champion-masteries/by-summoner/%s"
    uri = fmt.Sprintf(uri, ps.ByName("summoner_id"))

    log.Println("[GET] %v", uri)
    return GetData("GET", uri)
}

func MasteryGetChampion(ps *httprouter.Params) ([]byte, error) {
    // Get Champion Score
    // Find summoner score for a single champion

    uri := mastery + "champion-masteries/by-summoner/%s/by-champion/%s"
    uri = fmt.Sprintf(uri, ps.ByName("summoner_name"), ps.ByName("champion_id"))

    log.Println("Requesting MasteryGetChampion...")
    return GetData("GET", uri)
}

func MasterySummonerScore(ps *httprouter.Params) ([]byte, error) {
    // SummonerScore
    // Find the total score for all champions across the board

    uri := mastery + "scores/by-summoner/%s"
    uri = fmt.Sprintf(uri, ps.ByName("summoner_id"))

    log.Println("Requesting MasterySummonerScore...")
    return GetData("GET", uri)
}
