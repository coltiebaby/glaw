package clock

import (
	"time"
)

type Stopwatch struct {
	ticker *time.Ticker
	// timer  *time.Timer
}

func (s *Stopwatch) Tick() <-chan time.Time {
	return s.ticker.C
}

// func (s *Stopwatch) Stop() <-chan time.Time {
//     return s.timer.C
// }

func (s *Stopwatch) Finish() error {
	s.ticker.Stop()

	return nil
}

func NewClock(limit, seconds int) Clock {
	ticker := time.NewTicker(FillRate(limit, seconds))

	return &Stopwatch{
		ticker: ticker,
	}
}
