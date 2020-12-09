package glaw

import (
	"context"
	"fmt"
)

type QueueRequest struct {
	Type   string
	Queue  Queue
	Region Region
}

func (qr QueueRequest) String() string {
	template := `%sleagues/by-queue/%s`
	return fmt.Sprintf(template, qr.Type, qr.Queue)
}

func (c *Client) Queue(ctx context.Context, qr QueueRequest) (league League, err error) {
	req := Request{
		Method:  `GET`,
		Domain:  `league`,
		Version: V4,
		Region:  qr.Region,
		Uri:     qr.String(),
	}

	r, err := req.NewHttpRequest(ctx)
	if err != nil {
		return league, err
	}

	resp, err := c.Do(r)
	if err != nil {
		return league, err
	}

	err = ProcessResponse(resp, &league)
	return league, err
}

type LeagueRequest struct {
	ID     string
	Region Region
}

func (lr LeagueRequest) String() string {
	return fmt.Sprintf(`leagues/%s`, lr.ID)
}

func (c *Client) League(ctx context.Context, lr LeagueRequest) (league League, err error) {
	req := Request{
		Method:  `GET`,
		Domain:  `league`,
		Version: V4,
		Region:  lr.Region,
		Uri:     lr.String(),
	}

	r, err := req.NewHttpRequest(ctx)
	if err != nil {
		return league, err
	}

	resp, err := c.Do(r)
	if err != nil {
		return league, err
	}

	err = ProcessResponse(resp, &league)
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

type Queue string

const (
	Challenger  = `challenger`
	Master      = `master`
	GrandMaster = `grandmaster`
)

const (
	FLEX               Queue = `RANKED_FLEX`
	SOLO               Queue = `RANKED_SOLO_5x5`
	TEAM_FIGHT_TACTICS Queue = `RANKED_FLEX_TFT`
	TWISTED_TREELINE   Queue = `RANKED_FLEX_TT`
)

type Tier string

const (
	// CHALLENGER Tier = "CHALLENGER"
	DIAMOND  Tier = `DIAMOND`
	PLATINUM Tier = `PLATINUM`
	GOLD     Tier = `GOLD`
	SILVER   Tier = `SILVER`
	BRONZE   Tier = `BRONZE`
	IRON     Tier = `IRON`
)

type Division string

const (
	ONE   Division = `I`
	TWO   Division = `II`
	THREE Division = `III`
	FOUR  Division = `IV`
)

type Entry struct {
	MiniSeries   MiniSeries
	SummonerName string `json:"summonerName"`
	HotStreak    bool   `json:"hotStreak"`
	Wins         int    `json:"wins"`
	Veteran      bool   `json:"veteran"`
	Losses       int    `json:"losses"`
	Rank         string `json:"rank"`
	Inactive     bool   `json:"inactive"`
	FreshBlood   bool   `json:"freshBlood"`
	SummonerID   string `json:"summonerId"`
	LeaguePoints int    `json:"leaguePoints"`
}

type MiniSeries struct {
	Progress string `json:"progress"`
	Losses   int    `json:"losses"`
	Target   int    `json:"target"`
	Wins     int    `json:"wins"`
}

type League struct {
	Tier     string  `json:"tier"`
	LeagueID string  `json:"leagueId"`
	Entries  []Entry `json:"entries"`
	Queue    string  `json:"queue"`
	Name     string  `json:"name"`
}

// Has slightly more info than the Entry struct
type LeagueEntry struct {
	QueueType string `json:"queueType"`
	Tier      string `json:"tier"`
	Entry
}
