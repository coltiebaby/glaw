// Rate Limit

// The debug token from riot will handle twenty requests in a second or
// it will do 100 requests every two minutes.

package ratelimit

import (
	"context"
	"fmt"
	"time"
)

// Figures out the pace to fill the bucket.
func rate(limit int, dur time.Duration) time.Duration {
	return dur / time.Duration(limit)
}

type empty struct{}

type Limiter struct {
	Burst int
	Max   int
	Reset time.Duration

	stop  chan empty
	queue chan empty
	burst chan empty
}

func NewLimiter(burst, max int, dur time.Duration) *Limiter {
	limiter := &Limiter{
		Burst: burst,
		Max:   max,
		Reset: dur,
		queue: make(chan empty, max),
		burst: make(chan empty, burst),
	}

	return limiter
}

func (limiter *Limiter) fill() {
	var e empty
	for i := 0; i < limiter.Max; i++ {
		select {
		case limiter.queue <- e:
		default:
		}

		select {
		case limiter.burst <- e:
		default:
		}
	}
}

// Fill burst fills the burst bucket every second. If you drain this you
// have to wait for the queue bucket to refill.
func (limit *Limiter) fillBurst(ctx context.Context) {
	ticker := time.NewTicker(time.Second)

	f := func() {
		select {
		case <-ticker.C:
		case <-ctx.Done():
			return
		case <-limit.stop:
			return
		}

		for q := range limit.queue {
			select {
			case limit.burst <- q:
			default:
				goto exit
			}

		exit:
			break
		}
	}

	go f()
}

// Fill queue is the total amount of reqeusts you can make per X seconds
func (limit *Limiter) fillQueue(ctx context.Context) {
	ticker := time.NewTicker(rate(limit.Max, limit.Reset))

	var e empty
	f := func() {
		for {
			select {
			case <-ticker.C:
			case <-ctx.Done():
				return
			case <-limit.stop:
				return
			}

			select {
			case limit.queue <- e:
			default:
			}
		}
	}

	go f()
}

// Must Get will return an error if it cannot make a request or nil if it can
func (limit *Limiter) MustGet(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-limit.burst:
		return nil
	default:
	}

	return EmptyErr
}

// Get will wait until it can get a request or until the context is canceled
func (limit *Limiter) Get(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-limit.burst:
		return nil
	}
}

// Starts the fill routine functions
func (limit *Limiter) Start(ctx context.Context) {
	limit.fill()

	limit.stop = make(chan empty)

	limit.fillQueue(ctx)
	limit.fillBurst(ctx)
}

// Stops and cleans up the limiter
func (limit *Limiter) Stop() {
	close(limit.stop)
	close(limit.queue)
	close(limit.burst)
}

var EmptyErr error = fmt.Errorf("empty")
