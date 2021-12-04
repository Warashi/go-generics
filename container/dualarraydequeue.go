package container

import (
	"github.com/Warashi/go-generics/minmax"
	"github.com/Warashi/go-generics/zero"
)

type DualArrayDequeue[T any] struct {
	front ArrayStack[T]
	back  ArrayStack[T]
}

func (q *DualArrayDequeue[T]) Size() int {
	return q.front.Size() + q.back.Size()
}

func (q *DualArrayDequeue[T]) Get(i int) (T, error) {
	if i < q.front.Size() {
		return q.front.Get(q.front.Size() - i - 1)
	}
	if i < q.Size() {
		return q.back.Get(i - q.front.Size())
	}
	return zero.New[T](), ErrIndexOutOfRange
}

func (q *DualArrayDequeue[T]) Set(i int, x T) error {
	if i < q.front.Size() {
		return q.front.Set(q.front.Size()-i-1, x)
	}
	if i < q.Size() {
		return q.back.Set(i-q.front.Size(), x)
	}
	return ErrIndexOutOfRange
}

func (q *DualArrayDequeue[T]) Add(i int, x T) error {
	defer q.balance()
	if i < q.front.Size() {
		return q.front.Add(q.front.Size()-i, x)
	}
	if i < q.Size()+1 {
		return q.back.Add(i-q.front.Size(), x)
	}
	return ErrIndexOutOfRange
}

func (q *DualArrayDequeue[T]) Remove(i int) (T, error) {
	defer q.balance()
	if i < q.front.Size() {
		return q.front.Remove(q.front.Size() - i - 1)
	}
	if i < q.Size() {
		return q.back.Remove(i - q.front.Size())
	}
	return zero.New[T](), ErrIndexOutOfRange
}

func (q *DualArrayDequeue[T]) balance() {
	if q.back.Size() < 3*q.front.Size() && q.front.Size() < 3*q.back.Size() {
		return
	}
	size := q.Size()
	newFrontSize := size / 2
	newFront := make([]T, minmax.Max(2*newFrontSize, 1))
	for i := 0; i < newFrontSize; i++ {
		newFront[newFrontSize-i-1], _ = q.Get(i)
	}
	newBackSize := size - newFrontSize
	newBack := make([]T, minmax.Max(2*newBackSize, 1))
	for i := 0; i < newBackSize; i++ {
		newBack[i], _ = q.Get(newFrontSize + i)
	}

	q.front.array = newFront
	q.front.length = newFrontSize
	q.back.array = newBack
	q.back.length = newBackSize
}

func (s *DualArrayDequeue[T]) PushFront(x T) {
	// never occur ErrIndexOutOfRange, so we can ignore error
	_ = s.Add(0, x)
}

func (s *DualArrayDequeue[T]) PopFront() (T, error) {
	return s.Remove(0)
}

func (s *DualArrayDequeue[T]) PushBack(x T) {
	// never occur ErrIndexOutOfRange, so we can ignore error
	_ = s.Add(s.Size(), x)
}

func (s *DualArrayDequeue[T]) PopBack() (T, error) {
	return s.Remove(s.Size() - 1)
}
