package sequence

import (
	"github.com/Warashi/go-generics/monad"
	"github.com/Warashi/go-generics/types"
	"github.com/Warashi/go-generics/zero"
)

var (
	_ Sequence[any] = (*SliceSequence[any])(nil)
)

type SliceSequence[T any] struct {
	cursor int
	base   []T
}

func (s *SliceSequence[T]) Next() bool {
	if s.cursor+1 >= len(s.base) {
		return false
	}
	s.cursor++
	return true
}

func (s *SliceSequence[T]) Value() T {
	if s.cursor < 0 {
		return zero.New[T]()
	}
	return s.base[s.cursor]
}

func Of[T any](value ...T) Sequence[T] {
	return &SliceSequence[T]{
		cursor: -1,
		base:   value,
	}
}

type Sequence[T any] interface {
	Next() bool
	Value() T
}

func Collect[T any](s Sequence[T]) []T {
	var v []T
	for s.Next() {
		v = append(v, s.Value())
	}
	return v
}

func Flatten[T any](s Sequence[Sequence[T]]) Sequence[T] {
	return FlatMap[Sequence[T], T](s, types.Identity[Sequence[T]]{})
}

func Map[F, T any](from Sequence[F], f types.Function[F, T]) Sequence[T] {
	return monad.Map[Sequence[T]](MonadImpl[F, T]{}, from, f)
}
func FlatMap[F, T any](from Sequence[F], f types.Function[F, Sequence[T]]) Sequence[T] {
	return monad.FlatMap(MonadImpl[F, T]{}, from, f)
}
func Filter[T any](from Sequence[T], f types.Function[T, bool]) Sequence[T] {
	return monad.Filter(MonadImpl[T, T]{}, from, f)
}
func ForEach[T any](from Sequence[T], f types.Consumer[T]) {
	monad.Do[types.Void, Sequence[types.Void]](MonadImpl[T, types.Void]{}, from, f)
}
