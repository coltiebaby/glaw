package riot

import (
    "fmt"

    "github.com/julienschmidt/httprouter"
)

var talent_mastery_uri = fmt.Sprintf("/platform/%s/", Version)

func talent_init(router *httprouter.Router) {
    router.GET("/masteries/:summoner_id", paramsWithSummoner(masteriesGetSummonersMasteries))
}

func masteriesGetSummonersMasteries(ps *httprouter.Params, summoner *Summoner) ([]byte, error) {
    uri := talent_mastery_uri + "masteries/by-summoner/%v"
    uri = fmt.Sprintf(uri, summoner.ID)

    return GetData("GET", uri)
}
