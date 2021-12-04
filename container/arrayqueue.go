package container

import (
	"github.com/Warashi/go-generics/minmax"
	"github.com/Warashi/go-generics/zero"
)

type ArrayQueue[T any] struct {
	start, length int
	array         []T
}

func (q *ArrayQueue[T]) Size() int {
	return q.length
}

func (q *ArrayQueue[T]) resize() {
	n := make([]T, minmax.Max(1, 2*q.length))
	if q.array != nil {
		l := minmax.Min(q.length, len(q.array)-q.start)
		copy(n[:l], q.array[q.start:])
		copy(n[l:], q.array[:(q.start+q.length)%len(q.array)])
	}
	q.start = 0
	q.array = n
}

func (q *ArrayQueue[T]) Add(x T) {
	if len(q.array) <= q.length+1 {
		q.resize()
	}
	q.array[(q.start+q.length)%len(q.array)] = x
	q.length++
}

func (q *ArrayQueue[T]) Remove() (T, error) {
	if q.length <= 0 {
		return zero.New[T](), ErrIndexOutOfRange
	}
	x := q.array[q.start]
	q.start = (q.start + 1) % len(q.array)
	q.length--
	if 3*q.length < len(q.array) {
		q.resize()
	}
	return x, nil
}
