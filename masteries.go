package glaw

import (
	"fmt"
)

type MasteryRequest struct {
	Region              Region
	EncryptedSummonerId string
	ChampionId          int
}

func newMasteryRequest(query string) Request {
	return Request{
		Method:  `GET`,
		Domain:  `champion-mastery`,
		Version: V4,
		Uri:     query,
	}
}

func (c *Client) ChampionScore(ctx context.Context, mr MasteryRequest) (score int, err error) {
	return score, err
}

func (c *Client) ChampionMasteries(c glaw.ApiClient, mr MasteryRequest) (cm []ChampionMastery, err error) {
	return cm, err
}

func (c *Client) MasteriesByChampionId(c glaw.ApiClient, mr MasteryRequest) (cm ChampionMastery, err error) {
	return cm, err
}

type ChampionMastery struct {
	ChampionLevel                int    `json:"championLevel"`
	ChestGranted                 bool   `json:"chestGranted"`
	ChampionPoints               int    `json:"championPoints"`
	ChampionPointsSinceLastLevel int    `json:"championPointsSinceLastLevel"`
	ChampionPointsUntilNextLevel int    `json:"championPointsUntilNextLevel"`
	SummonerID                   string `json:"summonerId"`
	TokensEarned                 int    `json:"tokensEarned"`
	ChampionID                   int    `json:"championId"`
	LastPlayTime                 int64  `json:"lastPlayTime"`
}
