package core

type Featured struct {
	Games                 []FeaturedGame `json:"gameList"`
	ClientRefreshInterval int            `json:"clientRefreshInterval"`
}

type FeaturedGame struct {
	GameID            int64         `json:"gameId"`
	MapID             int           `json:"mapId"`
	GameMode          string        `json:"gameMode"`
	GameType          string        `json:"gameType"`
	GameQueueConfigID int           `json:"gameQueueConfigId"`
	Participants      []Participant `json:"participants"`
	Observers         Observer      `json:"observers"`
	PlatformID        string        `json:"platformId"`
	GameTypeConfigID  int           `json:"gameTypeConfigId"`
	BannedChampions   []Banned      `json:"bannedChampions"`
	GameStartTime     int64         `json:"gameStartTime"`
	GameLength        int           `json:"gameLength"`
}

type Observer struct {
	EncryptionKey string `json:"encryptionKey"`
}
