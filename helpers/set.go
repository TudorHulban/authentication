package helpers

import (
	"github.com/dolthub/swiss"
)

type KV[K, V comparable] struct {
	Key   K
	Value V
}

func NewImmutableSetFrom[K, V comparable](values []KV[K, V]) *swiss.Map[K, V] {
	result := swiss.NewMap[K, V](uint32(len(values)))

	for _, value := range values {
		result.Put(
			value.Key,
			value.Value,
		)
	}

	return result
}
