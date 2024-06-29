package helpers

import "github.com/dolthub/swiss"

func ForEachSlice[T, R any](values []T, process func(T) R) []R {
	result := make([]R, len(values), len(values))

	for ix, value := range values {
		result[ix] = process(value)
	}

	return result
}

func FilterWith(value string, set *swiss.Map[uint8, string]) []string {
	var result []string

	set.Iter(
		func(k uint8, v string) (stop bool) {
			if v != value {
				result = append(result, v)
			}

			return false
		},
	)

	return result
}
