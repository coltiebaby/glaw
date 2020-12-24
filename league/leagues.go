package league

import (
	"context"
	"fmt"

	"github.com/coltiebaby/glaw"
	"github.com/coltiebaby/glaw/league/core"
)

// Queue

type QueueRequest struct {
	Tier   core.Tier
	Queue  core.Queue
	Region glaw.Region
}

func (qr QueueRequest) String() string {
	template := `%sleagues/by-queue/%s`
	return fmt.Sprintf(template, qr.Tier, qr.Queue)
}

func (c *Client) Queue(ctx context.Context, qr QueueRequest) (league core.League, err error) {
	switch qr.Tier {
	case core.CHALLENGER, core.MASTER, core.GRANDMASTER:
	default:
		err = fmt.Errorf("tier must be challenger, master, or grandmaster")
		return league, err
	}

	uri := qr.String()
	req := NewRequest("GET", "league", uri, qr.Region, glaw.V4)

	err = c.Do(ctx, req, &league)
	return league, err
}

// League

type LeagueRequest struct {
	ID     string
	Region glaw.Region
}

func (lr LeagueRequest) String() string {
	return fmt.Sprintf(`leagues/%s`, lr.ID)
}

func (c *Client) League(ctx context.Context, lr LeagueRequest) (league core.League, err error) {
	uri := lr.String()
	req := NewRequest("GET", "league", uri, lr.Region, glaw.V4)

	err = c.Do(ctx, req, &league)
	return league, err
}

// Entry

// Entry Request will default to the by summoner search if you include it instead of doing the
// specific search.
type EntryRequest struct {
	SummonerId string
	Queue      core.Queue
	Tier       core.Tier
	Division   core.Divsion
	Region     glaw.Region
}

func (er EntryRequest) String() string {
	if er.SummonerId != "" {
		return fmt.Sprintf(`entries/by-summoner/%s`, er.SummonerId)
	}

	return fmt.Sprintf(`entries/%s/%s/%s`, er.Queue, er.Tier, er.Divison)
}

// Entry will grab a slice of entries. Uses the experimental api if you use challenger, master, or
// grandmaster.
func (c *Client) Entry(ctx context.Context, er EntryRequest) (entries []core.LeagueEntry, err error) {
	uri := er.String()
	var req Request

	switch qr.Tier {
	case core.CHALLENGER, core.MASTER, core.GRANDMASTER:
		req = NewRequest("GET", "league-exp", uri, er.Region, glaw.V4)
	default:
		req = NewRequest("GET", "league", uri, er.Region, glaw.V4)
	}

	err = c.Do(ctx, req, &entries)
	return entries, err
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
// func EntriesFactory(c glaw.ApiClient, queue Queue, tier Tier, division Division) func(int) ([]LeagueEntry, error) {
// 	uri := fmt.Sprintf(`entries/%s/%s/%s`, queue, tier, division)
// 	return func(page int) ([]LeagueEntry, error) {
// 		return Entries(c, uri, page)
// 	}
// }
//
// func Entries(c glaw.ApiClient, uri string, page int) (entries []LeagueEntry, err error) {
// 	entries, err = getEntries(c, uri, page)
// 	return entries, err
// }
//
// func EntriesBySummonerID(c glaw.ApiClient, id string) (entries []LeagueEntry, err error) {
// 	entries, err = getEntries(c, fmt.Sprintf(`entries/by-summoner/%s`, id), -1)
// 	return entries, err
// }
