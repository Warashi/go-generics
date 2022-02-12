package slice

import "github.com/Warashi/go-generics/monad"
import "github.com/Warashi/go-generics/types"

func Map[F, T any](from []F, f types.Function[F, T]) []T {
	return monad.Map[[]T](MonadImpl[F, T]{}, from, f)
}
func FlatMap[F, T any](from []F, f types.Function[F, []T]) []T {
	return monad.FlatMap(MonadImpl[F, T]{}, from, f)
}
func Filter[T any](from []T, f types.Function[T, bool]) []T {
	return monad.Filter(MonadImpl[T, T]{}, from, f)
}
func ForEach[T any](from []T, f types.Consumer[T]) {
	monad.Do[types.Void, []types.Void](MonadImpl[T, types.Void]{}, from, f)
}
