package core

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
	SummonerID   string `json:"summonerId"`
	SummonerName string `json:"summonerName"`
	LeaguePoints int    `json:"leaguePoints"`
	Rank         string `json:"rank"`
	Wins         int    `json:"wins"`
	Losses       int    `json:"losses"`
	Veteran      bool   `json:"veteran"`
	Inactive     bool   `json:"inactive"`
	FreshBlood   bool   `json:"freshBlood"`
	HotStreak    bool   `json:"hotStreak"`
}

type Rank struct {
	Tier     string  `json:"tier"`
	LeagueID string  `json:"leagueId"`
	Queue    string  `json:"queue"`
	Name     string  `json:"name"`
	Entries  []Entry `json:"entries"`
}
