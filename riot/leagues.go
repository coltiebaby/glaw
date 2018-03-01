package riot

import (
    "fmt"
    "encoding/json"
    "strings"
)

type Entry struct {
    HotStreak        bool   `json:"hotStreak"`
    Wins             int    `json:"wins"`
    Veteran          bool   `json:"veteran"`
    Losses           int    `json:"losses"`
    Rank             string `json:"rank"`
    PlayerOrTeamName string `json:"playerOrTeamName"`
    Inactive         bool   `json:"inactive"`
    PlayerOrTeamID   string `json:"playerOrTeamId"`
    FreshBlood       bool   `json:"freshBlood"`
    LeaguePoints     int    `json:"leaguePoints"`
}

type PlayerLeague struct {
	QueueType        string `json:"queueType"`
	HotStreak        bool   `json:"hotStreak"`
	Wins             int    `json:"wins"`
	Veteran          bool   `json:"veteran"`
	Losses           int    `json:"losses"`
	PlayerOrTeamID   string `json:"playerOrTeamId"`
	LeagueName       string `json:"leagueName"`
	PlayerOrTeamName string `json:"playerOrTeamName"`
	Inactive         bool   `json:"inactive"`
	Rank             string `json:"rank"`
	FreshBlood       bool   `json:"freshBlood"`
	LeagueID         string `json:"leagueId"`
	Tier             string `json:"tier"`
	LeaguePoints     int    `json:"leaguePoints"`
}

type League struct {
	Tier     string  `json:"tier"`
	Queue    string  `json:"queue"`
	LeagueID string  `json:"leagueId"`
	Name     string  `json:"name"`
	Entries  []Entry `json:"entries"`
}

func selectQueue(queue_type string) string {
    switch qt := strings.ToLower(queue_type); qt {
        case "solo":
            return "RANKED_SOLO_5x5"
        case "flex":
            return "RANKED_FLEX_SR"
        case "tt":
            return "RANKED_FLEX_TT"
        default:
            return "RANKED_SOLO_5x5"
    }

}

func GetQueue(league_type string, queue_type string) (League) {
    // /lol/league/v3/challengerleagues/by-queue/RANKED_SOLO_5x5
    var league League

    uri := "%sleagues/by-queue/%s"
    rr := &RiotRequest {
        Type: "league",
        Uri:  fmt.Sprintf(uri, strings.ToLower(league_type), selectQueue(queue_type)),
    }

    resp := rr.GetData()
    json.Unmarshal(resp, &league)

    return league
}

func GetSummonerLeagues(summoner Summoner) ([]PlayerLeague) {
    // /lol/league/v3/challengerleagues/by-queue/RANKED_SOLO_5x5
    player_leagues := make([]PlayerLeague, 0)

    uri := "positions/by-summoner/%d"
    rr := &RiotRequest {
        Type: "league",
        Uri:  fmt.Sprintf(uri, summoner.ID),
    }

    resp := rr.GetData()
    json.Unmarshal(resp, &player_leagues)

    return player_leagues
}

func GetLeague(league_id string) (League) {
    var league League

    uri := "leagues/%s"
    rr := &RiotRequest {
        Type: "league",
        Uri:  fmt.Sprintf(uri, league_id),
    }

    resp := rr.GetData()
    json.Unmarshal(resp, &league)

    return league
}
