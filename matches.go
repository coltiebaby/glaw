package glaw

import (
	"context"
	"fmt"
)

type MatchRequest struct {
	ID     string
	Region Region
}

func (mr MatchRequest) String() string {
	return fmt.Sprintf("matches/%s", mr.ID)
}

func (c *Client) Match(ctx context.Context, mr MatchRequest) (matches MatchStorage, err error) {
	req := Request{
		Method:  `GET`,
		Domain:  `match`,
		Version: V4,
		Region:  mr.Region,
		Uri:     mr.String(),
	}

	resp, err := c.Do(ctx, req)
	if err != nil {
		return matches, err
	}

	err = ProcessResponse(resp, &matches)
	return matches, err
}

type MatchesRequest struct {
	AccountID string
	Region    Region
}

func (mr MatchesRequest) String() string {
	return fmt.Sprintf("matchlists/by-account/%s", mr.AccountID)
}

func (c *Client) Matches(ctx context.Context, mr MatchRequest) (matches MatchStorage, err error) {
	req := Request{
		Method:  `GET`,
		Domain:  `match`,
		Version: V4,
		Region:  mr.Region,
		Uri:     mr.String(),
	}

	resp, err := c.Do(ctx, req)
	if err != nil {
		return matches, err
	}

	err = ProcessResponse(resp, &matches)
	return matches, err
}

type TimelineRequest struct {
	ID     string
	Region Region
}

func (tr TimelineRequest) String() string {
	return fmt.Sprintf("timelines/by-match/%s", tr.ID)
}

func (c *Client) Timeline(ctx context.Context, mr MatchRequest) (matches MatchStorage, err error) {
	req := Request{
		Method:  `GET`,
		Domain:  `match`,
		Version: V4,
		Region:  mr.Region,
		Uri:     mr.String(),
	}

	resp, err := c.Do(ctx, req)
	if err != nil {
		return matches, err
	}

	err = ProcessResponse(resp, &matches)
	return matches, err
}