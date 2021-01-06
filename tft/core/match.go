package core

type Metadata struct {
	DataVersion  string   `json:"data_version"`
	MatchID      string   `json:"match_id"`
	Participants []string `json:"participants"`
}

type Companion struct {
	ContentID string `json:"content_ID"`
	SkinID    int    `json:"skin_ID"`
	Species   string `json:"species"`
}

type Trait struct {
	Name        string `json:"name"`
	NumUnits    int    `json:"num_units"`
	Style       int    `json:"style"`
	TierCurrent int    `json:"tier_current"`
	TierTotal   int    `json:"tier_total"`
}

type Unit struct {
	CharacterID string `json:"character_id"`
	Items       []int  `json:"items"`
	Name        string `json:"name"`
	Rarity      int    `json:"rarity"`
	Tier        int    `json:"tier"`
	Chosen      string `json:"chosen,omitempty"`
}

type Participant struct {
	Companion            Companion `json:"companion"`
	GoldLeft             int       `json:"gold_left"`
	LastRound            int       `json:"last_round"`
	Level                int       `json:"level"`
	Placement            int       `json:"placement"`
	PlayersEliminated    int       `json:"players_eliminated"`
	Puuid                string    `json:"puuid"`
	TimeEliminated       float64   `json:"time_eliminated"`
	TotalDamageToPlayers int       `json:"total_damage_to_players"`
	Traits               []Trait   `json:"traits"`
	Units                []Unit    `json:"units"`
}

type Match struct {
	GameDatetime int64         `json:"game_datetime"`
	GameLength   float64       `json:"game_length"`
	GameVersion  string        `json:"game_version"`
	Participants []Participant `json:"participants"`
	QueueID      int           `json:"queue_id"`
	TftSetNumber int           `json:"tft_set_number"`
}

type MatchStorage struct {
	Metadata Metadata `json:"metadata"`
	Match    Match    `json:"info"`
}
