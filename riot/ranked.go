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

const (
    challengers = "challenger"
    masters     = "master"
)

var ranked_uri = fmt.Sprintf("/league/%s/", Version)

func ranked_init(router *httprouter.Router) {
    router.GET("/ranked/leaderboard/:tier_id/:queue_id", hasParams(rankedLeaderBoard))
    router.GET("/ranked/summoner/:summoner_name", hasParams(rankedSummoner))
    router.GET("/ranked/stats/:summoner_name", hasParams(rankedPositions))
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
            fmt.Println("Raise an error for find queue id.")
    }

    return queue_id, err
}

func findTierByID(tier_type string) (string, error) {
    var tier_id string
    var err error

    switch tier_type {
        case challengers, challengers + "s":
            tier_id = challengers
        case masters, masters + "s":
            tier_id = masters
        default:
            fmt.Println("Raise an error for find tier by id.")
    }

    return tier_id, err
}

func rankedLeaderBoard(ps *httprouter.Params) ([]byte, error) {
    tier_id, _ := findTierByID(ps.ByName("tier_id"))
    queue_id, _ := findQueueID(ps.ByName("queue_id"))

    uri := ranked_uri + fmt.Sprintf("%sleagues/by-queue/%s", tier_id, queue_id)
    return GetData("GET", uri)
}

func rankedSummoner(ps *httprouter.Params) ([]byte, error) {
    var summoner Summoner
    summoner = findSummonerByName(ps.ByName("summoner_name"))

    uri := ranked_uri + fmt.Sprintf("leagues/by-summoner/%v", summoner.ID)

    return GetData("GET", uri)
}

func rankedPositions(ps *httprouter.Params) ([]byte, error) {
    var summoner Summoner
    summoner = findSummonerByName(ps.ByName("summoner_name"))

    uri := ranked_uri + fmt.Sprintf("positions/by-summoner/%v", summoner.ID)

    return GetData("GET", uri)
}
