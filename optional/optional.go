package optional

import "github.com/Warashi/go-generics/zero"

type Optional[T any] struct {
	value *T
}

func (o Optional[T]) IsEmpty() bool {
	return o.value == nil
}

func (o Optional[T]) OrElse(e T) T {
	if o.IsEmpty() {
		return e
	}
	return *o.value
}

func (o Optional[T]) OrElseZero() T {
	if o.IsEmpty() {
		return zero.New[T]()
	}
	return *o.value
}

func New[T any](value T) Optional[T] {
	return Optional[T]{value: &value}
}

func Empty[T any]() Optional[T] {
	return Optional[T]{}
}

type Mapper[F, T any] interface {
	Apply(F) T
}

type MapperFunc[F, T any] func(F) T

func (f MapperFunc[F, T]) Apply(value F) T { return f(value) }
func Map[F, T any](o Optional[F], mapper Mapper[F, T]) Optional[T] {
	if o.IsEmpty() {
		return Empty[T]()
	}
	return New(mapper.Apply(*o.value))
}

type FlatMapper[F, T any, O Optional[T]] Mapper[F, O]
type FlatMapperFunc[F, T any] func(F) Optional[T]

func (f FlatMapperFunc[F, T]) Apply(value F) Optional[T] { return f(value) }
func FlatMap[F, T any](o Optional[F], mapper FlatMapper[F, T]) Optional[T] {
	if o.IsEmpty() {
		return Empty[T]()
	}
	return mapper.Apply(*o.value)
}
