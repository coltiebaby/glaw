package mastery

import (
	"fmt"

	"github.com/coltiebaby/glaw"
	"github.com/coltiebaby/glaw/v4"
)

var makeUri = v4.BuildUriFunc(`champion-mastery`)

func Score(c glaw.ApiClient, id string) (score int, err error) {
	req := c.NewRequest(makeUri(`champion-masteries/by-summoner/` + id))

	resp, err := c.Get(req)
	if err != nil {
		return score, err
	}

	err = glaw.GetResultFromResp(resp, &score)
	return score, err
}

func All(c glaw.ApiClient, id string) (cm []ChampionMastery, err error) {
	req := c.NewRequest(makeUri(`champion-masteries/by-summoner/` + id))

	resp, err := c.Get(req)
	if err != nil {
		return cm, err
	}

	err = glaw.GetResultFromResp(resp, &cm)
	return cm, err
}

func ByChampionId(c glaw.ApiClient, id string, championId int) (cm ChampionMastery, err error) {
	endpoint := fmt.Sprintf(`champion-masteries/by-summoner/%s/by-champion/%d`, id, championId)
	req := c.NewRequest(makeUri(endpoint))

	resp, err := c.Get(req)
	if err != nil {
		return cm, err
	}

	err = glaw.GetResultFromResp(resp, &cm)
	return cm, err
}
