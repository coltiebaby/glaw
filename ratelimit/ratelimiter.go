package ratelimit

import (
	"context"
)

func NewRateLimiter(burst, max int) *RateLimiter {
	return &RateLimiter{
        Burst: burst,
        Max: max,
		timers: make(map[int]*Limiter),
	}
}

type RateLimiter struct {
    Burst int
    Max int
	timers map[int]*Limiter
}

func (r *RateLimiter) Add(region int, l *Limiter) {
	r.timers[region] = l
}

func (r *RateLimiter) timer(region int) *Limiter {
	if limiter, ok := r.timers[region]; ok {
		return limiter
	}

	return NewLimiter(r.Burst, r.Max)
}

func (r *RateLimiter) MustGet(ctx context.Context, region int) error {
	return r.timer(region).MustGet(ctx)
}

func (r *RateLimiter) Get(ctx context.Context, region int) error {
	return r.timer(region).Get(ctx)
}
