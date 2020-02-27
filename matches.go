package glaw

import (
	"fmt"
)

type MatchRequest struct {
	ID string
}

func (mr MatchRequest) String() string {
	return fmt.Sprintf("matches/%s", mr.ID)
}

func (c *Client) Match(mr MatchRequest) (matches MatchStorage, err error) {
	return matches, err
}

type MatchesRequest struct {
	AccountID string
}

func (mr MatchesRequest) String() string {
	return fmt.Sprintf("matchlists/by-account/%s", mr.AccountID)
}

func (c *Client) Matches(mr MatchRequest) (matches MatchStorage, err error) {
	return matches, err
}

type TimelineRequest struct {
	ID string
}

func (tr TimelineRequest) String() string {
	return fmt.Sprintf("timelines/by-match/%s", mr.ID)
}

func (c *Client) Timeline(mr MatchRequest) (matches MatchStorage, err error) {
	return matches, err
}
