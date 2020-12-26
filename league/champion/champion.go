package champion

import (
	"context"
	"fmt"

	"github.com/coltiebaby/glaw"
	"github.com/coltiebaby/glaw/league"
	"github.com/coltiebaby/glaw/league/core"
)

type Client struct {
	client *league.Client
}

func New(c *league.Client) *Client {
	return &Client{
		client: c,
	}
}

func NewChampionClient(opts ...glaw.Option) (*Client, error) {
	c, err := league.NewClient(opts...)
	client := &Client{
		client: c,
	}

	return client, err
}

type FreeRotationRequest struct {
	Region glaw.Region
}

func (c *Client) GetFreeRotation(ctx context.Context, fcr FreeRotationRequest) (ci core.ChampionInfo, err error) {
	uri := `champion-rotations`
	req := league.NewRequest("GET", "platform", uri, fcr.Region, glaw.V3)

	err = c.client.Do(ctx, req, &ci)
	return ci, err
}

// Get Score
//
// Get the players total mastery score

type ScoreRequest struct {
	Region              glaw.Region
	EncryptedSummonerID string
}

func (sr ScoreRequest) String() string {
	return fmt.Sprintf("scores/by-summoner/%s", sr.EncryptedSummonerID)
}

func (c *Client) GetScore(ctx context.Context, mr MasteryRequest) (score int, err error) {
	uri := mr.String()
	req := league.NewRequest("GET", "champion-mastery", uri, mr.Region, glaw.V4)

	err = c.client.Do(ctx, req, &score)
	return score, err
}

// Get Masteries
//
// Get all champions in a slice that have the score

type MasteryRequest struct {
	Region              glaw.Region
	EncryptedSummonerID string
}

func (mr MasteryRequest) String() string {
	return fmt.Sprintf("champion-masteries/by-summoner/%s", mr.EncryptedSummonerID)
}

func (c *Client) GetMasteries(ctx context.Context, mr MasteryRequest) (cm []core.ChampionMastery, err error) {
	uri := mr.String()
	req := league.NewRequest("GET", "champion-mastery", uri, mr.Region, glaw.V4)

	err = c.client.Do(ctx, req, &cm)
	return cm, err
}

// Get Mastery by ChampionId
//
// Get a score for a wanted champion

type MasteriesRequest struct {
	Region              glaw.Region
	EncryptedSummonerID string
	ChampionID          int
}

func (mr MasteriesRequest) String() string {
	return fmt.Sprintf("champion-masteries/by-summoner/%s/by-champion/%d", mr.EncryptedSummonerID, mr.ChampionID)
}

func (c *Client) GetMastery(ctx context.Context, mr MasteryRequest) (cm core.ChampionMastery, err error) {
	uri := mr.String()
	req := league.NewRequest("GET", "champion-mastery", uri, mr.Region, glaw.V4)

	err = c.client.Do(ctx, req, &cm)
	return cm, err
}
