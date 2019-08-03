package leagues

import (
	"fmt"
	"strconv"

	"github.com/coltiebaby/g-law/riot"
	"github.com/coltiebaby/g-law/riot/v4"
)

func getLeague(uri string) (league League, err error) {
	req := riot.RiotRequest{
		Type:    `league`,
		Uri:     uri,
		Version: v4.VERSION,
	}

	req.Get(&league)
	return league, err
}

func getEntries(uri string, page int) (entries []LeagueEntry, err error) {
	var params map[string]string = make(map[string]string)
	if page > 0 {
		params[`page`] = strconv.Itoa(page)
	}

	req := riot.RiotRequest{
		Type:    `league`,
		Uri:     uri,
		Version: v4.VERSION,
		Params:  params,
	}

	err = req.Get(&entries)
	return entries, err
}

func Challengers(queue Queue) (league League, err error) {
	league, err = getLeague(fmt.Sprintf(`challengerleagues/by-queue/%s`, queue))
	return league, err
}

func Masters(queue Queue) (league League, err error) {
	league, err = getLeague(fmt.Sprintf(`masterleagues/by-queue/%s`, queue))
	return league, err
}

func GrandMasters(queue Queue) (league League, err error) {
	league, err = getLeague(fmt.Sprintf(`grandmasterleagues/by-queue/%s`, queue))
	return league, err
}

func GetLeagueByID(leagueID string) (league League, err error) {
	league, err = getLeague(fmt.Sprintf(`leagues/%s`, queue))
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
func EntriesFactory(queue Queue, tier Tier, division Division) func(int) ([]LeagueEntry, error) {
	uri := fmt.Sprintf(`entries/%s/%s/%s`, queue, tier, division)
	return func(page int) ([]LeagueEntry, error) {
		return Entries(uri, page)
	}
}

func Entries(uri string, page int) (entries []LeagueEntry, err error) {
	entries, err = getEntries(uri, page)
	return entries, err
}

func EntriesBySummonerID(id string) (entries []LeagueEntry, err error) {
	entries, err = getEntries(fmt.Sprintf(`entries/by-summoner/%s`, id), -1)
	return entries, err
}
