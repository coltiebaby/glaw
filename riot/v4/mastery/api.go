package mastery

import (
	"fmt"

	"github.com/coltiebaby/g-law/riot"
	"github.com/coltiebaby/g-law/riot/v4"
)

var makeUri = v4.BuildUriFunc(`champion-mastery`)

func All(id string) (cm []ChampionMastery, err error) {
	req := riot.Client.NewRequest(makeUri(`champion-masteries/by-summoner/` + id))

	err = req.Get(&cm)
	return cm, err
}

func Score(id string) (score int, err error) {
	req := riot.Client.NewRequest(makeUri(`champion-masteries/by-summoner/` + id))
	req := riot.RiotRequest{
		Type:    `champion-mastery`,
		Uri:     `scores/by-summoner/` + id,
		Version: v4.VERSION,
	}

	err = req.Get(&score)
	return score, err
}

func ByChampionId(id string, championId int) (cm ChampionMastery, err error) {
	endpoint := fmt.Sprintf(`champion-masteries/by-summoner/%s/by-champion/%d`, id, championId)
	req := riot.Client.NewRequest(makeUri(endpoint))

	err = req.Get(&cm)
	return cm, err
}
