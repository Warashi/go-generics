package slice

import "github.com/Warashi/go-generics/monad"
import "github.com/Warashi/go-generics/types"

var (
	_ monad.Monad[int, string, []int, []string]         = MonadImpl[int, string]{}
	_ monad.AdditiveMonad[int, string, []int, []string] = MonadImpl[int, string]{}
)

type MonadImpl[T, U any] struct{}

func (MonadImpl[T, U]) Unit(value U) []U {
	return []U{value}
}

func (MonadImpl[T, U]) Bind(src []T, f types.Function[T, []U]) []U {
	var result []U
	for _, v := range src {
		result = append(result, f.Apply(v)...)
	}
	return result
}

func (MonadImpl[T, U]) Zero() []U {
	return nil
}

func (MonadImpl[T, U]) Plus(a, b []T) []T {
	r := make([]T, len(a)+len(b))
	copy(r[:len(a)], a)
	copy(r[len(a):], b)
	return r
}
