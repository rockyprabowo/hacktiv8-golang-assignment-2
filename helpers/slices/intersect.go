package slices

import (
	"golang.org/x/exp/constraints"
	"rocky.my.id/git/h8-assignment-2/helpers/set"
)

// Intersect
// Returns the intersection of two slice with a type of T without the duplicates.
func Intersect[T constraints.Ordered](a, b []T) (diff []T) {
	long, short := SliceLongShort(a, b)

	sets := set.NewSet[T]()

	for _, v := range short {
		sets.Add(v)
	}

	for _, v := range long {
		if sets.Has(v) {
			diff = append(diff, v)
		}
	}

	return
}
