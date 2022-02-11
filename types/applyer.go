package types

var _ Applyer[int, bool] = ApplyerFunc[int, bool](nil)

type Applyer[F, T any] interface {
	Apply(F) T
}

type ApplyerFunc[F, T any] func(F) T

func (f ApplyerFunc[F, T]) Apply(value F) T { return f(value) }
