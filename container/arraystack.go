package container

import (
	"errors"

	"github.com/Warashi/go-algorithms/zero"
)

var ErrIndexOutOfRange = errors.New("index out of range")

type ArrayStack[T any] struct {
	length int
	array  []T
}

func (s *ArrayStack[T]) resize() {
	n := make([]T, 2*s.length)
	copy(n[:s.length], s.array[:s.length])
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

func (s *ArrayStack[T]) Add(i int, x T) {
	if len(s.array) <= s.length+1 {
		s.resize()
	}
	for j := s.length; i < j; j-- {
		s.array[j] = s.array[j-1]
	}
	s.array[i] = x
	s.length++
}

func (s *ArrayStack[T]) Remove(i int) T {
	ret := s.array[i]
	for j := i; j < s.length-1; j++ {
		s.array[j] = s.array[j+1]
	}
	s.length--
	if 3*s.length < len(s.array) {
		s.resize()
	}
	return ret
}
