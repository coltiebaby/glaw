package ratelimit

import (
	"context"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*250)
	defer cancel()

	limit := NewLimiter(10, 20, time.Minute)
	limit.fill()

	var count int
	var err error

	for count < 20 {
		err = limit.Get(ctx)
		if err != nil {
			break
		}

		count++
	}

	if err != context.DeadlineExceeded {
		t.Fatalf("expected an empty err but got %s", err)
	}

	if count != 10 {
		t.Fatalf("expected to only get 10; got %d instead", count)
	}
}

func TestMustGet(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*250)
	defer cancel()

	limit := NewLimiter(10, 20, time.Minute)
	limit.fill()

	var count int
	var err error

	for count < 10 {
		err = limit.MustGet(ctx)
		count++
	}

	if err != nil {
		t.Fatalf("expected no err but got %s", err)
	}

	err = limit.MustGet(ctx)

	if err != EmptyErr {
		t.Fatalf("expected an empty err but got %s", err)
	}

	if count != 10 {
		t.Fatalf("expected to only get 10; got %d instead", count)
	}
}
