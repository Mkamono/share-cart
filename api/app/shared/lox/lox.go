package lox

import "github.com/samber/lo"

func FromSlicePtr[T any](collection []*T) []T {
	return lo.Map(collection, func(x *T, _ int) T {
		if x == nil {
			return lo.Empty[T]()
		}
		return *x
	})
}
