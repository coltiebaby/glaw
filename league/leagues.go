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
	Division   core.Division
	Region     glaw.Region
	Page       int
}

func (er EntryRequest) String() string {
	if er.SummonerId != "" {
		return fmt.Sprintf(`entries/by-summoner/%s`, er.SummonerId)
	}

	template := `entries/%s/%s/%s`
	if er.Page > 1 {
		template = fmt.Sprintf("%s?page=%d", template, er.Page)
	}

	return fmt.Sprintf(template, er.Queue, er.Tier, er.Division)
}

// Entry will grab a slice of entries. Uses the experimental api if you use challenger, master, or
// grandmaster.
func (c *Client) Entry(ctx context.Context, er EntryRequest) (entries []core.LeagueEntry, err error) {
	uri := er.String()
	var req Request

	switch er.Tier {
	case core.CHALLENGER, core.MASTER, core.GRANDMASTER:
		req = NewRequest("GET", "league-exp", uri, er.Region, glaw.V4)
	default:
		req = NewRequest("GET", "league", uri, er.Region, glaw.V4)
	}

	err = c.Do(ctx, req, &entries)
	return entries, err
}

type EntryFetcher struct {
	client  *Client
	Request EntryRequest
	stop    bool
}

func NewEntryFetcher(c *Client, req EntryRequest) *EntryFetcher {
	req.Page = 1
	return &EntryFetcher{
		client:  c,
		Request: req,
	}
}

func (ef *EntryFetcher) Next(ctx context.Context) (entries []core.LeagueEntry, err error) {
	if ef.stop {
		return entries, NoEntriesErr
	}

	ef.Request.Page++
	entries, err = ef.client.Entry(ctx, ef.Request)

	if len(entries) == 0 {
		ef.stop = true
		err = NoEntriesErr
	}

	return entries, err
}

var NoEntriesErr error = fmt.Errorf("no more entries")
