package matches

import (
	"fmt"
	"net/url"

	"github.com/coltiebaby/g-law/riot"
	"github.com/coltiebaby/g-law/riot/v4"
)

func GetMatchlists(id string, values url.Values) (matches MatchStorage, err error) {
	uri := fmt.Sprintf(`matchlists/by-account/%s`, id)

	req := riot.RiotRequest{
		Type:    `match`,
		Uri:     uri,
		Version: v4.VERSION,
		Params:  values,
	}

	req.Get(&matches)
	return matches, err
}
