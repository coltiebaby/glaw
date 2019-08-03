package champions

import (
	"github.com/coltiebaby/g-law/riot"
	"github.com/coltiebaby/g-law/riot/v3"
)

func FreeChampions() (ci ChampionInfo, err error) {
	req := riot.RiotRequest{
		Type:    `platform`,
		Uri:     `champion-rotations`,
		Version: v3.VERSION,
	}

	req.Get(&ci)
	return ci, err
}
