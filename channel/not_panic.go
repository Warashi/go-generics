package channel

func NoPanic[T any](base chan T) NonPanic[T] {
	return NonPanic[T]{base: base}
}

type NonPanic[T any] struct {
	base chan T
}

func (ch NonPanic[T]) Send(v T) {
	defer func() { recover() }()
	ch.base <- v
}

func (ch NonPanic[T]) C() <-chan T {
	return ch.base
}

func (ch NonPanic[T]) Close() {
	defer func() { recover() }()
	close(ch.base)
}
