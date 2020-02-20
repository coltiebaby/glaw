package mastery

import (
	"fmt"

	"github.com/coltiebaby/glaw"
	"github.com/coltiebaby/glaw/v4"
)

type MasteryRequest struct {
    Region Region
    EncryptedSummonerId string
    ChampionId int
}

func newMasteryRequest(query string) Request {
    return Request {
        Method: `GET`
        Domain: `champion-mastery`
        Version: V4,
        Uri: query,
    }
}

func (c *Client) ChampionScore(ctx context.Context, mr MasteryRequest) (score int, err error) {
    req := newMasteryRequest(fmt.Sprintf(`scores/by-summoner/%s`, mr.EncryptedSummonerId))
    req.Region = mr.Region

    resp, err := c.Do(req.NewHttpRequestWithCtx(ctx))
    if err != nil {
        return ci, err
    }

    err = ProcessRequest(resp, &score)

    return score, err
}

func (c *Client) ChampionMasteries(c glaw.ApiClient, mr MasteryRequest) (cm []ChampionMastery, err error) {
    req := newMasteryRequest(fmt.Sprintf(`champion-masteries/by-summoner/%s`, mr.EncryptedSummonerId))
    req.Region = mr.Region

    resp, err := c.Do(req.NewHttpRequestWithCtx(ctx))
    if err != nil {
        return ci, err
    }

    err = ProcessRequest(resp, &cm)

    return cm, err
}

func (c *Client) MasteriesByChampionId(c glaw.ApiClient, mr MasteryRequest) (cm ChampionMastery, err error) {
    template := `champion-masteries/by-summoner/%s/by-champion/%d`
    req := newMasteryRequest(fmt.Sprintf(template, mr.EncryptedSummonerId, mr.ChampionId))
    req.Region = mr.Region

    resp, err := c.Do(req.NewHttpRequestWithCtx(ctx))
    if err != nil {
        return ci, err
    }

    err = ProcessRequest(resp, &cm)
	return cm, err
}
