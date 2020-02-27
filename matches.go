package glaw

import (
	"fmt"
)

type MatchRequest struct {
	ID     string
	Region Region
}

func (mr MatchRequest) String() string {
	return fmt.Sprintf("matches/%s", mr.ID)
}

func (c *Client) Match(mr MatchRequest) (matches MatchStorage, err error) {
	req := Request{
		Method:  `GET`,
		Domain:  `match`,
		Version: V4,
		Region:  mr.Region,
		Uri:     mr.String(),
	}

	resp, err := c.Do(req.NewHttpRequestWithCtx(ctx))
	if err != nil {
		return matches, err
	}

	err = ProcessRequest(resp, &matches)
	return matches, err
}

type MatchesRequest struct {
	AccountID string
	Region    Region
}

func (mr MatchesRequest) String() string {
	return fmt.Sprintf("matchlists/by-account/%s", mr.AccountID)
}

func (c *Client) Matches(mr MatchRequest) (matches MatchStorage, err error) {
	req := Request{
		Method:  `GET`,
		Domain:  `match`,
		Version: V4,
		Region:  mr.Region,
		Uri:     mr.String(),
	}

	resp, err := c.Do(req.NewHttpRequestWithCtx(ctx))
	if err != nil {
		return matches, err
	}

	err = ProcessRequest(resp, &matches)
	return matches, err
}

type TimelineRequest struct {
	ID     string
	Region Region
}

func (tr TimelineRequest) String() string {
	return fmt.Sprintf("timelines/by-match/%s", mr.ID)
}

func (c *Client) Timeline(mr MatchRequest) (matches MatchStorage, err error) {
	req := Request{
		Method:  `GET`,
		Domain:  `match`,
		Version: V4,
		Region:  mr.Region,
		Uri:     mr.String(),
	}

	resp, err := c.Do(req.NewHttpRequestWithCtx(ctx))
	if err != nil {
		return matches, err
	}

	err = ProcessRequest(resp, &matches)
	return matches, err
}
