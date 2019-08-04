package matches

type Player struct {
	CurrentPlatformID string `json:"currentPlatformId"`
	SummonerName      string `json:"summonerName"`
	MatchHistoryURI   string `json:"matchHistoryUri"`
	PlatformID        string `json:"platformId"`
	CurrentAccountID  string `json:"currentAccountId"`
	ProfileIcon       int    `json:"profileIcon"`
	SummonerID        string `json:"summonerId"`
	AccountID         string `json:"accountId"`
}

type ParticipantID struct {
	Player        Player `json:"player"`
	ParticipantID int    `json:"participantId"`
}

type Team struct {
	FirstDragon          bool          `json:"firstDragon"`
	Bans                 []interface{} `json:"bans"`
	FirstInhibitor       bool          `json:"firstInhibitor"`
	Win                  string        `json:"win"`
	FirstRiftHerald      bool          `json:"firstRiftHerald"`
	FirstBaron           bool          `json:"firstBaron"`
	BaronKills           int           `json:"baronKills"`
	RiftHeraldKills      int           `json:"riftHeraldKills"`
	FirstBlood           bool          `json:"firstBlood"`
	TeamID               int           `json:"teamId"`
	FirstTower           bool          `json:"firstTower"`
	VilemawKills         int           `json:"vilemawKills"`
	InhibitorKills       int           `json:"inhibitorKills"`
	TowerKills           int           `json:"towerKills"`
	DominionVictoryScore int           `json:"dominionVictoryScore"`
	DragonKills          int           `json:"dragonKills"`
}

type Match struct {
	SeasonID              int             `json:"seasonId"`
	QueueID               int             `json:"queueId"`
	GameID                int64           `json:"gameId"`
	ParticipantIdentities []ParticipantID `json:"participantIdentities"`
	GameVersion           string          `json:"gameVersion"`
	PlatformID            string          `json:"platformId"`
	GameMode              string          `json:"gameMode"`
	MapID                 int             `json:"mapId"`
	GameType              string          `json:"gameType"`
	Teams                 []Team          `json:"teams"`
	Participants          []Participant   `json:"participants"`
	GameDuration          int             `json:"gameDuration"`
	GameCreation          int64           `json:"gameCreation"`
}
