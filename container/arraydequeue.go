package container

import (
	"github.com/Warashi/go-generics/minmax"
	"github.com/Warashi/go-generics/zero"
)

type ArrayDequeue[T any] struct {
	start, length int
	array         []T
}

func (q *ArrayDequeue[T]) Get(i int) (T, error) {
	if i < 0 || q.length <= i {
		return zero.New[T](), ErrIndexOutOfRange
	}
	return q.array[(q.start+i)%len(q.array)], nil
}

func (q *ArrayDequeue[T]) Set(i int, x T) (T, error) {
	if i < 0 || q.length <= i {
		return zero.New[T](), ErrIndexOutOfRange
	}
	y := q.array[(q.start+i)%len(q.array)]
	q.array[(q.start+i)%len(q.array)] = x
	return y, nil
}

func (q *ArrayDequeue[T]) Add(i int, x T) error {
	if i < 0 || q.length < i {
		return ErrIndexOutOfRange
	}
	if len(q.array) < q.length+1 {
		q.resize()
	}
	if i < q.length/2 {
		if q.start == 0 {
			q.start = len(q.array) - 1
		} else {
			q.start--
		}
		q.rotateLeft(0, i-1)
	} else {
		q.rotateRight(i, q.length)
	}
	q.array[(q.start+i)%len(q.array)] = x
	q.length++
	return nil
}

func (q *ArrayDequeue[T]) Remove(i int) (T, error) {
	if i < 0 || q.length <= i {
		return zero.New[T](), ErrIndexOutOfRange
	}
	x := q.array[(q.start+i)%len(q.array)]
	if i < q.length/2 {
		q.rotateRight(0, i-1)
		q.start = (q.start + 1) % len(q.array)
	} else {
		q.rotateLeft(i+1, q.length-1)
	}
	q.length--
	if 3*q.length < len(q.array) {
		q.resize()
	}
	return x, nil
}

func (q *ArrayDequeue[T]) rotateLeft(start, end int) {
	for k := start; k < end; k++ {
		q.array[(q.start+k)%len(q.array)] =
			q.array[(q.start+k+1)%len(q.array)]
	}
}

func (q *ArrayDequeue[T]) rotateRight(start, end int) {
	for k := end; start < k; k-- {
		q.array[(q.start+k)%len(q.array)] =
			q.array[(q.start+k-1)%len(q.array)]
	}
}

func (q *ArrayDequeue[T]) resize() {
	n := make([]T, minmax.Max(1, 2*q.length))
	if q.array != nil {
		l := minmax.Min(q.length, len(q.array)-q.start)
		copy(n[:l], q.array[q.start:])
		copy(n[l:], q.array[:(q.start+q.length)%len(q.array)])
	}
	q.start = 0
	q.array = n
}

func (s *ArrayDequeue[T]) PushFront(x T) {
	// never occur ErrIndexOutOfRange, so we can ignore error
	_ = s.Add(0, x)
}

func (s *ArrayDequeue[T]) PopFront() (T, error) {
	return s.Remove(0)
}

func (s *ArrayDequeue[T]) PushBack(x T) {
	// never occur ErrIndexOutOfRange, so we can ignore error
	_ = s.Add(s.length, x)
}

func (s *ArrayDequeue[T]) PopBack() (T, error) {
	return s.Remove(s.length - 1)
}
