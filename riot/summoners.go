package riot

import (
    "fmt"
    "encoding/json"
)

type Summoner struct {
	ID            int    `json:"id"`
	AccountID     int    `json:"accountId"`
	Name          string `json:"name"`
	ProfileIconID int    `json:"profileIconId"`
	RevisionDate  int64  `json:"revisionDate"`
	SummonerLevel int    `json:"summonerLevel"`
}


func GetSummonerByAccount(account_id int) (Summoner) {
    // /lol/summoner/v3/summoners/by-account/{accountId}
    var summoner Summoner

    rr := &RiotRequest {
        Type: "summoner",
        Uri:  fmt.Sprintf("summoners/by-account/%s", account_id),
    }

    resp := rr.GetData()
    json.Unmarshal(resp, &summoner)

    return summoner
}


func GetSummonerById(summoner_id int) (Summoner) {
    // /lol/summoner/v3/summoners/{summonerId}
    var summoner Summoner

    rr := &RiotRequest {
        Type: "summoner",
        Uri:  fmt.Sprintf("summoners/%s", summoner_id),
    }

    resp := rr.GetData()
    json.Unmarshal(resp, &summoner)

    return summoner
}


func GetSummonerByName(name string) (Summoner) {
    // /lol/summoner/v3/summoners/by-name/{summonerName}
    var summoner Summoner

    rr := &RiotRequest {
        Type: "summoner",
        Uri:  fmt.Sprintf("summoners/by-name/%s", name),
    }

    resp := rr.GetData()
    json.Unmarshal(resp, &summoner)

    return summoner
}
