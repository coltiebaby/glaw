package glaw

import (
	"context"
	"fmt"
)

type SummonerRequest struct {
	Type   int
	ID     string
	Region Region
}

func (sr SummonerRequest) String() string {
	uri := sr.ID

	switch sr.Type {
	case SummonerID:
	case SummonerName:
		uri = fmt.Sprintf("by-name/%s", uri)
	case SummonerPUUID:
		uri = fmt.Sprintf("by-puuid/%s", uri)
	case SummonerAccountID:
		uri = fmt.Sprintf("by-account", uri)
	}

	return uri
}

func (c *Client) Summoner(ctx context.Context, sr SummonerRequest) (summoner Summoner, err error) {
	req := Request{
		Method:  `GET`,
		Domain:  `summoner`,
		Version: V4,
		Region:  sr.Region,
		Uri:     sr.String(),
	}

	r, err := req.NewHttpRequestWithCtx(ctx)
	if err != nil {
		return summoner, err
	}

	resp, err := c.Do(r)
	if err != nil {
		return summoner, err
	}

	err = ProcessResponse(resp, &summoner)
	return summoner, err
}

type Summoner struct {
	ID            string `json:"id"`
	AccountID     string `json:"accountId"`
	PUUID         string `json:"puuid"`
	Name          string `json:"name"`
	SummonerLevel int    `json:"summonerLevel"`
	ProfileIconID int    `json:"profileIconId"`
	// Date summoner was last modified specified as epoch milliseconds.
	// The following events will update this timestamp: profile icon change,
	// playing the tutorial or advanced tutorial, finishing a game, summoner name change
	RevisionDate int64 `json:"revisionDate"`
}

const (
	SummonerName = iota
	SummonerID
	SummonerAccountID
	SummonerPUUID
)
