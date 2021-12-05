package sequence

import (
	"github.com/Warashi/go-generics/zero"
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

type Applyer[T, R any] interface {
	Apply(T) R
}

type ApplyerFunc[T, R any] func(T) R

func (f ApplyerFunc[T, R]) Apply(v T) R {
	return f(v)
}

type Consumer[T any] interface {
	Consume(T)
}

type ConsumerFunc[T any] func(T)

func (f ConsumerFunc[T]) Consume(v T) {
	f(v)
}

func ForEach[T any](s Sequence[T], c Consumer[T]) {
	for s.Next() {
		c.Consume(s.Value())
	}
}

func Consume[T any](s Sequence[T]) []T {
	var v []T
	for s.Next() {
		v = append(v, s.Value())
	}
	return v
}

type MapSequence[T, R any] struct {
	base    Sequence[T]
	applyer Applyer[T, R]
}

func (s *MapSequence[T, R]) Next() bool {
	return s.base.Next()
}

func (s *MapSequence[T, R]) Value() R {
	return s.applyer.Apply(s.base.Value())
}

func Map[T, R any](s Sequence[T], a Applyer[T, R]) Sequence[R] {
	return &MapSequence[T, R]{
		base:    s,
		applyer: a,
	}
}

type FlatMapSequence[T, R any] struct {
	base    Sequence[T]
	current Sequence[R]
	applyer Applyer[T, Sequence[R]]
}

func (s *FlatMapSequence[T, R]) Next() bool {
	for {
		if s.current != nil && s.current.Next() {
			return true
		}
		if !s.base.Next() {
			return false
		}
		s.current = s.applyer.Apply(s.base.Value())
	}
}

func (s *FlatMapSequence[T, R]) Value() R {
	if s.current == nil {
		return zero.New[R]()
	}
	return s.current.Value()
}

func FlatMap[T, R any](s Sequence[T], a Applyer[T, Sequence[R]]) Sequence[R] {
	return &FlatMapSequence[T, R]{
		base:    s,
		applyer: a,
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

func Flatten[T any](s Sequence[Sequence[T]]) *FlattenSequence[T] {
	return &FlattenSequence[T]{
		base: s,
	}
}
