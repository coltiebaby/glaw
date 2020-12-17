// Rate Limit

// The debug token from riot will handle twenty requests in a second or
// it will do 100 requests every two minutes.

package ratelimit

import (
	"context"
	"fmt"
	"time"
)

const ResetDur time.Duration = time.Second * 2

type empty struct{}

type Limiter struct {
	Burst int
	Max   int
	Reset time.Duration

	queue chan empty
	burst chan empty
}

func NewLimiter(burst, max int, dur time.Duration) *Limiter {
	var e empty

	limiter := &Limiter{
		Burst: burst,
		Max:   max,
		Reset: dur,
		queue: make(chan empty, max),
		burst: make(chan empty, burst),
	}

	for i := 0; i < max; i++ {
		select {
		case limiter.queue <- e:
		default:
		}

		select {
		case limiter.burst <- e:
		default:
		}
	}

	return limiter
}

func (limit *Limiter) fillBurst(ctx context.Context) {
	ticker := time.NewTicker(time.Second)

	f := func() {
		select {
		case <-ticker.C:
		case <-ctx.Done():
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

func (limit *Limiter) fillQueue(ctx context.Context) {
	var e empty
	ticker := time.NewTicker(time.Duration(limit.Max) / limit.Reset)

	f := func() {
		for {
			select {
			case <-ticker.C:
			case <-ctx.Done():
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

func (limit *Limiter) Get(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-limit.burst:
		return nil
	}
}

func (limit *Limiter) Fill(ctx context.Context) {
	limit.fillQueue(ctx)
	limit.fillBurst(ctx)
}

func (limit *Limiter) Stop() {
	close(limit.queue)
	close(limit.burst)
}

var EmptyErr error = fmt.Errorf("empty")
