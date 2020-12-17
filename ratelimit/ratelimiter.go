package ratelimit

import (
	"context"
	"time"
)

func NewRateLimiter(burst, max int, dur time.Duration) *RateLimiter {
	ctx, cancel := context.WithCancel(context.Background())

	return &RateLimiter{
		ctx:    ctx,
		cancel: cancel,
		Burst:  burst,
		Max:    max,
		Reset:  dur,
		timers: make(map[int]*Limiter),
	}
}

type RateLimiter struct {
	ctx    context.Context
	cancel func()

	Burst int
	Max   int
	Reset time.Duration

	timers map[int]*Limiter
}

func (r *RateLimiter) Stop() {
	r.cancel()
}

func (r *RateLimiter) timer(region int) *Limiter {
	if limiter, ok := r.timers[region]; ok {
		return limiter
	}

	limiter := NewLimiter(r.Burst, r.Max, r.Reset)
	limiter.Start(r.ctx)
	r.timers[region] = limiter

	return limiter
}

func (r *RateLimiter) MustGet(ctx context.Context, region int) error {
	return r.timer(region).MustGet(ctx)
}

func (r *RateLimiter) Get(ctx context.Context, region int) error {
	return r.timer(region).Get(ctx)
}
