package optional

import "github.com/Warashi/go-generics/zero"
import "github.com/Warashi/go-generics/types"
import "github.com/Warashi/go-generics/monad"
import "github.com/google/go-cmp/cmp"
import "github.com/google/go-cmp/cmp/cmpopts"

type Optional[T any] struct {
	value *T
}

func (o Optional[T]) Equal(o2 Optional[T]) bool {
	if o.IsEmpty() && o2.IsEmpty() {
		return true
	}
	if !o.IsEmpty() && !o2.IsEmpty() {
		return cmp.Equal(o.OrElseZero(), o2.OrElseZero(), cmpopts.IgnoreUnexported())
	}
	return false
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

func Map[F, T any](o Optional[F], f types.Applyer[F, T]) Optional[T] {
	return monad.Map[Optional[T]](MonadImpl[F, T]{}, o, f)
}

func FlatMap[F, T any](o Optional[F], f types.Applyer[F, Optional[T]]) Optional[T] {
	return monad.FlatMap(MonadImpl[F, T]{}, o, f)
}
