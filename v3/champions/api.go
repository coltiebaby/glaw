package champions

import (
	"github.com/coltiebaby/glaw"
	"github.com/coltiebaby/glaw/v3"
)

var buildUri = v3.BuildUriFunc(`platform`)

func FreeChampions(c glaw.ApiClient) (ci ChampionInfo, err error) {
	req := c.NewRequest(buildUri(`champion-rotations`))

	resp, err := c.Get(req)
	if err != nil {
		return ci, err
	}

	err = glaw.GetResultFromResp(resp, &ci)
	return ci, err
}
