package glaw

import (
	"fmt"
)

type ScoreRequest struct {
	Region              Region
	EncryptedSummonerId string
}

func (sr ScoreRequest) String() string {
	return fmt.Sprintf("scores/by-summoner/%s", mr.EncryptedSummonerID)
}

func (c *Client) ChampionScore(ctx context.Context, mr MasteryRequest) (score int, err error) {
	req := Request{
		Method:  `GET`,
		Domain:  `champion-mastery`,
		Version: V4,
		Region:  lr.Region,
		Uri:     lr.String(),
	}

	resp, err := c.Do(req.NewHttpRequestWithCtx(ctx))
	if err != nil {
		return score, err
	}

	err = ProcessRequest(resp, &score)
	return score, err
}

type MasteryRequest struct {
	Region              Region
	EncryptedSummonerID string
}

func (mr MasteryRequest) String() string {
	return fmt.Sprintf("champion-masteries/by-summoner/%s", mr.EncryptedSummonerID)
}

func (c *Client) ChampionMasteries(c glaw.ApiClient, mr MasteryRequest) (cm []ChampionMastery, err error) {
	req := Request{
		Method:  `GET`,
		Domain:  `champion-mastery`,
		Version: V4,
		Region:  lr.Region,
		Uri:     lr.String(),
	}

	resp, err := c.Do(req.NewHttpRequestWithCtx(ctx))
	if err != nil {
		return cm, err
	}

	err = ProcessRequest(resp, &cm)
	return cm, err
}

type MasteriesRequest struct {
	Region              Region
	EncryptedSummonerID string
	ChampionID          int
}

func (mr MasteriesRequest) String() string {
	return fmt.Sprintf("champion-masteries/by-summoner/%s/by-champion/%s", EncryptedSummonerID, ChampionID)
}

func (c *Client) MasteriesByChampionId(mr MasteryRequest) (cm ChampionMastery, err error) {
	req := Request{
		Method:  `GET`,
		Domain:  `champion-mastery`,
		Version: V4,
		Region:  mr.Region,
		Uri:     mr.String(),
	}

	resp, err := c.Do(req.NewHttpRequestWithCtx(ctx))
	if err != nil {
		return cm, err
	}

	err = ProcessRequest(resp, &cm)
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
