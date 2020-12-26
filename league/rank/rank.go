package rank

import (
	"context"
	"fmt"

	"github.com/coltiebaby/glaw"
	"github.com/coltiebaby/glaw/league"
	"github.com/coltiebaby/glaw/league/core"
)

type Client struct {
	client *league.Client
}

func New(c *league.Client) *Client {
	return &Client{
		client: c,
	}
}

func NewRankClient(opts ...glaw.Option) (*Client, error) {
	c, err := league.NewClient(opts...)
	client := &Client{
		client: c,
	}

	return client, err
}

// Queue
//
// View the top tier players and rankings (challenger, master, grandmaster)

type QueueRequest struct {
	Tier   core.Tier
	Queue  core.Queue
	Region glaw.Region
}

func (qr QueueRequest) String() string {
	template := `%sleagues/by-queue/%s`
	return fmt.Sprintf(template, qr.Tier, qr.Queue)
}

func (c *Client) GetQueue(ctx context.Context, qr QueueRequest) (l core.League, err error) {
	switch qr.Tier {
	case core.CHALLENGER, core.MASTER, core.GRANDMASTER:
	default:
		err = fmt.Errorf("tier must be challenger, master, or grandmaster")
		return l, err
	}

	uri := qr.String()
	req := league.NewRequest("GET", "league", uri, qr.Region, glaw.V4)

	err = c.client.Do(ctx, req, &l)
	return l, err
}

// League
//
//

type LeagueRequest struct {
	Id     string
	Region glaw.Region
}

func (lr LeagueRequest) String() string {
	return fmt.Sprintf(`leagues/%s`, lr.Id)
}

func (c *Client) GetLeague(ctx context.Context, lr LeagueRequest) (l core.League, err error) {
	uri := lr.String()
	req := league.NewRequest("GET", "league", uri, lr.Region, glaw.V4)

	err = c.client.Do(ctx, req, &l)
	return l, err
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
func (c *Client) GetEntry(ctx context.Context, er EntryRequest) (entries []core.LeagueEntry, err error) {
	uri := er.String()
	var req league.Request

	switch er.Tier {
	case core.CHALLENGER, core.MASTER, core.GRANDMASTER:
		req = league.NewRequest("GET", "league-exp", uri, er.Region, glaw.V4)
	default:
		req = league.NewRequest("GET", "league", uri, er.Region, glaw.V4)
	}

	err = c.client.Do(ctx, req, &entries)
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
	entries, err = ef.client.GetEntry(ctx, ef.Request)

	if len(entries) == 0 {
		ef.stop = true
		err = NoEntriesErr
	}

	return entries, err
}

var NoEntriesErr error = fmt.Errorf("no more entries")
