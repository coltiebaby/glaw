package riot

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/coltiebaby/g-law/config"
	log "github.com/sirupsen/logrus"
)

var (
	Client = &http.Client{}
	c      = config.GetConfig()
)

func logResponse(code int, url string) {
	respLog := log.WithFields(log.Fields{name: "riot", "url": url, "status": code})
	switch code {
	case 200, 201:
		respLog.Debug()
	case 404:
		respLog.Warning("Not Found!")
	case 401:
		respLog.Fatal("API Token is down for the count...")
	default:
		respLog.Warning()
	}
}

type RiotRequest struct {
	Type    string
	Uri     string
	Version string
	Params  map[string]string
}

func (rr RiotRequest) GetData() (*http.Response, error) {
	values := url.Values{}
	for k, v := range rr.Params {
		values.Add(k, v)
	}

	u := &url.URL{
		Scheme:   "https",
		Host:     "na1.api.riotgames.com",
		Path:     fmt.Sprintf("lol/%s/%s/%s", rr.Type, rr.Version, rr.Uri),
		RawQuery: values.Encode(),
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		RiotLog.Fatal(err)
		return &http.Response{}, err
	}

	req.Header.Add("X-Riot-Token", c.Api.Token)
	resp, err := Client.Do(req)
	if err != nil {
		RiotLog.Fatal(err)
		return &http.Response{}, err
	}

	logResponse(resp.StatusCode, u.String())

	return resp, nil
}
