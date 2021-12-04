package minmax

import (
	"constraints"
)

func ArgMin[T constraints.Ordered](data []T) int {
	argmin := 0
	for i := 1; i < len(data); i++ {
		if data[i] < data[argmin] {
			argmin = i
		}
	}
	return argmin
}

func ArgMax[T constraints.Ordered](data []T) int {
	argmax := 0
	for i := 1; i < len(data); i++ {
		if data[argmax] < data[i] {
			argmax = i
		}
	}
	return argmax
}
