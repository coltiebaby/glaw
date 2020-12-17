package league

import (
	"context"
	"fmt"

	"github.com/coltiebaby/glaw"
	"github.com/coltiebaby/glaw/league/core"
)

type MatchRequest struct {
	ID     string
	Region glaw.Region
}

func (mr MatchRequest) String() string {
	return fmt.Sprintf("matches/%s", mr.ID)
}

func (c *Client) Match(ctx context.Context, mr MatchRequest) (matches core.MatchStorage, err error) {
	uri := mr.String()
	req := NewRequest("GET", "match", uri, mr.Region, glaw.V4)

	err = c.Do(ctx, req, &matches)
	return matches, err

}

type MatchesRequest struct {
	AccountID string
	Region    glaw.Region
}

func (mr MatchesRequest) String() string {
	return fmt.Sprintf("matchlists/by-account/%s", mr.AccountID)
}

func (c *Client) Matches(ctx context.Context, mr MatchRequest) (matches core.MatchStorage, err error) {
	uri := mr.String()
	req := NewRequest("GET", "match", uri, mr.Region, glaw.V4)

	err = c.Do(ctx, req, &matches)
	return matches, err
}

type TimelineRequest struct {
	ID     string
	Region glaw.Region
}

func (tr TimelineRequest) String() string {
	return fmt.Sprintf("timelines/by-match/%s", tr.ID)
}

func (c *Client) Timeline(ctx context.Context, mr MatchRequest) (matches core.MatchStorage, err error) {
	uri := mr.String()
	req := NewRequest("GET", "match", uri, mr.Region, glaw.V4)

	err = c.Do(ctx, req, &matches)
	return matches, err
}
