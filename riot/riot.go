package riot

import (
	"net/http"
	"net/url"

	"github.com/coltiebaby/g-law/config"
	"github.com/coltiebaby/g-law/ratelimit"
	"github.com/coltiebaby/g-law/ratelimit/clock"
	"github.com/coltiebaby/g-law/ratelimit/jar"
)

var (
	c = config.FromEnv()
	// Client = NewClient(REGION_NA, c.EnableRateLimiting)
	Client ApiClient = NewRiotClient(REGION_NA, c.EnableRateLimiting)
)

// SetupRateLimit creates a new Default Limiter to monitor our requests
// This is not required if you end up making your own thing.
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
	Get(ApiRequest) (*http.Response, error)
}

type ApiRequest interface {
	AddParameter(string, string)
	SetParameters(url.Values)
	Encode() string
	Uri() string
}
