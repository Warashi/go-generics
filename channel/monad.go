package channel

import "github.com/Warashi/go-generics/monad"
import "github.com/Warashi/go-generics/types"

var _ monad.Monad[int, string, <-chan int, <-chan string] = MonadImpl[int, string]{}

type MonadImpl[T, U any] struct{}

func (MonadImpl[T, U]) Unit(value U) <-chan U {
	ch := make(chan U, 1)
	ch <- value
	close(ch)
	return ch
}

func (MonadImpl[T, U]) Bind(src <-chan T, f types.Applyer[T, <-chan U]) <-chan U {
	result := make(chan U)
	go func() {
		defer close(result)
		for v := range src {
			for vv := range f.Apply(v) {
				result <- vv
			}
		}
	}()
	return result
}
