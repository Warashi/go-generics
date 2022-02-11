package optional

import "github.com/Warashi/go-generics/types"

type MonadImpl[T, U any] struct{}

func (MonadImpl[T, U]) Unit(value U) Optional[U] {
	return New(value)
}

func (MonadImpl[T, U]) Bind(src Optional[T], f types.Applyer[T, Optional[U]]) Optional[U] {
	if src.IsEmpty() {
		return Empty[U]()
	}
	return f.Apply(src.OrElseZero())
}
