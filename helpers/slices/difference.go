package slices

import (
	"golang.org/x/exp/constraints"
	"rockyprabowo/assignment-2/helpers/set"
)

func SliceLongShort[T any](a, b []T) (long, short []T) {
	if len(b) > len(a) {
		return b, a
	}
	return a, b
}

func Diff[T constraints.Ordered](a, b []T) (diff []T) {
	long, short := SliceLongShort(a, b)

	sets := set.NewSet[T]()

	for _, v := range short {
		sets.Add(v)
	}

	for _, v := range long {
		if !sets.Has(v) {
			diff = append(diff, v)
		}
	}

	return
}

func DiffWithDuplicates[T constraints.Ordered](a, b []T) (diff []T) {
	long, short := SliceLongShort(a, b)
	longIter, shortIter := 0, 0

	for longIter < len(long) && shortIter < len(short) {
		if long[longIter] == short[longIter] {
			longIter++
			shortIter++
		} else if long[longIter] < short[shortIter] {
			diff = append(diff, long[longIter])
			longIter++
		} else {
			diff = append(diff, short[shortIter])
			shortIter++
		}
	}

	diff = append(diff, a[longIter:]...)
	diff = append(diff, b[shortIter:]...)

	return
}
