package riot

import (
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

type ApiClient interface {
	NewRequest(string) ApiRequest
}

type ApiRequest interface {
	Get(interface{}) *errors.RequestError
	AddParameter(string, string)
	SetParameters(url.Values)
}

func NewClient(enabled bool) ApiClient {
	var limiter *ratelimit.RateLimit
	if enabled {
		limiter = ratelimit.Start()
	}

	return &RiotClient{
		rateLimitEnabled: enabled,
		limiter:          limiter,
	}
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
