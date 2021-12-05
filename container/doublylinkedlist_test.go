package container_test

import (
	"testing"

	"github.com/Warashi/go-generics/container"
)

func TestDoublyLinkedList(t *testing.T) {
	t.Run("Push-Pop/Front-Front", func(t *testing.T) {
		s := container.DoublyLinkedList[int]{}
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
	t.Run("Push-Pop/Back-Front", func(t *testing.T) {
		s := container.DoublyLinkedList[int]{}
		for i := 0; i < 5; i++ {
			s.PushBack(i)
		}
		for i := 0; i < 5; i++ {
			got, err := s.PopFront()
			if err != nil {
				t.Errorf("s.PopBack() returned error %v, wantErr %v", err, false)
			}
			if got != i {
				t.Errorf("s.PopBack() = %d, want %d", got, i)
			}
		}
	})
	t.Run("Push-Pop/Front-Back", func(t *testing.T) {
		s := container.DoublyLinkedList[int]{}
		for i := 0; i < 5; i++ {
			s.PushFront(i)
		}
		for i := 0; i < 5; i++ {
			got, err := s.PopBack()
			if err != nil {
				t.Errorf("s.PopFront() returned error %v, wantErr %v", err, false)
			}
			if got != i {
				t.Errorf("s.PopFront() = %d, want %d", got, i)
			}
		}
	})
	t.Run("Push-Pop/Back-Back", func(t *testing.T) {
		s := container.DoublyLinkedList[int]{}
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
		s := container.DoublyLinkedList[int]{}
		if _, err := s.PopFront(); err == nil {
			t.Errorf("s.PopFront() returned error %v, wantErr %v", err, true)
		}
	})
	t.Run("Pop does not panic/Back", func(t *testing.T) {
		s := container.DoublyLinkedList[int]{}
		if _, err := s.PopBack(); err == nil {
			t.Errorf("s.PopBack() returned error %v, wantErr %v", err, true)
		}
	})
}
