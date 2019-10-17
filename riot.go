package glaw

import (
	"net/http"
	"net/url"

	"github.com/coltiebaby/glaw/ratelimit"
)

// SetupRateLimit creates a new Default Limiter to monitor our requests
// This is not required if you end up making your own thing.
func SetupRateLimiter(enabled bool) ratelimit.Limiter {
	limiter := ratelimit.NewRateLimiter(enabled)

	for r, _ := range Regions {
		rl := ratelimit.NewRateLimit(100, 20, 120)
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
