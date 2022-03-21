package hedging

import (
	"context"
	"time"
)

type result[T any] struct {
	value T
	err   error
}

func Do[T any](parentCtx context.Context, delay time.Duration, f func(context.Context) (T, error)) (T, error) {
	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	ch1, ch2 := make(chan result[T], 1), make(chan result[T], 1)

	go func() {
		defer close(ch1)
		var r result[T]
		r.value, r.err = f(parentCtx)
		ch1 <- r
	}()
	go func() {
		defer close(ch2)
		select {
		case <-ctx.Done():
			return
		case <-time.After(delay):
		}
		var r result[T]
		r.value, r.err = f(parentCtx)
		ch2 <- r
	}()

	select {
	case v := <-ch1:
		return v.value, v.err
	case v := <-ch2:
		return v.value, v.err
	}
}
