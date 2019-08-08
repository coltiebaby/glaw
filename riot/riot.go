package riot

import (
	"net/http"
	"net/url"

	"github.com/coltiebaby/g-law/config"
	"github.com/coltiebaby/g-law/ratelimit"
	"github.com/coltiebaby/g-law/ratelimit/clock"
	"github.com/coltiebaby/g-law/ratelimit/jar"
	"github.com/coltiebaby/g-law/riot/errors"
)

var (
	c = config.FromEnv()
	// Client = NewClient(REGION_NA, c.EnableRateLimiting)
	Client = NewClient(REGION_NA, c.EnableRateLimiting)
)

func SetupRateLimiter(enabled bool) ratelimit.Limiter {
	limiter := ratelimit.NewRateLimiter(enabled)

	for r, _ := range Regions {
		c := clock.NewClock(100, 120)
		j := jar.NewBucket(20)

		rl := ratelimit.NewRateLimit(c, j)

		limiter.Add(int(r), rl)
	}

	return limiter
}

type ApiClient interface {
	NewRequest(string) ApiRequest
	ChangeRegion(Region)
}

type ApiRequest interface {
	Get(interface{}) *errors.RequestError
	AddParameter(string, string)
	SetParameters(url.Values)
}

func NewClient(region Region, enabled bool) ApiClient {
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
