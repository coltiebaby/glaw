package core

import (
	"encoding/json"
	"time"
)

type Queue string

const (
	Competitive Queue = "competitive"
	Unrank      Queue = "unrank"
	SpikeRush   Queue = "spikerush"
)

type RecentMatches struct {
	CurrentTime time.Time `json:"currentTime"`
	Ids         []string
}

func (recent *RecentMatches) UnmarshalJSON(b []byte) error {
	type phase struct {
		CurrentTime int64 `json:"currentTime"`
	}

	var p phase
	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	var r RecentMatches
	if err := json.Unmarshal(b, &r); err != nil {
		return err
	}

	r.CurrentTime = time.Unix(p.CurrentTime, 0)

	*recent = r

	return nil
}

type Matchlist struct {
	PUUID   string
	Entries []Entry
}

type Entry struct {
	MatchId   string
	TeamId    string
	StartTime time.Time
}

func (entry *Entry) UnmarshalJSON(b []byte) error {
	type phase struct {
		StartTime int64 `json:"startTime"`
	}

	var p phase
	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	var e Entry
	if err := json.Unmarshal(b, &e); err != nil {
		return err
	}

	e.StartTime = time.Unix(p.StartTime, 0)

	*entry = e

	return nil
}

type Match struct {
	MatchInfo    MatchInfo     `json:"matchInfo"`
	Players      []Player      `json:"players"`
	Teams        []Team        `json:"teams"`
	RoundResults []RoundResult `json:"roundResults"`
}

type MatchInfo struct {
	MatchId            string `json:"matchId"`
	MapId              string `json:"mapId"`
	GameLength         time.Duration
	GameStart          time.Time
	ProvisioningFlowId string `json:"provisioningFlowId"`
	IsCompleted        bool   `json:"isCompleted"`
	CustomGameName     string `json:"customGameName"`
	QueueId            string `json:"queueId"`
	GameMode           string `json:"gameMode"`
	IsRanked           bool   `json:"isRanked"`
	SeasonId           string `json:"seasonId"`
}

func (mi *MatchInfo) UnmarshalJSON(b []byte) error {
	type phase struct {
		GameLength int64 `json:"gameLengthMillis"`
		GameStart  int64 `json:"gameStartMillis"`
	}

	var p phase
	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	var m MatchInfo
	if err := json.Unmarshal(b, &m); err != nil {
		return err
	}

	m.GameLength = time.Duration(p.GameLength) * time.Second
	m.GameStart = time.Unix(p.GameStart, 0)

	*mi = m

	return nil
}

type Player struct {
	Puuid           string     `json:"puuid"`
	TeamId          string     `json:"teamId"`
	PartyId         string     `json:"partyId"`
	CharacterId     string     `json:"characterId"`
	Stats           PlayerStat `json:"stats"`
	CompetitiveTier int        `json:"competitiveTier"`
	PlayerCard      string     `json:"playerCard"`
	PlayerTitle     string     `json:"playerTitle"`
}

type PlayerStat struct {
	Score        int           `json:"score"`
	RoundsPlayed int           `json:"roundsPlayed"`
	Kills        int           `json:"kills"`
	Deaths       int           `json:"deaths"`
	Assists      int           `json:"assists"`
	Playtime     time.Duration `json:"playtimeMillis"`
	AbilityCasts AbilityCast   `json:"abilityCasts"`
}

func (ps *PlayerStat) UnmarshalJSON(b []byte) error {
	type phase struct {
		Playtime int64 `json:"playtimeMillis"`
	}

	var p phase
	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	var stats PlayerStat
	if err := json.Unmarshal(b, &ps); err != nil {
		return err
	}

	stats.Playtime = time.Duration(p.Playtime) * time.Millisecond

	*ps = stats

	return nil
}

type AbilityCast struct {
	GrenadeCasts  int `json:"grenadeCasts"`
	Ability1Casts int `json:"ability1Casts"`
	Ability2Casts int `json:"ability2Casts"`
	UltimateCasts int `json:"ultimateCasts"`
}

type Team struct {
	TeamId       string `json:"teamId"`
	Won          bool   `json:"won"`
	RoundsPlayed int    `json:"roundsPlayed"`
	RoundsWon    int    `json:"roundsWon"`
	NumPoints    int    `json:"numPoints"`
}

type RoundResult struct {
	RoundNum              int              `json:"roundNum"`
	RoundResult           string           `json:"roundResult"`
	RoundCeremony         string           `json:"roundCeremony"`
	WinningTeam           string           `json:"winningTeam"`
	BombPlanter           string           `json:"bombPlanter"`
	BombDefuser           string           `json:"bombDefuser"`
	PlantRoundTime        int              `json:"plantRoundTime"`
	PlantPlayerLocations  []PlayerLocation `json:"plantPlayerLocations"`
	PlantLocation         Location         `json:"plantLocation"`
	PlantSite             string           `json:"plantSite"`
	DefuseRoundTime       int              `json:"defuseRoundTime"`
	DefusePlayerLocations []PlayerLocation `json:"defusePlayerLocations"`
	DefuseLocation        Location         `json:"defuseLocation"`
	PlayerStats           []PlayerStat     `json:"playerStats"`
	RoundResultCode       string           `json:"roundResultCode"`
}

type PlayerLocation struct {
	Puuid       string   `json:"puuid"`
	ViewRadians float64  `json:"viewRadians"`
	Location    Location `json:"location"`
}

type Location struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type PlayerRoundStat struct {
	Puuid   string   `json:"puuid"`
	Kills   []Kill   `json:"kills"`
	Damage  []Damage `json:"damage"`
	Score   int      `json:"score"`
	Economy Economy  `json:"economy"`
	Ability Ability  `json:"ability"`
}

type Kill struct {
	TimeSinceGameStart  time.Duration
	TimeSinceRoundStart time.Duration
	Killer              string           `json:"killer"`
	Victim              string           `json:"victim"`
	VictimLocation      Location         `json:"victimLocation"`
	Assistants          []string         `json:"assistants"`
	PlayerLocations     []PlayerLocation `json:"playerLocations"`
	FinishingDamage     FinishingDamage  `json:"finishingDamage"`
}

func (kill *Kill) UnmarshalJSON(b []byte) error {
	type phase struct {
		TimeSinceGameStart  int64 `json:"timeSinceGameStartMillis"`
		TimeSinceRoundStart int64 `json:"timeSinceRoundStartMillis"`
	}

	var p phase
	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}

	var k Kill
	if err := json.Unmarshal(b, &k); err != nil {
		return err
	}

	k.TimeSinceGameStart = time.Duration(p.TimeSinceGameStart) * time.Millisecond
	k.TimeSinceRoundStart = time.Duration(p.TimeSinceRoundStart) * time.Millisecond

	*kill = k

	return nil
}

type FinishingDamage struct {
	DamageType          string `json:"damageType"`
	DamageItem          string `json:"damageItem"`
	IsSecondaryFireMode bool   `json:"isSecondaryFireMode"`
}

type Damage struct {
	Receiver  string `json:"receiver"`
	Damage    int    `json:"damage"`
	Legshots  int    `json:"legshots"`
	Bodyshots int    `json:"bodyshots"`
	Headshots int    `json:"headshots"`
}

type Economy struct {
	LoadoutValue int    `json:"loadoutValue"`
	Weapon       string `json:"weapon"`
	Armor        string `json:"armor"`
	Remaining    int    `json:"remaining"`
	Spent        int    `json:"spent"`
}

type Ability struct {
	GrenadeEffects  string `json:"grenadeEffects"`
	Ability1Effects string `json:"ability1Effects"`
	Ability2Effects string `json:"ability2Effects"`
	UltimateEffects string `json:"ultimateEffects"`
}
