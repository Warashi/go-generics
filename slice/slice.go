package slice

import "github.com/Warashi/go-generics/monad"
import "github.com/Warashi/go-generics/types"

func Map[F, T any](from []F, f types.Applyer[F, T]) []T {
	return monad.Map[[]T](MonadImpl[F, T]{}, from, f)
}
func FlatMap[F, T any](from []F, f types.Applyer[F, []T]) []T {
	return monad.FlatMap(MonadImpl[F, T]{}, from, f)
}
