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
			if got := s.Pop(); got != i {
				t.Errorf("s.Pop() = %d, want %d", got, i)
			}
		}
	})
}
