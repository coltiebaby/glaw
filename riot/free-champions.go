package riot

import (
    "fmt"
    "encoding/json"
)

type FreeChampion struct {
	RankedPlayEnabled bool `json:"rankedPlayEnabled"`
	BotEnabled        bool `json:"botEnabled"`
	BotMmEnabled      bool `json:"botMmEnabled"`
	Active            bool `json:"active"`
	FreeToPlay        bool `json:"freeToPlay"`
	ID                int  `json:"id"`
}

type FreeChampions struct {
    Champions []FreeChampion `json:"champions"`
}

func GetFreeChampions(show_free_champions bool) (FreeChampions) {
    free_champs := FreeChampions{}
    params := map[string]string{"freeToPlay": fmt.Sprint(show_free_champions)}

    rr := &RiotRequest {
        Type: "platform",
        Uri:  "champions",
        Params: params,
    }

    resp := rr.GetData()
    err := json.Unmarshal(resp, &free_champs)
    if err != nil {
        panic(err)
    }

    return free_champs
}

func GetFreeChampion(champion_id int) (FreeChampion) {
    free_champ := FreeChampion{}

    rr := &RiotRequest {
        Type: "platform",
        Uri:  fmt.Sprintf("champions/%d", champion_id),
    }

    resp := rr.GetData()
    err := json.Unmarshal(resp, &free_champ)
    if err != nil {
        panic(err)
    }

    return free_champ

}
