package ratelimit

func NewRateLimiter(enabled bool) *RateLimiter {
	return &RateLimiter{
		enabled: enabled,
		timers:  make(map[int]Limit),
	}
}

type RateLimiter struct {
	enabled bool
	timers  map[int]Limit
}

func (r *RateLimiter) Activate() {
	r.enabled = true
}

func (r *RateLimiter) Deactivate() {
	r.enabled = false
}

func (r *RateLimiter) Add(region int, l Limit) {
	r.timers[region] = l
}

func (r *RateLimiter) Take(region int) {
	r.timers[region].Take()
}
