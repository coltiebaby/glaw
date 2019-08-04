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

func GetMatch(match_id string) (match Match, err error) {
	req := riot.RiotRequest{
		Type:    "match",
		Uri:     "matches/" + match_id,
		Version: v4.VERSION,
	}

	err = req.Get(&match)
	return match, err
}

func GetTimeline(match_id string) (tl Timeline, err error) {
	req := riot.RiotRequest{
		Type:    "match",
		Uri:     "timelines/by-match/" + match_id,
		Version: v4.VERSION,
	}

	err = req.Get(&tl)
	return tl, err
}
