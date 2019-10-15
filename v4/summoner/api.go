package summoner

import (
	"github.com/coltiebaby/glaw"
	"github.com/coltiebaby/glaw/v4"
)

var buildUri = v4.BuildUriFunc(`summoner`)

func get(c glaw.ApiClient, endpoint string) (summoner Summoner, err error) {
	req := c.NewRequest(buildUri(endpoint))

	resp, err := c.Get(req)
	if err != nil {
		return summoner, err
	}

	err = glaw.GetResultFromResp(resp, &summoner)
	return summoner, err
}

func ByName(c glaw.ApiClient, name string) (summoner Summoner, err error) {
	summoner, err = get(c, "summoners/by-name/"+name)
	return summoner, err
}

func ByAccountID(c glaw.ApiClient, id string) (summoner Summoner, err error) {
	summoner, err = get(c, "summoners/by-account/"+id)
	return summoner, err
}

func ByPUUID(c glaw.ApiClient, puuid string) (summoner Summoner, err error) {
	summoner, err = get(c, "summoners/by-puuid/"+puuid)
	return summoner, err
}

func ByID(c glaw.ApiClient, id string) (summoner Summoner, err error) {
	summoner, err = get(c, "summoners/"+id)
	return summoner, err
}
