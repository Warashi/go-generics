package minmax

import (
	"golang.org/x/exp/constraints"

	"github.com/Warashi/go-generics/zero"
)

func Min[T constraints.Ordered](values ...T) T {
	if len(values) == 0 {
		return zero.New[T]()
	}
	m := values[0]
	if len(values) == 1 {
		return m
	}
	for _, v := range values[1:] {
		if v < m {
			m = v
		}
	}
	return m
}

func Max[T constraints.Ordered](values ...T) T {
	if len(values) == 0 {
		return zero.New[T]()
	}
	m := values[0]
	if len(values) == 1 {
		return m
	}
	for _, v := range values[1:] {
		if m < v {
			m = v
		}
	}
	return m
}
