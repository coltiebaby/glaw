package ratelimit

import (
	"time"
)

type RateLimit struct {
	On         bool
	reset      int
	limit      int
	burstLimit int
	gauge      chan time.Time
}

func NewLimit(limit, burstLimit, reset int) *RateLimit {
	gauge := make(chan time.Time, burstLimit)
	return &RateLimit{
		On:         true,
		reset:      reset,
		limit:      limit,
		burstLimit: burstLimit,
		gauge:      gauge,
	}
}

func (rl *RateLimit) Increase(t time.Time) {
	if rl.Len() < rl.Max() {
		rl.gauge <- t
	}
}

func (rl *RateLimit) Len() int {
	return len(rl.gauge)
}

func (rl *RateLimit) Max() int {
	return rl.burstLimit
}

func (rl *RateLimit) Stop() {
	close(rl.gauge)
}

func (rl *RateLimit) Rate() time.Duration {
	ms := time.Duration(rl.reset) * time.Second // 10s
	return ms / time.Duration(rl.limit)
}

func (rl *RateLimit) Request() bool {
	if !rl.On {
		return true
	}

	<-rl.gauge
	return true
}

func Start() *RateLimit {
	rl := NewLimit(100, 20, 120)

	// Fills the channel at a fixed rate to replace all requests made
	go func() {
		ticker := time.NewTicker(rl.Rate())
		defer ticker.Stop()

		for {
			rl.Increase(<-ticker.C)
		}
	}()

	return rl
}
