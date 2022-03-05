package hedging

import (
	"context"
	"time"
)

func Hedge[T any](delay time.Duration, f func() (T, error)) (T, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	type result struct {
		value T
		err   error
	}

	ch1, ch2 := make(chan result, 1), make(chan result, 1)

	go func() {
		defer close(ch1)
		var r result
		r.value, r.err = f()
		ch1 <- r
	}()
	go func() {
		defer close(ch2)
		select {
		case <-ctx.Done():
			return
		case <-time.After(delay):
		}
		var r result
		r.value, r.err = f()
		ch2 <- r
	}()

	select {
	case v := <-ch1:
		return v.value, v.err
	case v := <-ch2:
		return v.value, v.err
	}
}
