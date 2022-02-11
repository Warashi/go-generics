package monad_test

import "testing"
import "github.com/Warashi/go-generics/monad"
import "github.com/Warashi/go-generics/slice"
import "github.com/Warashi/go-generics/types"
import "github.com/Warashi/go-generics/optional"
import "github.com/google/go-cmp/cmp"

func TestMap(t *testing.T) {
	t.Run("slice", func(t *testing.T) {
		odd := types.Applyer[int, bool](types.ApplyerFunc[int, bool](func(v int) bool { return v%2 != 0 }))

		s1 := []int{1, 4, 7, 10}
		s2 := monad.Map[[]bool](slice.MonadImpl[int, bool]{}, s1, odd)

		if got, want := s2, []bool{true, false, true, false}; !cmp.Equal(got, want) {
			t.Errorf("got %v, want %v, diff %v", got, want, cmp.Diff(got, want))
		}
	})
	t.Run("option", func(t *testing.T) {
		odd := types.Applyer[int, bool](types.ApplyerFunc[int, bool](func(v int) bool { return v%2 != 0 }))

		s1 := optional.New(1)
		s2 := monad.Map[optional.Optional[bool]](optional.MonadImpl[int, bool]{}, s1, odd)

		if got, want := s2, optional.New(true); !cmp.Equal(got, want) {
			t.Errorf("got %v, want %v, diff %v", got, want, cmp.Diff(got, want))
		}
	})
}
