package rank

import (
	"context"
	"fmt"
	"strings"

	"github.com/coltiebaby/glaw"
	"github.com/coltiebaby/glaw/tft"
	"github.com/coltiebaby/glaw/tft/core"
)

type Client struct {
	client *tft.Client
}

func New(c *tft.Client) *Client {
	return &Client{
		client: c,
	}
}

func NewRankClient(opts ...glaw.Option) (*Client, error) {
	c, err := tft.NewClient(opts...)
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
	Region glaw.Region
}

func (qr QueueRequest) String() string {
	return strings.ToLower(string(qr.Tier))
}

func (c *Client) GetQueue(ctx context.Context, qr QueueRequest) (l core.Rank, err error) {
	switch qr.Tier {
	case core.CHALLENGER, core.MASTER, core.GRANDMASTER:
	default:
		err = fmt.Errorf("tier must be challenger, master, or grandmaster")
		return l, err
	}

	uri := qr.String()
	req := tft.NewRequest("GET", "league", uri, qr.Region, glaw.V1)

	err = c.client.Do(ctx, req, &l)
	return l, err
}

// League
//
//

type RankRequest struct {
	Id     string
	Region glaw.Region
}

func (lr RankRequest) String() string {
	return fmt.Sprintf(`league/%s`, lr.Id)
}

func (c *Client) GetRank(ctx context.Context, lr RankRequest) (l core.Rank, err error) {
	uri := lr.String()
	req := tft.NewRequest("GET", "league", uri, lr.Region, glaw.V1)

	err = c.client.Do(ctx, req, &l)
	return l, err
}

// Entry

// Entry Request will default to the by summoner search if you include it instead of doing the
// specific search.
type EntryRequest struct {
	SummonerId string
	Tier       core.Tier
	Division   core.Division
	Region     glaw.Region
	Page       int
}

func (er EntryRequest) String() string {
	if er.SummonerId != "" {
		return fmt.Sprintf(`entries/by-summoner/%s`, er.SummonerId)
	}

	template := `entries/%s/%s`
	if er.Page > 1 {
		template = fmt.Sprintf("%s?page=%d", template, er.Page)
	}

	return fmt.Sprintf(template, er.Tier, er.Division)
}

// Entry will grab a slice of entries. Uses the experimental api if you use challenger, master, or
// grandmaster.
func (c *Client) GetEntry(ctx context.Context, er EntryRequest) (entries []core.Entry, err error) {
	uri := er.String()
	req := tft.NewRequest("GET", "league", uri, er.Region, glaw.V1)

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

func (ef *EntryFetcher) Next(ctx context.Context) (entries []core.Entry, err error) {
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
