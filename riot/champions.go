package riot

import (
    "fmt"

    "github.com/julienschmidt/httprouter"
)

var champion_uri = fmt.Sprintf("/platform/%s/champions/", Version)

func champion_init(router *httprouter.Router) {
    router.GET("/champions/info", hasParams(summonerFindByName))
    router.GET("/champions/id/:champion_id", hasParams(summonerFindByAccount))
}

func championFindAll(_ *httprouter.Params) ([]byte, error) {
    return GetData("GET", champion_uri)
}

func championFindByID(ps *httprouter.Params) ([]byte, error) {
    uri := champion_uri + "%s"
    uri = fmt.Sprintf(uri, ps.ByName("champion_id"))

    return GetData("GET", uri)
}
