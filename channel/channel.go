package channel

import (
	"github.com/Warashi/go-generics/monad"
	"github.com/Warashi/go-generics/types"
)

func Map[F, T any](from <-chan F, f types.Function[F, T]) <-chan T {
	return monad.Map[<-chan T](MonadImpl[F, T]{}, from, f)
}
func FlatMap[F, T any](from <-chan F, f types.Function[F, <-chan T]) <-chan T {
	return monad.FlatMap(MonadImpl[F, T]{}, from, f)
}
func Filter[T any](from <-chan T, f types.Function[T, bool]) <-chan T {
	return monad.Filter(MonadImpl[T, T]{}, from, f)
}
func ForEach[T any](from <-chan T, f types.Consumer[T]) {
	monad.Do[types.Void, <-chan types.Void](MonadImpl[T, types.Void]{}, from, f)
}
