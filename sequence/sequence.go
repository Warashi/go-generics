package sequence

import (
	"github.com/Warashi/go-generics/zero"
	"github.com/Warashi/go-generics/types"
)

var (
	_ Sequence[any] = (*SliceSequence[any])(nil)
	_ Sequence[any] = (*MapSequence[any, any])(nil)
	_ Sequence[any] = (*FlatMapSequence[any, any])(nil)
	_ Sequence[any] = (*FlattenSequence[any])(nil)
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

func ForEach[T any](s Sequence[T], c types.Consumer[T]) {
	for s.Next() {
		c.Accept(s.Value())
	}
}

func Collect[T any](s Sequence[T]) []T {
	var v []T
	for s.Next() {
		v = append(v, s.Value())
	}
	return v
}

type MapSequence[T, R any] struct {
	base    Sequence[T]
	function types.Function[T, R]
}

func (s *MapSequence[T, R]) Next() bool {
	return s.base.Next()
}

func (s *MapSequence[T, R]) Value() R {
	return s.function.Apply(s.base.Value())
}

func Map[T, R any](s Sequence[T], f types.Function[T, R]) Sequence[R] {
	return &MapSequence[T, R]{
		base:    s,
		function: f,
	}
}

type FlatMapSequence[T, R any] struct {
	base    Sequence[T]
	current Sequence[R]
	function types.Function[T, Sequence[R]]
}

func (s *FlatMapSequence[T, R]) Next() bool {
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

func (s *FlatMapSequence[T, R]) Value() R {
	if s.current == nil {
		return zero.New[R]()
	}
	return s.current.Value()
}

func FlatMap[T, R any](s Sequence[T], f types.Function[T, Sequence[R]]) Sequence[R] {
	return &FlatMapSequence[T, R]{
		base:    s,
		function: f,
	}
}

type FlattenSequence[T any] struct {
	base Sequence[Sequence[T]]
}

func (s *FlattenSequence[T]) Next() bool {
	for s.base.Value() == nil || !s.base.Value().Next() {
		if ok := s.base.Next(); !ok {
			return false
		}
	}
	return true
}

func (s *FlattenSequence[T]) Value() T {
	return s.base.Value().Value()
}

func Flatten[T any](s Sequence[Sequence[T]]) Sequence[T] {
	return &FlattenSequence[T]{
		base: s,
	}
}

type FilterSequence[T any] struct {
	base   Sequence[T]
	filter types.Function[T, bool]
}

func (s *FilterSequence[T]) Next() bool {
	for s.base.Next() {
		if s.filter.Apply(s.base.Value()) {
			return true
		}
	}
	return false
}

func (s *FilterSequence[T]) Value() T {
	return s.base.Value()
}

func Filter[T any](s Sequence[T], filter types.Function[T, bool]) Sequence[T] {
	return &FilterSequence[T]{
		base:   s,
		filter: filter,
	}
}
