package slices

func Map[T any, U any](slice []T, f func(T) U) (mapped []U) {
	mapped = make([]U, len(slice))

	for i, v := range slice {
		mapped[i] = f(v)
	}

	return
}

func Filter[T any](slice []T, f func(T) bool) (filtered []T) {
	for _, v := range slice {
		if f(v) {
			filtered = append(filtered, v)
		}
	}
	return
}
