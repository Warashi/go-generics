package argminmax_test

import (
	"math/rand"
	"strconv"
	"testing"

	"github.com/Warashi/go-generics/argminmax"
)

func TestArgMax(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "zero length", args: args{data: []int{}}, want: 0},
		{name: "one length", args: args{data: []int{0}}, want: 0},
		{name: "two length, first is smaller", args: args{data: []int{0, 1}}, want: 1},
		{name: "two length, first is larger", args: args{data: []int{1, 0}}, want: 0},
		{name: "three length, third is the largest", args: args{data: []int{0, 1, 2}}, want: 2},
		{name: "three length, second is the largest", args: args{data: []int{0, 2, 1}}, want: 1},
		{name: "three length, third is the largest", args: args{data: []int{1, 0, 2}}, want: 2},
		{name: "three length, second is the largest", args: args{data: []int{1, 2, 0}}, want: 1},
		{name: "three length, first is the largest", args: args{data: []int{2, 0, 1}}, want: 0},
		{name: "three length, first is the largest", args: args{data: []int{2, 1, 0}}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := argminmax.ArgMax(tt.args.data); got != tt.want {
				t.Errorf("ArgMax() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArgMin(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "zero length", args: args{data: []int{}}, want: 0},
		{name: "one length", args: args{data: []int{0}}, want: 0},
		{name: "two length, first is smaller", args: args{data: []int{0, 1}}, want: 0},
		{name: "two length, first is larger", args: args{data: []int{1, 0}}, want: 1},
		{name: "three length, third is the largest", args: args{data: []int{0, 1, 2}}, want: 0},
		{name: "three length, second is the largest", args: args{data: []int{0, 2, 1}}, want: 0},
		{name: "three length, third is the largest", args: args{data: []int{1, 0, 2}}, want: 1},
		{name: "three length, second is the largest", args: args{data: []int{1, 2, 0}}, want: 2},
		{name: "three length, first is the largest", args: args{data: []int{2, 0, 1}}, want: 1},
		{name: "three length, first is the largest", args: args{data: []int{2, 1, 0}}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := argminmax.ArgMin(tt.args.data); got != tt.want {
				t.Errorf("ArgMin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkArgMin(b *testing.B) {
	for i := 1024; i <= 1024*1024; i *= 2 {
		b.Run(strconv.Itoa(i), func(b *testing.B) {
			rand.Seed(1)
			list := make([]int, i)
			for i := range list {
				list[i] = i
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				rand.Shuffle(len(list), func(i, j int) { list[i], list[j] = list[j], list[i] })
				b.StartTimer()
				a := argminmax.ArgMin(list)
				_ = a
			}
		})
	}
}

func BenchmarkArgMax(b *testing.B) {
	for i := 1024; i <= 1024*1024; i *= 2 {
		b.Run(strconv.Itoa(i), func(b *testing.B) {
			rand.Seed(1)
			list := make([]int, i)
			for i := range list {
				list[i] = i
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				rand.Shuffle(len(list), func(i, j int) { list[i], list[j] = list[j], list[i] })
				b.StartTimer()
				a := argminmax.ArgMax(list)
				_ = a
			}
		})
	}
}
