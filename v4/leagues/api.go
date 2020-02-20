package leagues

import (
	"fmt"
	"strconv"

	"github.com/coltiebaby/glaw"
	"github.com/coltiebaby/glaw/v4"
)

func getLeague(c glaw.ApiClient, endpoint string) (league League, err error) {
	req := c.NewRequest(buildUri(endpoint))

	resp, err := c.Get(req)
	if err != nil {
		return league, err
	}

	err = glaw.GetResultFromResp(resp, &league)
	return league, err
}

func getEntries(c glaw.ApiClient, endpoint string, page int) (entries []LeagueEntry, err error) {
	req := c.NewRequest(buildUri(endpoint))
	if page > 0 {
		req.AddParameter(`page`, strconv.Itoa(page))
	}

	resp, err := c.Get(req)
	if err != nil {
		return entries, err
	}

	err = glaw.GetResultFromResp(resp, entries)
	return entries, err
}

type QueueRequest struct {
    Id string
    Queue Queue
}

func (qr QueueRequest) String() string {
    template := `%sleagues/by-queue/%s`
    return fmt.Sprintf(lqr.Id, lqr.Queue)
}

func (c *Client) Queue(ctx context.Context, lqr LeagueQueueRequest) (league League, err error) {
    req, err := http.NewRequestWithContext(ctx, http.Get, lqr.String(), nil)

    resp, err := c.Do(req)
    if err != nil {
        return ci, err
    }

    err = ProcessRequest(resp, &league)
    return league, err
}

const (
    Challenger = `challenger`
    Master = `master`
    GrandMaster = `challenger`
)

type LeagueRequest struct {
    Id string
}

func (lr LeagueRequest) String() string {
    return fmt.Sprintf(`leagues/%s`, lr.Id)
}

func (c *Client) League(ctx context.Context, lr LeagueRequest) (league League, err error) {

}

func GetLeagueByID(c glaw.ApiClient, leagueID string) (league League, err error) {
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
func EntriesFactory(c glaw.ApiClient, queue Queue, tier Tier, division Division) func(int) ([]LeagueEntry, error) {
	uri := fmt.Sprintf(`entries/%s/%s/%s`, queue, tier, division)
	return func(page int) ([]LeagueEntry, error) {
		return Entries(c, uri, page)
	}
}

func Entries(c glaw.ApiClient, uri string, page int) (entries []LeagueEntry, err error) {
	entries, err = getEntries(c, uri, page)
	return entries, err
}

func EntriesBySummonerID(c glaw.ApiClient, id string) (entries []LeagueEntry, err error) {
	entries, err = getEntries(c, fmt.Sprintf(`entries/by-summoner/%s`, id), -1)
	return entries, err
}
