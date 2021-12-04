package container_test

import (
	"testing"

	"github.com/Warashi/go-generics/container"
)

func TestArrayQueue(t *testing.T) {
	t.Run("Remove-from-empty", func(t *testing.T) {
		q := container.ArrayQueue[int]{}
		if _, err := q.Remove(); err == nil {
			t.Errorf("q.Remove() returned error %v, wantErr %v", err, true)
		}
	})
	t.Run("Add-Remove/Batch", func(t *testing.T) {
		q := container.ArrayQueue[int]{}
		for i := 0; i < 256; i++ {
			q.Add(i)
		}
		for i := 0; i < 256; i++ {
			got, err := q.Remove()
			if err != nil {
				t.Errorf("q.Remove() returned error %v, wantErr %v", err, false)
			}
			if got != i {
				t.Errorf("q.Remove() = %d, want %d", got, i)
			}
		}
	})
	t.Run("Add-Remove/Iterate", func(t *testing.T) {
		q := container.ArrayQueue[int]{}
		for i := 0; i < 256; i++ {
			q.Add(i)
			got, err := q.Remove()
			if err != nil {
				t.Errorf("q.Remove() returned error %v, wantErr %v", err, false)
			}
			if got != i {
				t.Errorf("q.Remove() = %d, want %d", got, i)
			}
		}
	})
}
