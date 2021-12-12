package channel_test

import (
	"sync"
	"testing"

	"github.com/Warashi/go-generics/channel"
)

func TestInfinite(t *testing.T) {
	const loop = 100
	in, out := channel.Infinite[int]()
	last := -1
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < loop; i++ {
			t.Logf("write %d", i)
			in <- i
		}
		t.Log("write done")
		close(in)
	}()
	for v := range out {
		t.Logf("consume %d", v)
		if last+1 != v {
			t.Fatalf("expected %v, want %v", last+1, v)
		}
		last = v
	}
	if last != loop-1 {
		t.Errorf("last = %v, want %v", last, loop-1)
	}
	wg.Wait()
}
