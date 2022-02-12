package types

var (
	_ Function[int, bool]           = Closure[int, bool](nil)
	_ BiFunction[int, string, bool] = BiClosure[int, string, bool](nil)
)

type Function[F, T any] interface {
	Apply(F) T
}

type Closure[F, T any] func(F) T

func (f Closure[F, T]) Apply(value F) T { return f(value) }

type BiFunction[F1, F2, T any] interface {
	Apply(F1, F2) T
}

type BiClosure[F1, F2, T any] func(F1, F2) T

func (f BiClosure[F1, F2, T]) Apply(f1 F1, f2 F2) T { return f(f1, f2) }

type Consumer[T any] interface {
	Accept(T)
}

type ConsumerClosure[T any] func(T)

func (f ConsumerClosure[T]) Accept(value T) { f(value) }
