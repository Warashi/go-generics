package types

var (
	_ Function[int, bool]           = Closure[int, bool](nil)
	_ BiFunction[int, string, bool] = BiClosure[int, string, bool](nil)
	_ Consumer[int]                 = ConsumerClosure[int](nil)

	_ Function[int, int] = Identity[int]{}
)

type Identity[T any] struct{}

func (Identity[T]) Apply(v T) T { return v }

type Function[F, T any] interface {
	Apply(F) T
}

func NewFunction[F, T any](f func(F) T) Function[F, T] {
	return Closure[F, T](f)
}

type Closure[F, T any] func(F) T

func (f Closure[F, T]) Apply(value F) T { return f(value) }

func NewBiFunction[F1, F2, T any](f func(F1, F2) T) BiFunction[F1, F2, T] {
	return BiClosure[F1, F2, T](f)
}

type BiFunction[F1, F2, T any] interface {
	Apply(F1, F2) T
}

type BiClosure[F1, F2, T any] func(F1, F2) T

func (f BiClosure[F1, F2, T]) Apply(f1 F1, f2 F2) T { return f(f1, f2) }

func NewConsumer[T any](f func(T)) Consumer[T] {
	return ConsumerClosure[T](f)
}

type Consumer[T any] interface {
	Accept(T)
}

type ConsumerClosure[T any] func(T)

func (f ConsumerClosure[T]) Accept(value T) { f(value) }
