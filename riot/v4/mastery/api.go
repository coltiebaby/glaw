package mastery

import (
	"fmt"

	"github.com/coltiebaby/g-law/riot"
	"github.com/coltiebaby/g-law/riot/v4"
)

var makeUri = v4.BuildUriFunc(`champion-mastery`)

func All(c riot.ApiClient, id string) (cm []ChampionMastery, err error) {
	req := c.NewRequest(makeUri(`champion-masteries/by-summoner/` + id))

	err = req.Get(&cm)
	return cm, err
}

func Score(c riot.ApiClient, id string) (score int, err error) {
	req := c.NewRequest(makeUri(`champion-masteries/by-summoner/` + id))

	err = req.Get(&score)
	return score, err
}

func ByChampionId(c riot.ApiClient, id string, championId int) (cm ChampionMastery, err error) {
	endpoint := fmt.Sprintf(`champion-masteries/by-summoner/%s/by-champion/%d`, id, championId)
	req := c.NewRequest(makeUri(endpoint))

	err = req.Get(&cm)
	return cm, err
}
