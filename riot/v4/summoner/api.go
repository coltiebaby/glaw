package summoner

import (
	"github.com/coltiebaby/g-law/riot"
	"github.com/coltiebaby/g-law/riot/v4"
)

func get(uri string) (summoner Summoner, err error) {
	req := riot.RiotRequest{
		Type:    `summoner`,
		Uri:     uri,
		Version: v4.VERSION,
	}

	req.Get(&summoner)
	return summoner, err
}

func ByName(name string) (summoner Summoner, err error) {
	summoner, err = get("summoners/by-name/" + name)
	return summoner, err
}
