package glaw

import (
	"fmt"
)

type SummonerRequest struct {
	Type int
	ID   string
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

func (c *Client) Summoner() (summoner Summoner, err error) {

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
