package channel

import (
	"sync"

	"github.com/Warashi/go-generics/container"
)

type queue[T any] struct {
	once  sync.Once
	mu    sync.Mutex
	cond  *sync.Cond
	array container.ArrayQueue[T]
	done  bool
}

func (q *queue[T]) setup() {
	q.once.Do(func() {
		q.mu.Lock()
		defer q.mu.Unlock()
		q.cond = sync.NewCond(&q.mu)
	})
}

func (q *queue[T]) Add(x T) {
	q.setup()
	q.mu.Lock()
	defer q.cond.Signal()
	defer q.mu.Unlock()
	q.array.Add(x)
}

func (q *queue[T]) Channel() chan T {
	q.setup()
	c := make(chan T)
	go func() {
		defer close(c)
		for {
			q.cond.L.Lock()
		wait:
			for {
				if q.array.Size() > 0 {
					break wait
				}
				if q.done {
					q.cond.L.Unlock()
					return
				}
				q.cond.Wait()
			}
			v, err := q.array.Remove()
			if err != nil {
				panic(err)
			}
			c <- v
			q.cond.L.Unlock()
		}
	}()
	return c
}

func (q *queue[T]) Close() {
	q.mu.Lock()
	defer q.cond.Broadcast()
	defer q.mu.Unlock()
	q.done = true
}

func Infinite[T any]() (in chan<- T, out <-chan T) {
	i := make(chan T)
	var q queue[T]
	go func() {
		defer q.Close()
		for v := range i {
			q.Add(v)
		}
	}()
	return i, q.Channel()
}
