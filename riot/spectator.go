package riot

import (
    "fmt"
    "encoding/json"
)

type BannedChampion struct {
    PickedTurn int `json:"pickTurn"`
    ChampionID int `json:"championId"`
    TeamID int `json:"teamId"`
}

type FeaturedGames struct {
	ClientRefreshInterval int `json:"clientRefreshInterval"`
    Games []Game `json:"gameList"`
}

type Game struct {
		GameID            int64  `json:"gameId"`
		GameStartTime     int64  `json:"gameStartTime"`
		PlatformID        string `json:"platformId"`
		GameMode          string `json:"gameMode"`
		MapID             int    `json:"mapId"`
		GameType          string `json:"gameType"`
		GameQueueConfigID int    `json:"gameQueueConfigId"`
		Observers         Observer `json:"observers"`
		Participants []Participant `json:"participants"`
		GameLength      int           `json:"gameLength"`
		BannedChampions []BannedChampion `json:"bannedChampions"`
}

type Observer struct {
    EncryptionKey string `json:"encryptionKey"`
}

type Perks struct {
    PerkStyle    int   `json:"perkStyle"`
    PerkIds      []int `json:"perkIds"`
    PerkSubStyle int   `json:"perkSubStyle"`
}

type Participant struct {
    ProfileIconID int    `json:"profileIconId"`
    ChampionID    int    `json:"championId"`
    SummonerName  string `json:"summonerName"`
    SummonerID int `json:"summonerId"`
    Bot           bool   `json:"bot"`
    Spell2ID      int    `json:"spell2Id"`
    TeamID        int    `json:"teamId"`
    Spell1ID      int    `json:"spell1Id"`
    Perks         Perks  `json:"perks"`
    GameCustomizationObjects []GameCustomizationObject `json:"gameCustomizationObjects"`
}

type GameCustomizationObject struct {
    category string `json:"category"`
    content string `json:"content"`
}

func GetFeaturedGames() (FeaturedGames) {
    // https://na1.api.riotgames.com/lol/spectator/v3/featured-games
    var featured FeaturedGames

    rr := &RiotRequest {
        Type: "spectator",
        Uri:  "featured-games",
    }

    resp := rr.GetData()
    json.Unmarshal(resp, &featured)

    return featured
}

func GetCurrentGame(summoner Summoner) (Game) {
    // /lol/spectator/v3/active-games/by-summoner/{summonerId}

    var game Game

    uri := fmt.Sprintf("active-games/by-summoner/%d", summoner.ID)
    rr := &RiotRequest {
        Type: "spectator",
        Uri:  uri,
    }

    resp := rr.GetData()
    json.Unmarshal(resp, &game)

    return game
}
