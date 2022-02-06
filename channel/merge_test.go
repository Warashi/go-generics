package channel_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/Warashi/go-generics/channel"
)

func TestOrDone(t *testing.T) {
	t.Run("normal case", func(t *testing.T) {
		ch := make(chan int, 1)
		och := channel.OrDone(context.Background(), ch)
		ch <- 1
		close(ch)
		if got := <-och; got != 1 {
			t.Errorf("got %d, want %d", got, 1)
		}
	})
	t.Run("context canceled", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		ch := make(chan int)
		och := channel.OrDone(ctx, ch)
		cancel()
		if got, ok := <-och; got != 0 || ok {
			t.Errorf("got %d and %t, want %d and %t", got, ok, 0, false)
		}
	})
}

func TestMerge(t *testing.T) {
	chs := make([]<-chan int, 100)
	for i := range chs {
		ch := make(chan int, 1)
		ch <- i
		close(ch)
		chs[i] = ch
	}

	gots := make([]int, 0, 100)
	for got := range channel.Merge(context.Background(), chs...) {
		gots = append(gots, got)
	}
	want := make([]int, 100)
	for i := range want {
		want[i] = i
	}
	opt := cmpopts.SortSlices(func(i, j int) bool { return i < j })
	if !cmp.Equal(gots, want, opt) {
		t.Errorf("got %v, want %v, diff %v", gots, want, cmp.Diff(gots, want, opt))
	}
}
