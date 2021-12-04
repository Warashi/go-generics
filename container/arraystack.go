package container

import (
	"errors"

	"github.com/Warashi/go-generics/zero"
)

var ErrIndexOutOfRange = errors.New("index out of range")

type ArrayStack[T any] struct {
	length int
	array  []T
}

func (s *ArrayStack[T]) Size() int {
	return s.length
}

func (s *ArrayStack[T]) resize() {
	n := make([]T, 2*s.length)
	if len(s.array) > 0 {
		copy(n[:s.length], s.array[:s.length])
	}
	s.array = n
}

func (s *ArrayStack[T]) Len() int {
	return s.length
}

func (s *ArrayStack[T]) Get(i int) (T, error) {
	if i < 0 || s.length <= i {
		return zero.New[T](), ErrIndexOutOfRange
	}
	return s.array[i], nil
}

func (s *ArrayStack[T]) Set(i int, x T) error {
	if i < 0 || s.length <= i {
		return ErrIndexOutOfRange
	}
	s.array[i] = x
	return nil
}

func (s *ArrayStack[T]) Add(i int, x T) error {
	if i < 0 || s.length < i {
		return ErrIndexOutOfRange
	}
	s.length++
	if len(s.array) <= s.length+1 {
		s.resize()
	}
	for j := s.length; i < j; j-- {
		s.array[j] = s.array[j-1]
	}
	s.array[i] = x
	return nil
}

func (s *ArrayStack[T]) Remove(i int) (T, error) {
	if i < 0 || s.length <= i {
		return zero.New[T](), ErrIndexOutOfRange
	}
	ret := s.array[i]
	for j := i; j < s.length-1; j++ {
		s.array[j] = s.array[j+1]
	}
	s.length--
	if 3*s.length < len(s.array) {
		s.resize()
	}
	return ret, nil
}

func (s *ArrayStack[T]) Push(x T) {
	// never occur ErrIndexOutOfRange, so we can ignore error
	_ = s.Add(s.length, x)
}

func (s *ArrayStack[T]) Pop() (T, error) {
	return s.Remove(s.length - 1)
}
