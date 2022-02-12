package monad_test

import "testing"
import "github.com/Warashi/go-generics/monad"
import "github.com/Warashi/go-generics/slice"
import "github.com/Warashi/go-generics/types"
import "github.com/Warashi/go-generics/optional"
import "github.com/google/go-cmp/cmp"

func TestMap(t *testing.T) {
	t.Run("slice", func(t *testing.T) {
		odd := types.Function[int, bool](types.Closure[int, bool](func(v int) bool { return v%2 != 0 }))

		s1 := []int{1, 4, 7, 10}
		s2 := monad.Map[[]bool](slice.MonadImpl[int, bool]{}, s1, odd)

		if got, want := s2, []bool{true, false, true, false}; !cmp.Equal(got, want) {
			t.Errorf("got %v, want %v, diff %v", got, want, cmp.Diff(got, want))
		}
	})
	t.Run("option", func(t *testing.T) {
		odd := types.Function[int, bool](types.Closure[int, bool](func(v int) bool { return v%2 != 0 }))

		s1 := optional.New(1)
		s2 := monad.Map[optional.Optional[bool]](optional.MonadImpl[int, bool]{}, s1, odd)

		if got, want := s2, optional.New(true); !cmp.Equal(got, want) {
			t.Errorf("got %v, want %v, diff %v", got, want, cmp.Diff(got, want))
		}
	})
}
func TestFilter(t *testing.T) {
	t.Run("slice", func(t *testing.T) {
		odd := types.Function[int, bool](types.Closure[int, bool](func(v int) bool { return v%2 != 0 }))

		s1 := []int{1, 4, 7, 10}
		s2 := monad.Filter(slice.MonadImpl[int, int]{}, s1, odd)

		if got, want := s2, []int{1, 7}; !cmp.Equal(got, want) {
			t.Errorf("got %v, want %v, diff %v", got, want, cmp.Diff(got, want))
		}
	})
	t.Run("option", func(t *testing.T) {
		odd := types.Function[int, bool](types.Closure[int, bool](func(v int) bool { return v%2 != 0 }))
		impl := optional.MonadImpl[int, int]{}

		cases := map[string]struct {
			in   optional.Optional[int]
			want optional.Optional[int]
		}{
			"odd":   {optional.New(1), optional.New(1)},
			"even":  {optional.New(2), optional.Empty[int]()},
			"empty": {optional.Empty[int](), optional.Empty[int]()},
		}

		for n, c := range cases {
			n, c := n, c
			t.Run(n, func(t *testing.T) {
				if got := monad.Filter(impl, c.in, odd); !cmp.Equal(got, c.want) {
					t.Errorf("got %v, want %v, diff %v", got, c.want, cmp.Diff(got, c.want))
				}
			})
		}

		t.Run("non-empty", func(t *testing.T) {
			s1 := optional.New(1)
			s2 := monad.Filter(impl, s1, odd)

			if got, want := s2, optional.New(1); !cmp.Equal(got, want) {
				t.Errorf("got %v, want %v, diff %v", got, want, cmp.Diff(got, want))
			}
		})
	})
}
