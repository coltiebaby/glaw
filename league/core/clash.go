package core

import (
	"encoding/json"
	"time"
)

type ClashPlayer struct {
	SummonerID string `json:"summonerID"`
	TeamID     string `json:"teamID"`
	Position   string `json:"postion"` // (Legal values: UNSELECTED, FILL, TOP, JUNGLE, MIDDLE, BOTTOM, UTILITY)
	Role       string `json:"role"`    // (Legal values: CAPTAIN, MEMBER)
}

type ClashTeam struct {
	ID           string        `json:"id"`
	TournamentId int           `json:"tournamentId"`
	Name         string        `json:"name"`
	IconID       int           `json:"iconId"`
	Tier         int           `json:"tier"`
	Captain      string        `json:"captain"`
	Abbreviation string        `json:"abbreviation"`
	Players      []ClashPlayer `json:"players"`
}

type ClashTournament struct {
	ID               int                    `json:"id"`
	ThemeID          int                    `json:"themeId"`
	NameKey          string                 `json:"nameKey"`
	NameKeySecondary string                 `json:"nameKeySecondary"`
	Schedule         []ClashTournamentPhase `json:"schedule"`
}

type ClashTournamentPhase struct {
	ID               int       `json:"id"`
	RegistrationTime time.Time `json:"registrationTime"`
	StartTime        time.Time `json:"startTime"`
	Cancelled        bool      `json:"cancelled"`
}

type phase struct {
	ID               int   `json:"id"`
	RegistrationTime int64 `json:"registrationTime"`
	StartTime        int64 `json:"startTime"`
	Cancelled        bool  `json:"cancelled"`
}

func (ct *ClashTournamentPhase) UnmarshalJSON(b []byte) error {
	var p phase
	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	ctp := ClashTournamentPhase{
		ID:               p.ID,
		RegistrationTime: time.Unix(p.RegistrationTime, 0),
		StartTime:        time.Unix(p.StartTime, 0),
		Cancelled:        p.Cancelled,
	}

	*ct = ctp

	return nil
}
