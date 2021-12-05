package sequence_test

import (
	"testing"

	"github.com/Warashi/go-generics/sequence"
	"github.com/google/go-cmp/cmp"
)

func TestOf(t *testing.T) {
	s := sequence.Of(0, 1, 2)
	wants := []int{0, 1, 2}
	for _, want := range wants {
		if ok := s.Next(); !ok {
			t.Fatalf("s.Next() = %v, want %v", ok, true)
		}
		if v := s.Value(); v != want {
			t.Fatalf("s.Value() = %v, want %v", v, want)
		}
	}
}

func TestMap(t *testing.T) {
	type f = sequence.ApplyerFunc[int, int]
	s := sequence.Map[int, int](
		sequence.Of(0, 1, 2),
		f(func(x int) int { return x + 1 }),
	)
	wants := []int{1, 2, 3}
	for _, want := range wants {
		if ok := s.Next(); !ok {
			t.Fatalf("s.Next() = %v, want %v", ok, true)
		}
		if v := s.Value(); v != want {
			t.Fatalf("s.Value() = %v, want %v", v, want)
		}
	}
}

func TestFlatten(t *testing.T) {
	s := sequence.Flatten(
		sequence.Of(
			sequence.Of(0, 1, 2),
			sequence.Of(3, 4, 5),
		))
	wants := []int{0, 1, 2, 3, 4, 5}
	for _, want := range wants {
		if ok := s.Next(); !ok {
			t.Fatalf("s.Next() = %v, want %v", ok, true)
		}
		if v := s.Value(); v != want {
			t.Fatalf("s.Value() = %v, want %v", v, want)
		}
	}
}

func TestMapSlice(t *testing.T) {
	type f = sequence.ApplyerFunc[int, sequence.Sequence[int]]
	s := sequence.Map[int, sequence.Sequence[int]](
		sequence.Of(0, 1, 2),
		f(func(i int) sequence.Sequence[int] { return sequence.Of(i, i*2) }),
	)
	wants := [][]int{{0, 0}, {1, 2}, {2, 4}}
	for _, want := range wants {
		if ok := s.Next(); !ok {
			t.Fatalf("s.Next() = %v, want %v", ok, true)
		}
		if v := s.Value(); !cmp.Equal(sequence.Collect(v), want) {
			t.Fatalf("s.Value() = %v, want %v", v, want)
		}
	}
}

func TestFlatMap(t *testing.T) {
	type f = sequence.ApplyerFunc[int, sequence.Sequence[int]]
	s := sequence.FlatMap[int, int](
		sequence.Of(0, 1, 2),
		f(func(i int) sequence.Sequence[int] { return sequence.Of(i, i*2) }),
	)
	wants := []int{0, 0, 1, 2, 2, 4}
	for i, want := range wants {
		if ok := s.Next(); !ok {
			t.Fatalf("s.Next() = %v, want %v", ok, true)
		}
		if v := s.Value(); v != want {
			t.Fatalf("s.Value() = %v, want %v, i = %d", v, want, i)
		}
	}
}

func TestFilter(t *testing.T) {
	type f = sequence.ApplyerFunc[int, bool]
	s := sequence.Filter[int](
		sequence.Of(0, 1, 2, 3, 4, 5, 6),
		f(func(i int) bool { return i%2 == 0 }),
	)
	wants := []int{0, 2, 4, 6}
	for _, want := range wants {
		if ok := s.Next(); !ok {
			t.Fatalf("s.Next() = %v, want %v", ok, true)
		}
		if v := s.Value(); v != want {
			t.Fatalf("s.Value() = %v, want %v", v, want)
		}
	}
}
