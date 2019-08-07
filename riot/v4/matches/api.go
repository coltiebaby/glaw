package matches

import (
	"fmt"
	"net/url"

	"github.com/coltiebaby/g-law/riot"
	"github.com/coltiebaby/g-law/riot/v4"
)

var buildUri = v4.BuildUriFunc(`match`)

func GetMatchlists(c riot.ApiClient, id string, values url.Values) (matches MatchStorage, err error) {
	uri := buildUri(fmt.Sprintf(`matchlists/by-account/%s`, id))

	req := c.NewRequest(uri)
	req.SetParameters(values)

	req.Get(&matches)
	return matches, err
}

func GetMatch(c riot.ApiClient, match_id string) (match Match, err error) {
	uri := buildUri("matches/" + match_id)
	req := c.NewRequest(uri)

	err = req.Get(&match)
	return match, err
}

func GetTimeline(c riot.ApiClient, match_id string) (tl Timeline, err error) {
	uri := buildUri("timelines/by-match/" + match_id)
	req := c.NewRequest(uri)

	err = req.Get(&tl)
	return tl, err
}
