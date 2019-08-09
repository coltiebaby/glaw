package leagues

import (
	"fmt"
	"strconv"

	"github.com/coltiebaby/g-law/riot"
	"github.com/coltiebaby/g-law/riot/v4"
)

var buildUri = v4.BuildUriFunc(`league`)

func getLeague(c riot.ApiClient, endpoint string) (league League, err error) {
	req := c.NewRequest(buildUri(endpoint))

	resp, err := c.Get(req)
	if err != nil {
		return league, err
	}

	err = riot.GetResultFromResp(resp, &league)
	return league, err
}

func getEntries(c riot.ApiClient, endpoint string, page int) (entries []LeagueEntry, err error) {
	req := c.NewRequest(buildUri(endpoint))
	if page > 0 {
		req.AddParameter(`page`, strconv.Itoa(page))
	}

	resp, err := c.Get(req)
	if err != nil {
		return entries, err
	}

	err = riot.GetResultFromResp(resp, entries)
	return entries, err
}

func Challengers(c riot.ApiClient, queue Queue) (league League, err error) {
	league, err = getLeague(c, fmt.Sprintf(`challengerleagues/by-queue/%s`, queue))
	return league, err
}

func Masters(c riot.ApiClient, queue Queue) (league League, err error) {
	league, err = getLeague(c, fmt.Sprintf(`masterleagues/by-queue/%s`, queue))
	return league, err
}

func GrandMasters(c riot.ApiClient, queue Queue) (league League, err error) {
	league, err = getLeague(c, fmt.Sprintf(`grandmasterleagues/by-queue/%s`, queue))
	return league, err
}

func GetLeagueByID(c riot.ApiClient, leagueID string) (league League, err error) {
	league, err = getLeague(c, fmt.Sprintf(`leagues/%s`, leagueID))
	return league, err
}

// EntriesFactory builds out a closer to make it easier to get multiple pages
// You'll have to manually check if there's anymore pages
//
//  f := EntriesFactory(leagues.SOLO, leagues.GOLD, leagues.ONE)
//  for {
//       page := 1
//       if results, err := f(page); err != nil {
//            if len(results) == 0 {
//                 break
//            }
//            // ...do something with results
//       }
//       page = page + 1
//  }
func EntriesFactory(c riot.ApiClient, queue Queue, tier Tier, division Division) func(int) ([]LeagueEntry, error) {
	uri := fmt.Sprintf(`entries/%s/%s/%s`, queue, tier, division)
	return func(page int) ([]LeagueEntry, error) {
		return Entries(c, uri, page)
	}
}

func Entries(c riot.ApiClient, uri string, page int) (entries []LeagueEntry, err error) {
	entries, err = getEntries(c, uri, page)
	return entries, err
}

func EntriesBySummonerID(c riot.ApiClient, id string) (entries []LeagueEntry, err error) {
	entries, err = getEntries(c, fmt.Sprintf(`entries/by-summoner/%s`, id), -1)
	return entries, err
}
