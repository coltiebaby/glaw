package champions

import (
	"github.com/coltiebaby/g-law/riot"
	"github.com/coltiebaby/g-law/riot/v3"
)

var buildUri = v3.BuildUriFunc(`platform`)

func FreeChampions(c riot.ApiClient) (ci ChampionInfo, err error) {
	req := c.NewRequest(buildUri(`champion-rotations`))

	req.Get(&ci)
	return ci, err
}
