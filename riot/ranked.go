package riot

import (
    "fmt"

    "github.com/julienschmidt/httprouter"
)

const (
    flex =             "RANKED_FLEX_SR"
    solo =             "RANKED_SOLO_5x5"
    twisted_treeline = "RANKED_FLEX_TT"
)

var ranked_uri = fmt.Sprintf("/league/%s/", Version)

func ranked_init(router *httprouter.Router) {
    router.GET("/ranked/leaderboard/:queue_id", hasParams(rankedLeaderBoard))
    router.GET("/ranked/summoner/:summoner_name", hasParams(rankedSummoner))
}

func findQueueID(queue_type string) (string, error) {
    var queue_id string
    var err error

	switch queue_type {
        case "flex", "flex-queue":
            queue_id = flex
        case "solo", "solo-queue":
            queue_id = solo
        case "tt", "twisted", "twisted-treeline":
            queue_id = twisted_treeline
        default:
            fmt.Println("Raise an error for rankedLEADERBOARD")
    }

    return queue_id, err
}

func rankedLeaderBoard(ps *httprouter.Params) ([]byte, error) {
    queue_id, _ := findQueueID(ps.ByName("queue_id"))

    uri := ranked_uri + fmt.Sprintf("challengerleagues/by-queue/%s", queue_id)
    return GetData("GET", uri)
}

func rankedSummoner(ps *httprouter.Params) ([]byte, error) {
    var summoner Summoner
    summoner = findSummonerByName(ps.ByName("summoner_name"))

    uri := ranked_uri + fmt.Sprintf("leagues/by-summoner/%v", summoner.ID)

    return GetData("GET", uri)
}
