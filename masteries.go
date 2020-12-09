package glaw

import (
	"context"
	"fmt"
)

type ScoreRequest struct {
	Region              Region
	EncryptedSummonerID string
}

func (sr ScoreRequest) String() string {
	return fmt.Sprintf("scores/by-summoner/%s", sr.EncryptedSummonerID)
}

func (c *Client) ChampionScore(ctx context.Context, mr MasteryRequest) (score int, err error) {
	req := Request{
		Method:  `GET`,
		Domain:  `champion-mastery`,
		Version: V4,
		Region:  mr.Region,
		Uri:     mr.String(),
	}

	r, err := req.NewHttpRequest(ctx)
	if err != nil {
		return score, err
	}

	resp, err := c.Do(r)
	if err != nil {
		return score, err
	}

	err = ProcessResponse(resp, &score)
	return score, err
}

type MasteryRequest struct {
	Region              Region
	EncryptedSummonerID string
}

func (mr MasteryRequest) String() string {
	return fmt.Sprintf("champion-masteries/by-summoner/%s", mr.EncryptedSummonerID)
}

func (c *Client) ChampionMasteries(ctx context.Context, mr MasteryRequest) (cm []ChampionMastery, err error) {
	req := Request{
		Method:  `GET`,
		Domain:  `champion-mastery`,
		Version: V4,
		Region:  mr.Region,
		Uri:     mr.String(),
	}

	r, err := req.NewHttpRequest(ctx)
	if err != nil {
		return cm, err
	}

	resp, err := c.Do(r)
	if err != nil {
		return cm, err
	}

	err = ProcessResponse(resp, &cm)
	return cm, err
}

type MasteriesRequest struct {
	Region              Region
	EncryptedSummonerID string
	ChampionID          int
}

func (mr MasteriesRequest) String() string {
	return fmt.Sprintf("champion-masteries/by-summoner/%s/by-champion/%d", mr.EncryptedSummonerID, mr.ChampionID)
}

func (c *Client) MasteriesByChampionId(ctx context.Context, mr MasteryRequest) (cm ChampionMastery, err error) {
	req := Request{
		Method:  `GET`,
		Domain:  `champion-mastery`,
		Version: V4,
		Region:  mr.Region,
		Uri:     mr.String(),
	}

	r, err := req.NewHttpRequest(ctx)
	if err != nil {
		return cm, err
	}

	resp, err := c.Do(r)
	if err != nil {
		return cm, err
	}

	err = ProcessResponse(resp, &cm)
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
