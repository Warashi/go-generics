package optional

import (
	"strconv"
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
			value: pointer.Of(0),
			want:  false,
		},
		{
			name:  "non-zero",
			value: pointer.Of(1),
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

func TestOptional_MapOrElseZero(t *testing.T) {
	mapper := MapperFunc[int, string](strconv.Itoa)
	tests := []struct {
		name  string
		value *int
		want  string
	}{
		{
			name:  "empty",
			value: nil,
		},
		{
			name:  "zero",
			value: pointer.Of(0),
			want:  "0",
		},
		{
			name:  "non-zero",
			value: pointer.Of(1),
			want:  "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := Map[int, string](Optional[int]{value: tt.value}, mapper)
			if got := o.OrElseZero(); got != tt.want {
				t.Errorf("OrElseZero() = %v, want = %v", got, tt.want)
			}
		})
	}

}
