package container_test

import (
	"testing"

	"github.com/Warashi/go-generics/container"
)

func TestArrayStack(t *testing.T) {
	t.Run("Push-Pop", func(t *testing.T) {
		s := container.ArrayStack[int]{}
		for i := 0; i < 5; i++ {
			s.Push(i)
		}
		for i := 4; 0 <= i; i-- {
			got, err := s.Pop()
			if err != nil {
				t.Errorf("s.Pop() returned error %v, wantErr %v", err, false)
			}
			if got != i {
				t.Errorf("s.Pop() = %d, want %d", got, i)
			}
		}
	})
	t.Run("Pop does not panic", func(t *testing.T) {
		s := container.ArrayStack[int]{}
		if _, err := s.Pop(); err == nil {
			t.Errorf("s.Pop() returned error %v, wantErr %v", err, true)
		}
	})
	t.Run("Set-Get", func(t *testing.T) {
		s := container.ArrayStack[int]{}
		for i := 0; i < 5; i++ {
			s.Push(0)
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
