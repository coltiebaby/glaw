package ratelimit

import (
	"github.com/coltiebaby/g-law/ratelimit/clock"
	"github.com/coltiebaby/g-law/ratelimit/jar"
)

type RateLimit struct {
	burst int

	jar jar.Jar
}

func NewRateLimit(c clock.Clock, j jar.Jar) (rl *RateLimit) {
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
