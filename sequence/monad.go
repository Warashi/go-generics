package sequence

import (
	"github.com/Warashi/go-generics/monad"
	"github.com/Warashi/go-generics/types"
	"github.com/Warashi/go-generics/zero"
)

var (
	_ monad.Monad[int, string, Sequence[int], Sequence[string]]         = MonadImpl[int, string]{}
	_ monad.AdditiveMonad[int, string, Sequence[int], Sequence[string]] = MonadImpl[int, string]{}
)

type MonadImpl[T, U any] struct{}

func (MonadImpl[T, U]) Unit(value U) Sequence[U] {
	return Of(value)
}

func (MonadImpl[T, U]) Zero() Sequence[U] {
	return Of[U]()
}

func (MonadImpl[T, U]) Bind(src Sequence[T], f types.Function[T, Sequence[U]]) Sequence[U] {
	return &BindSequence[T, U]{
		base:     src,
		function: f,
	}
}

func (MonadImpl[T, U]) Plus(a, b Sequence[T]) Sequence[T] {
	return &PlusSequence[T]{
		second: false,
		base:   [2]Sequence[T]{a, b},
	}
}

type BindSequence[T, U any] struct {
	base     Sequence[T]
	current  Sequence[U]
	function types.Function[T, Sequence[U]]
}

func (s *BindSequence[T, U]) Next() bool {
	for {
		if s.current != nil && s.current.Next() {
			return true
		}
		if !s.base.Next() {
			return false
		}
		s.current = s.function.Apply(s.base.Value())
	}
}

func (s *BindSequence[T, U]) Value() U {
	if s.current == nil {
		return zero.New[U]()
	}
	return s.current.Value()
}

type PlusSequence[T any] struct {
	second bool
	base   [2]Sequence[T]
}

func (s *PlusSequence[T]) Next() bool {
	if !s.second && s.base[0].Next() {
		return true
	}
	s.second = true
	return s.base[1].Next()
}

func (s *PlusSequence[T]) Value() T {
	if !s.second {
		return s.base[0].Value()
	}
	return s.base[1].Value()
}
