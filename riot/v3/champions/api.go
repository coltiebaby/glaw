package champions

import (
	"github.com/coltiebaby/g-law/riot"
	"github.com/coltiebaby/g-law/riot/v3"
)

var buildUri = BuildUriFunc(`platform`)

func FreeChampions() (ci ChampionInfo, err error) {
	req := riot.Client.NewRequest(buildUri(`champion-rotations`))

	req.Get(&ci)
	return ci, err
}
