package league

import (
	"context"
	"fmt"

	"github.com/coltiebaby/glaw"
	"github.com/coltiebaby/glaw/league/core"
)

type ScoreRequest struct {
	Region              glaw.Region
	EncryptedSummonerID string
}

func (sr ScoreRequest) String() string {
	return fmt.Sprintf("scores/by-summoner/%s", sr.EncryptedSummonerID)
}

func (c *Client) ChampionScore(ctx context.Context, mr MasteryRequest) (score int, err error) {
	uri := mr.String()
	req := NewRequest("GET", "champion-mastery", uri, mr.Region, glaw.V4)

	err = c.Do(ctx, req, &score)
	return score, err
}

type MasteryRequest struct {
	Region              glaw.Region
	EncryptedSummonerID string
}

func (mr MasteryRequest) String() string {
	return fmt.Sprintf("champion-masteries/by-summoner/%s", mr.EncryptedSummonerID)
}

func (c *Client) ChampionMasteries(ctx context.Context, mr MasteryRequest) (cm []core.ChampionMastery, err error) {
	uri := mr.String()
	req := NewRequest("GET", "champion-mastery", uri, mr.Region, glaw.V4)

	err = c.Do(ctx, req, &cm)
	return cm, err
}

type MasteriesRequest struct {
	Region              glaw.Region
	EncryptedSummonerID string
	ChampionID          int
}

func (mr MasteriesRequest) String() string {
	return fmt.Sprintf("champion-masteries/by-summoner/%s/by-champion/%d", mr.EncryptedSummonerID, mr.ChampionID)
}

func (c *Client) MasteriesByChampionId(ctx context.Context, mr MasteryRequest) (cm core.ChampionMastery, err error) {
	uri := mr.String()
	req := NewRequest("GET", "champion-mastery", uri, mr.Region, glaw.V4)

	err = c.Do(ctx, req, &cm)
	return cm, err
}
