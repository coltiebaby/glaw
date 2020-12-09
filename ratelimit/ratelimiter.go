package ratelimit

import (
	"context"
)

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		timers: make(map[int]*Limiter),
	}
}

type RateLimiter struct {
	timers map[int]*Limiter
}

func (r *RateLimiter) Add(region int, l *Limiter) {
	r.timers[region] = l
}

func (r *RateLimiter) timer(region int) *Limiter {
	if limiter, ok := r.timers[region]; ok {
		return limiter
	}

	return NewLimiter(20, 100)
}

func (r *RateLimiter) MustGet(ctx context.Context, region int) error {
	return r.timer(region).MustGet(ctx)
}

func (r *RateLimiter) Get(ctx context.Context, region int) error {
	return r.timer(region).Get(ctx)
}
