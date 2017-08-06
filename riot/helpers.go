package riot

import (
    "encoding/json"
    "fmt"
    "strconv"

    "github.com/julienschmidt/httprouter"
)

func getSummoner(summoner_type string) Summoner {
    if isInt(summoner_type) {
        summoner_id, _ := strconv.Atoi(summoner_type)
        return Summoner{summoner_id, summoner_id, "USED_ID"}
    }

    return findSummonerByName(summoner_type)
}

func isInt(s string) bool {
    if _, err := strconv.ParseInt(s, 10, 64); err == nil {
        return true
    }
    return false
}

func findSummonerByName(summoner_name string) (Summoner) {
    var (
        body []byte
        err  error
        s    Summoner
    )

    params := httprouter.Params{
        httprouter.Param{"summoner_id", summoner_name},
    }

    body, err = summonerFindByName(&params)

    if err = json.Unmarshal(body, &s); err != nil {
        fmt.Println("error here for unmarshal")
    }

    return s
}
