package riot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/coltiebaby/g-law/config"
	"github.com/coltiebaby/g-law/ratelimit"
	"github.com/coltiebaby/g-law/riot/errors"
)

var (
	c      = config.FromEnv()
	Client = NewClient(c.EnableRateLimiting)
)

type ApiRequest interface {
	Get(v interface{}) *errors.RequestError
	AddParameter(key, value string)
	SetParameters(url.Values)
}

type RiotClient struct {
	rateLimitEnabled bool
	limiter          *ratelimit.RateLimit
}

func NewClient(enabled bool) *RiotClient {
	var limiter *ratelimit.RateLimit
	if enabled {
		limiter = ratelimit.Start()
	}

	return &RiotClient{
		rateLimitEnabled: enabled,
		limiter:          limiter,
	}
}

func (rc *RiotClient) NewRequest(uri string) (req ApiRequest) {
	req = &RiotRequest{
		uri: uri,
	}

	if rc.rateLimitEnabled {
		rc.limiter.Request()
	}

	return req
}

type RiotRequest struct {
	uri    string
	params url.Values
}

func isBad(code int) bool {
	return (code >= 200 && code < 300) != true
}

func get(u *url.URL) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return resp, err
	}

	req.Header.Add("X-Riot-Token", c.Token)
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func (rr *RiotRequest) AddParameter(key, value string) {
	rr.params.Add(key, value)
}

func (rr *RiotRequest) SetParameters(params url.Values) {
	rr.params = params
}

func (rr RiotRequest) Get(v interface{}) *errors.RequestError {
	u := &url.URL{
		Scheme:   "https",
		Host:     "na1.api.riotgames.com",
		Path:     fmt.Sprintf("lol/%s", rr.uri),
		RawQuery: rr.params.Encode(),
	}

	resp, err := get(u)
	if err != nil {
		return errors.NewErrorFromString(err.Error())
	}

	if err = json.NewDecoder(resp.Body).Decode(v); err != nil {
		return errors.NewErrorFromString(err.Error())
	}

	if isBad(resp.StatusCode) {
		err = errors.NewRequestError(resp)
	}

	return nil
}
