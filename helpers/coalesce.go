package helpers

func Coalesce[T comparable](first, second T) T {
	var zeroValue T

	if first == zeroValue {
		return second
	}

	return first
}
