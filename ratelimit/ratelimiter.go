package ratelimit

import (
    "context"
)

func NewRateLimiter(enabled bool) *RateLimiter {
	return &RateLimiter{
		enabled: enabled,
		timers:  make(map[int]Limiter),
	}
}

type RateLimiter struct {
	enabled bool
	timers  map[int]Limiter
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

func (r *RateLimiter) MustGet(ctx context.Context, region int) error {
	return r.timers[region].MustGet()
}

func (r *RateLimiter) Get(ctx context.Context, region int) error {
	return r.timers[region].Get()
}
