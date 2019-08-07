package matches

import (
	"fmt"
	"net/url"

	"github.com/coltiebaby/g-law/riot"
	"github.com/coltiebaby/g-law/riot/v4"
)

var buildUri = v4.BuildUriFunc(`match`)

func GetMatchlists(id string, values url.Values) (matches MatchStorage, err error) {
	uri := buildUri(fmt.Sprintf(`matchlists/by-account/%s`, id))

	req := riot.Client.NewRequest(uri)
	req.SetParameters(values)

	req.Get(&matches)
	return matches, err
}

func GetMatch(match_id string) (match Match, err error) {
	uri := buildUri("matches/" + match_id)
	req := riot.Client.NewRequest(uri)

	err = req.Get(&match)
	return match, err
}

func GetTimeline(match_id string) (tl Timeline, err error) {
	uri := buildUri("timelines/by-match/" + match_id)
	req := riot.Client.NewRequest(uri)

	err = req.Get(&tl)
	return tl, err
}
