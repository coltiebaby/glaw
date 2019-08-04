package riot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/coltiebaby/g-law/config"
)

var (
	Client = &http.Client{}
	// TODO: Replace
	c = func() *config.Config {
		x := config.NewConfig()
		x.FromEnv()
		return x

	}()
)

type RiotRequest struct {
	Type    string
	Uri     string
	Version string
	Params  url.Values
}

func get(u *url.URL) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return resp, err
	}

	req.Header.Add("X-Riot-Token", c.Token)
	resp, err = Client.Do(req)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (rr *RiotRequest) AddParameter(key, value string) {
	rr.Params.Add(key, value)
}

func (rr RiotRequest) Get(v interface{}) *RequestError {
	u := &url.URL{
		Scheme:   "https",
		Host:     "na1.api.riotgames.com",
		Path:     fmt.Sprintf("lol/%s/%s/%s", rr.Type, rr.Version, rr.Uri),
		RawQuery: rr.Params.Encode(),
	}

	resp, err := get(u)
	if err != nil {
		return NewErrorFromString(err.Error())
	}

	if err = json.NewDecoder(resp.Body).Decode(v); err != nil {
		return NewErrorFromString(err.Error())
	}

	if isBad(resp.StatusCode) {
		err = NewRequestError(resp)
	}

	return nil
}
