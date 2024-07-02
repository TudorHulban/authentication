package helpers

func High[T any](slice []T) T {
	if slice == nil {
		var zero T

		return zero
	}

	return slice[len(slice)-1]
}

func Low[T any](slice []T) T {
	if slice == nil {
		var zero T

		return zero
	}

	return slice[0]
}
