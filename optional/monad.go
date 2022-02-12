package optional

import "github.com/Warashi/go-generics/monad"
import "github.com/Warashi/go-generics/types"

var _ monad.Monad[int, string, Optional[int], Optional[string]] = MonadImpl[int, string]{}
var _ monad.AdditiveMonad[int, string, Optional[int], Optional[string]] = MonadImpl[int, string]{}

type MonadImpl[T, U any] struct{}

func (MonadImpl[T, U]) Unit(value U) Optional[U] {
	return New(value)
}

func (MonadImpl[T, U]) Bind(src Optional[T], f types.Function[T, Optional[U]]) Optional[U] {
	if src.IsEmpty() {
		return Empty[U]()
	}
	return f.Apply(src.OrElseZero())
}

func (MonadImpl[T, U]) Zero() Optional[U] {
	return Empty[U]()
}

func (MonadImpl[T, U]) Plus(a Optional[T], b Optional[T]) Optional[T] {
	if a.IsEmpty() {
		return b
	}
	return a
}
