package summoner

import (
	"github.com/coltiebaby/g-law/riot"
	"github.com/coltiebaby/g-law/riot/v4"
)

var buildUri = v4.BuildUriFunc(`summoner`)

func get(endpoint string) (summoner Summoner, err error) {
	req := riot.Client.NewRequest(buildUri(endpoint))
	req.Get(&summoner)
	return summoner, err
}

func ByName(name string) (summoner Summoner, err error) {
	summoner, err = get("summoners/by-name/" + name)
	return summoner, err
}

func ByAccountID(id string) (summoner Summoner, err error) {
	summoner, err = get("summoners/by-account/" + id)
	return summoner, err
}

func ByPUUID(puuid string) (summoner Summoner, err error) {
	summoner, err = get("summoners/by-puuid/" + puuid)
	return summoner, err
}

func ByID(id string) (summoner Summoner, err error) {
	summoner, err = get("summoners/" + id)
	return summoner, err
}
