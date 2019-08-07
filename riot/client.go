package riot

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/coltiebaby/g-law/ratelimit"
	"github.com/coltiebaby/g-law/riot/errors"
)

type RiotClient struct {
	region           Region
	rateLimitEnabled bool
	limiter          *ratelimit.RateLimit
}

func (rc *RiotClient) NewRequest(uri string) (req ApiRequest) {
	req = &RiotRequest{
		region: rc.region,
		uri:    uri,
	}

	if rc.rateLimitEnabled {
		rc.limiter.Request()
	}

	return req
}

type RiotRequest struct {
	region Region
	uri    string
	params url.Values
}

func (rr *RiotRequest) AddParameter(key, value string) {
	rr.params.Add(key, value)
}

func (rr *RiotRequest) SetParameters(params url.Values) {
	rr.params = params
}

func (rr RiotRequest) Get(v interface{}) *errors.RequestError {
	platform := RegionsPlatform[rr.region]
	host := fmt.Sprintf("%s.api.riotgames.com", strings.ToLower(platform))

	u := &url.URL{
		Scheme:   "https",
		Host:     host,
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
