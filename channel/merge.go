package channel

import (
	"context"
)

func Merge[T any](ctx context.Context, ch ...<-chan T) <-chan T {
	switch len(ch) {
	case 0:
		return OrDone[T](ctx, nil)
	case 1:
		return OrDone(ctx, ch[0])
	case 2:
		break
	default:
		left, right := ch[:len(ch)/2], ch[len(ch)/2:]
		return Merge(ctx, Merge(ctx, left...), Merge(ctx, right...))
	}
	ret := make(chan T)
	go func() {
		defer close(ret)

		left, right := ch[0], ch[1]
	loop:
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-left:
				if !ok {
					left = nil
					continue loop
				}
				select {
				case <-ctx.Done():
					return
				case ret <- v:
				}
			case v, ok := <-right:
				if !ok {
					right = nil
					continue loop
				}
				select {
				case <-ctx.Done():
					return
				case ret <- v:
				}
			}
			if left == nil && right == nil {
				return
			}
		}
	}()
	return ret
}

func OrDone[T any](ctx context.Context, ch <-chan T) <-chan T {
	rch := make(chan T)
	go func() {
		defer close(rch)
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-ch:
				if !ok {
					return
				}
				select {
				case rch <- v:
				case <-ctx.Done():
				}
			}
		}
	}()
	return rch
}
