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

	resp, err := c.Get(req)
	if err != nil {
		return matches, err
	}

	err = riot.GetResultFromResp(resp, &matches)
	return matches, err
}

func GetMatch(c riot.ApiClient, match_id string) (match Match, err error) {
	uri := buildUri("matches/" + match_id)
	req := c.NewRequest(uri)

	resp, err := c.Get(req)
	if err != nil {
		return match, err
	}

	err = riot.GetResultFromResp(resp, &match)
	return match, err
}

func GetTimeline(c riot.ApiClient, match_id string) (tl Timeline, err error) {
	uri := buildUri("timelines/by-match/" + match_id)
	req := c.NewRequest(uri)

	resp, err := c.Get(req)
	if err != nil {
		return tl, err
	}

	err = riot.GetResultFromResp(resp, &tl)
	return tl, err
}
