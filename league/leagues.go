package league

import (
	"context"
	"fmt"

	"github.com/coltiebaby/glaw"
	"github.com/coltiebaby/glaw/league/core"
)

type QueueRequest struct {
	Type   string
	Queue  core.Queue
	Region glaw.Region
}

func (qr QueueRequest) String() string {
	template := `%sleagues/by-queue/%s`
	return fmt.Sprintf(template, qr.Type, qr.Queue)
}

func (c *Client) Queue(ctx context.Context, qr QueueRequest) (league core.League, err error) {
	req := glaw.Request{
		Method:  `GET`,
		Domain:  `league`,
		Version: glaw.V4,
		Region:  qr.Region,
		Uri:     qr.String(),
	}

	resp, err := c.Do(ctx, req)
	if err != nil {
		return league, err
	}

	err = glaw.ProcessResponse(resp, &league)
	return league, err
}

type LeagueRequest struct {
	ID     string
	Region glaw.Region
}

func (lr LeagueRequest) String() string {
	return fmt.Sprintf(`leagues/%s`, lr.ID)
}

func (c *Client) League(ctx context.Context, lr LeagueRequest) (league core.League, err error) {
	req := glaw.Request{
		Method:  `GET`,
		Domain:  `league`,
		Version: glaw.V4,
		Region:  lr.Region,
		Uri:     lr.String(),
	}

	resp, err := c.Do(ctx, req)
	if err != nil {
		return league, err
	}

	err = glaw.ProcessResponse(resp, &league)
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
