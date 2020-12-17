package core

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
