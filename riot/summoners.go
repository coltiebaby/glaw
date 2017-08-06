package riot

import (
    "fmt"

    "github.com/julienschmidt/httprouter"
)

type Summoner struct {
    ACCOUNT_ID int `json:"accountId"`
    ID int `json:"id"`
    NAME string `json:"name"`
}

var summoner_uri = fmt.Sprintf("/summoner/%s/summoners/", Version)

func summoner_init(router *httprouter.Router) {
    router.GET("/summoner/name/:summoner_id", hasParams(summonerFindByName))
    router.GET("/summoner/account/:summoner_id", hasParams(summonerFindByAccount))
    router.GET("/summoner/id/:summoner_id", hasParams(summonerFindByID))
}

func summonerFindByName(ps *httprouter.Params) ([]byte, error) {
    uri := summoner_uri + "by-name/%s"
    uri = fmt.Sprintf(uri, ps.ByName("summoner_id"))

    return GetData("GET", uri)
}

func summonerFindByAccount(ps *httprouter.Params) ([]byte, error) {
    uri := summoner_uri + "by-account/%s"
    uri = fmt.Sprintf(uri, ps.ByName("summoner_id"))

    return GetData("GET", uri)
}

func summonerFindByID(ps *httprouter.Params) ([]byte, error) {
    uri := summoner_uri + "%s"
    uri = fmt.Sprintf(uri, ps.ByName("summoner_id"))

    return GetData("GET", uri)
}
