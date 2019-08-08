package clock

import (
	"time"
)

func FillRate(limit, seconds int) time.Duration {
	ms := time.Duration(seconds) * time.Second
	return ms / time.Duration(limit)
}

type Clock interface {
	Finish() error
	Tick() <-chan time.Time
	// Stop() <-chan time.Time
}
