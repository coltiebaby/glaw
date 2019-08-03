package mastery

import (
	"fmt"

	"github.com/coltiebaby/g-law/riot"
	"github.com/coltiebaby/g-law/riot/v4"
)

func All(id string) (cm []ChampionMastery, err error) {
	req := riot.RiotRequest{
		Type:    `champion-mastery`,
		Uri:     `champion-masteries/by-summoner/` + id,
		Version: v4.VERSION,
	}

	err = req.Get(&cm)
	return cm, err
}

func Score(id string) (score int, err error) {
	req := riot.RiotRequest{
		Type:    `champion-mastery`,
		Uri:     `scores/by-summoner/` + id,
		Version: v4.VERSION,
	}

	err = req.Get(&score)
	return score, err
}

func ByChampionId(id string, championId int) (cm ChampionMastery, err error) {
	req := riot.RiotRequest{
		Type:    `champion-mastery`,
		Uri:     fmt.Sprintf(`champion-masteries/by-summoner/%s/by-champion/%d`, id, championId),
		Version: v4.VERSION,
	}

	err = req.Get(&cm)
	return cm, err
}
