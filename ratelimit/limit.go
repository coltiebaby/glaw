package ratelimit

import (
	"github.com/coltiebaby/glaw/ratelimit/clock"
	"github.com/coltiebaby/glaw/ratelimit/jar"
)

type RateLimit struct {
	burst int

	jar jar.Jar
}

func NewRateLimit(max, burst, seconds int) (rl *RateLimit) {
	c := clock.NewClock(max, seconds)
	j := jar.NewBucket(burst)
	rl = newRateLimit(c, j)

	return rl
}

func newRateLimit(c clock.Clock, j jar.Jar) (rl *RateLimit) {
	rl = &RateLimit{
		jar: j,
	}

	go func() {
		defer c.Finish()
		for {
			select {
			// case <-c.Stop():
			//     break
			case <-c.Tick():
				rl.Recharge()
			}
		}
	}()

	return rl
}

func (r *RateLimit) Recharge() {
	if !r.jar.Full() {
		r.jar.Give()
	}
}

func (r *RateLimit) Take() {
	r.jar.Take()
}
