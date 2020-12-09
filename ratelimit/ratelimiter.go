package ratelimit

import (
	"context"
	"time"
)

func NewRateLimiter(burst, max int, dur time.Duration) *RateLimiter {
	return &RateLimiter{
		Burst:  burst,
		Max:    max,
		Reset:  dur,
		timers: make(map[int]*Limiter),
	}
}

type RateLimiter struct {
	Burst int
	Max   int
	Reset time.Duration

	timers map[int]*Limiter
}

func (r *RateLimiter) Add(region int, l *Limiter) {
	r.timers[region] = l
}

func (r *RateLimiter) timer(region int) *Limiter {
	if limiter, ok := r.timers[region]; ok {
		return limiter
	}

	return NewLimiter(r.Burst, r.Max, r.Reset)
}

func (r *RateLimiter) MustGet(ctx context.Context, region int) error {
	return r.timer(region).MustGet(ctx)
}

func (r *RateLimiter) Get(ctx context.Context, region int) error {
	return r.timer(region).Get(ctx)
}
