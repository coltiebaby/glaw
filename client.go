package glaw

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/coltiebaby/glaw/ratelimit"
	"github.com/coltiebaby/glaw/errors"
)

type RiotClient struct {
	region           Region
	rateLimitEnabled bool
	limiter          ratelimit.Limiter
}

func NewRiotClient(region Region, enabled bool) *RiotClient {
	var rateLimiter ratelimit.Limiter
	if enabled {
		rateLimiter = SetupRateLimiter(enabled)
	}

	return &RiotClient{
		rateLimitEnabled: enabled,
		region:           region,
		limiter:          rateLimiter,
	}
}

func (rc *RiotClient) NewRequest(uri string) (req ApiRequest) {
	req = &RiotRequest{
		uri: uri,
	}

	return req
}

func (rc *RiotClient) Get(req ApiRequest) (resp *http.Response, err error) {
	platform := RegionsPlatform[rc.region]
	host := fmt.Sprintf("%s.api.riotgames.com", strings.ToLower(platform))

	u := &url.URL{
		Scheme:   "https",
		Host:     host,
		Path:     fmt.Sprintf("lol/%s", req.Uri()),
		RawQuery: req.Encode(),
	}

	if rc.rateLimitEnabled {
		rc.limiter.Take(int(rc.region))
	}

	if resp, err = get(u); err != nil {
		err = errors.NewErrorFromString(err.Error())
	}

	return resp, err

}

func (rc *RiotClient) ChangeRegion(region Region) {
	rc.region = region
}

type RiotRequest struct {
	uri    string
	params url.Values
}

func (rr *RiotRequest) Encode() string {
	return rr.params.Encode()
}

func (rr *RiotRequest) Uri() string {
	return rr.uri
}

func (rr *RiotRequest) AddParameter(key, value string) {
	rr.params.Add(key, value)
}

func (rr *RiotRequest) SetParameters(params url.Values) {
	rr.params = params
}

func GetResultFromResp(resp *http.Response, v interface{}) (err error) {
	if e := json.NewDecoder(resp.Body).Decode(v); e != nil {
		return errors.NewErrorFromString(e.Error())
	}

	if isBad(resp.StatusCode) {
		err = errors.NewRequestError(resp)
	}

	return err
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

	return resp, err
}
