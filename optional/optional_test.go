package optional

import (
	"testing"

	"github.com/Warashi/go-generics/pointer"
)

func TestOptional_IsEmpty(t *testing.T) {
	tests := []struct {
		name  string
		value *int
		want  bool
	}{
		{
			name:  "empty",
			value: nil,
			want:  true,
		},
		{
			name:  "zero",
			value: pointer.New(0),
			want:  false,
		},
		{
			name:  "non-zero",
			value: pointer.New(1),
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := Optional[int]{value: tt.value}
			if got := o.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want = %v", got, tt.want)
			}
		})
	}
}
