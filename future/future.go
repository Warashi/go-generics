package future

import "context"
import "github.com/Warashi/go-generics/channel"
import "github.com/Warashi/go-generics/zero"

type result[T any] struct {
	value T
	error error
}

type Result[T any] struct {
	ch channel.NonPanic[result[T]]
}

func (r Result[T]) Get(ctx context.Context) (T, error) {
	select {
	case <-ctx.Done():
		return zero.New[T](), ctx.Err()
	case got := <-r.ch.C():
		return got.value, got.error
	}
}

type Task[T any] interface {
	Do(context.Context) (T, error)
}

type TaskFunc[T any] func(ctx context.Context) (T, error)

func (f TaskFunc[T]) Do(ctx context.Context) (T, error) { return f(ctx) }

func Do[T any](ctx context.Context, t Task[T]) Result[T] {
	ch := channel.NoPanic(make(chan result[T], 1))
	go func() {
		defer ch.Close()
		var r result[T]
		r.value, r.error = t.Do(ctx)
		ch.Send(r)
	}()
	return Result[T]{ch: ch}
}
