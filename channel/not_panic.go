package channel

func NoPanic[T any](base chan T) NotPanic[T] {
	return NotPanic[T]{base: base}
}

type NotPanic[T any] struct {
	base chan T
}

func (ch NotPanic[T]) Send(v T) {
	defer func() { recover() }()
	ch.base <- v
}

func (ch NotPanic[T]) Receive() T {
	return <-ch.base
}

func (ch NotPanic[T]) Close() {
	defer func() { recover() }()
	close(ch.base)
}
