package container

import (
	"fmt"
	"math"

	"github.com/Warashi/go-generics/zero"
)

type RootishArrayStack[T any] struct {
	length int
	blocks ArrayStack[[]T]
}

func (s *RootishArrayStack[T]) Size() int {
	return s.length
}

func (*RootishArrayStack[T]) iToBlock(i int) int {
	return int(math.Ceil(-3 + math.Sqrt(9+8*float64(i))/2))
}

func (s *RootishArrayStack[T]) Get(i int) (T, error) {
	if i < 0 || s.length <= i {
		return zero.New[T](), ErrIndexOutOfRange
	}
	b := s.iToBlock(i)
	j := i - b*(b+1)/2
	block, err := s.blocks.Get(j)
	if err != nil {
		return zero.New[T](), fmt.Errorf("s.blocks.Get(%d): %w", j, err)
	}
	return block[j], nil
}

func (s *RootishArrayStack[T]) Set(i int, x T) (T, error) {
	if i < 0 || s.length <= i {
		return zero.New[T](), ErrIndexOutOfRange
	}
	b := s.iToBlock(i)
	j := i - b*(b+1)/2
	block, err := s.blocks.Get(j)
	if err != nil {
		return zero.New[T](), fmt.Errorf("s.blocks.Get(%d): %w", j, err)
	}
	y := block[j]
	block[j] = x
	return y, nil
}

func (s *RootishArrayStack[T]) Add(i int, x T) error {
	if i < 0 || s.length+1 <= i {
		return ErrIndexOutOfRange
	}
	r := s.blocks.Size()
	if r*(r+1)/2 < s.length+1 {
		s.glow()
	}
	s.length++
	s.rotateRight(i, s.length-1)
	s.Set(i, x)
	return nil
}

func (s *RootishArrayStack[T]) Remove(i int) (T, error) {
	if i < 0 || s.length <= i {
		return zero.New[T](), ErrIndexOutOfRange
	}
	x, err := s.Get(i)
	if err != nil {
		return zero.New[T](), err
	}
	s.rotateLeft(i+1, s.length)
	s.length--

	r := s.blocks.Size()
	if s.length <= (r-2)*(r-1)/2 {
		s.shrink()
	}

	return x, nil
}

func (s *RootishArrayStack[T]) Push(x T) {
	// never occur ErrIndexOutOfRange, so we can ignore error
	_ = s.Add(s.length, x)
}

func (s *RootishArrayStack[T]) Pop() (T, error) {
	return s.Remove(s.length - 1)
}

func (s *RootishArrayStack[T]) rotateLeft(start, end int) {
	for i := start - 1; i < end-1; i++ {
		x, _ := s.Get(i + 1)
		s.Set(i, x)
	}
}

func (s *RootishArrayStack[T]) rotateRight(start, end int) {
	for i := end; start < i; i-- {
		x, _ := s.Get(i - 1)
		s.Set(i, x)
	}
}

func (s *RootishArrayStack[T]) glow() {
	s.blocks.Push(make([]T, s.blocks.Size()+1))
}

func (s *RootishArrayStack[T]) shrink() {
	r := s.blocks.Size()
	for 0 < r && s.length <= (r-2)*(r-1)/2 {
		s.blocks.Pop()
		r--
	}
}
