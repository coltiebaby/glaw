package ratelimit

import (
	"testing"
	"time"
)

func parse(t string) time.Duration {
	d, _ := time.ParseDuration(t)
	return d
}

func TestLimitRate(t *testing.T) {
	type Tester struct {
		Limit    int
		Duration int
		Result   time.Duration
	}

	tests := []Tester{
		Tester{Limit: 10, Duration: 1, Result: parse(`100ms`)},
		Tester{Limit: 100, Duration: 120, Result: parse(`1200ms`)},
		Tester{Limit: 5, Duration: 10, Result: parse(`2000ms`)},
	}

	for _, test := range tests {
		limitAndBurst := test.Limit
		limiter := NewLimit(limitAndBurst, limitAndBurst, test.Duration)

		a := limiter.Rate()
		b := test.Result

		if a != b {
			t.Errorf("Results did not match: %s != %s", a, b)
		}
	}
}
