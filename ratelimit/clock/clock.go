// A fancy stopwatch that monitors the rate we fill ratelimit.jar.Jar
package clock

import (
	"time"
)

// Figures out the pace to fill the bucket.
func FillRate(maxLimit, seconds int) time.Duration {
	ms := time.Duration(seconds) * time.Second
	return ms / time.Duration(maxLimit)
}

// Basic timer that helps monitor the rate we fill the bucket
type Clock interface {
	Finish() error
	Tick() <-chan time.Time
	// Stop() <-chan time.Time
}
