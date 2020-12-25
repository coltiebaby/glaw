package core

type Spectator string

const (
	NONE      Spectator = "NONE"
	LOBBYONLY Spectator = "LOBBYONLY"
	ALL       Spectator = "ALL"
)

type Map string

const (
	SUMMONERS_RIFT Map = "SUMMONERS_RIFT"
	TREELINE       Map = "TWISTED_TREELINE"
	HOWLING_ABYSS  Map = "HOWLING_ABYSS"
)

type GameType string

const (
	BLIND_PICK       GameType = "BLIND_PICK"
	DRAFT_MODE       GameType = "DRAFT_MODE"
	ALL_RANDOM       GameType = "ALL_RANDOM"
	TOURNAMENT_DRAFT GameType = "TOURNAMENT_DRAFT"
)

type Queue string

const (
	FLEX               Queue = `RANKED_FLEX`
	SOLO               Queue = `RANKED_SOLO_5x5`
	TEAM_FIGHT_TACTICS Queue = `RANKED_FLEX_TFT`
	TWISTED_TREELINE   Queue = `RANKED_FLEX_TT`
)

type Tier string

const (
	CHALLENGER  Tier = `CHALLENGER`
	MASTER      Tier = `MASTER`
	GRANDMASTER Tier = `GRANDMASTER`
	DIAMOND     Tier = `DIAMOND`
	PLATINUM    Tier = `PLATINUM`
	GOLD        Tier = `GOLD`
	SILVER      Tier = `SILVER`
	BRONZE      Tier = `BRONZE`
	IRON        Tier = `IRON`
)

type Division string

const (
	ONE   Division = `I`
	TWO   Division = `II`
	THREE Division = `III`
	FOUR  Division = `IV`
)

type Entry struct {
	MiniSeries   MiniSeries
	SummonerName string `json:"summonerName"`
	HotStreak    bool   `json:"hotStreak"`
	Wins         int    `json:"wins"`
	Veteran      bool   `json:"veteran"`
	Losses       int    `json:"losses"`
	Rank         string `json:"rank"`
	Inactive     bool   `json:"inactive"`
	FreshBlood   bool   `json:"freshBlood"`
	SummonerID   string `json:"summonerId"`
	LeaguePoints int    `json:"leaguePoints"`
}

type MiniSeries struct {
	Progress string `json:"progress"`
	Losses   int    `json:"losses"`
	Target   int    `json:"target"`
	Wins     int    `json:"wins"`
}

type League struct {
	Tier     string  `json:"tier"`
	LeagueID string  `json:"leagueId"`
	Entries  []Entry `json:"entries"`
	Queue    string  `json:"queue"`
	Name     string  `json:"name"`
}

// Has slightly more info than the Entry struct
type LeagueEntry struct {
	QueueType string `json:"queueType"`
	Tier      string `json:"tier"`
	Entry
}
