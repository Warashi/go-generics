package container_test

import (
	"testing"

	"github.com/Warashi/go-generics/container"
)

func TestArrayDequeue(t *testing.T) {
	t.Run("Push-Pop/Front", func(t *testing.T) {
		s := container.ArrayDequeue[int]{}
		for i := 0; i < 5; i++ {
			s.PushFront(i)
		}
		for i := 4; 0 <= i; i-- {
			got, err := s.PopFront()
			if err != nil {
				t.Errorf("s.PopFront() returned error %v, wantErr %v", err, false)
			}
			if got != i {
				t.Errorf("s.PopFront() = %d, want %d", got, i)
			}
		}
	})
	t.Run("Push-Pop/Back", func(t *testing.T) {
		s := container.ArrayDequeue[int]{}
		for i := 0; i < 5; i++ {
			s.PushBack(i)
		}
		for i := 4; 0 <= i; i-- {
			got, err := s.PopBack()
			if err != nil {
				t.Errorf("s.PopBack() returned error %v, wantErr %v", err, false)
			}
			if got != i {
				t.Errorf("s.PopBack() = %d, want %d", got, i)
			}
		}
	})
	t.Run("Pop does not panic/Front", func(t *testing.T) {
		s := container.ArrayDequeue[int]{}
		if _, err := s.PopFront(); err == nil {
			t.Errorf("s.PopFront() returned error %v, wantErr %v", err, true)
		}
	})
	t.Run("Pop does not panic/Back", func(t *testing.T) {
		s := container.ArrayDequeue[int]{}
		if _, err := s.PopBack(); err == nil {
			t.Errorf("s.PopBack() returned error %v, wantErr %v", err, true)
		}
	})
	t.Run("Set-Get", func(t *testing.T) {
		s := container.ArrayDequeue[int]{}
		for i := 0; i < 5; i++ {
			s.PushFront(0)
		}
		for i := 0; i < 5; i++ {
			s.Set(i, i)
			got, err := s.Get(i)
			if err != nil {
				t.Errorf("s.Get() returned error %v, wantErr %v", err, false)
			}
			want := i
			if got != want {
				t.Errorf("s.Get() = %v, want %v", got, want)
			}
		}
	})
}
