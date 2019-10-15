package matches

type MatchStorage struct {
	Matches    []MatchInfo `json:"matches"`
	EndIndex   int         `json:"endIndex"`
	StartIndex int         `json:"startIndex"`
	TotalGames int         `json:"totalGames"`
}

type MatchInfo struct {
	Lane       string `json:"lane"`
	GameID     int64  `json:"gameId"`
	Champion   int    `json:"champion"`
	PlatformID string `json:"platformId"`
	Timestamp  int64  `json:"timestamp"`
	Queue      int    `json:"queue"`
	Role       string `json:"role"`
	Season     int    `json:"season"`
}
