package helpers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSet(t *testing.T) {
	m := NewImmutableSetFrom[string, struct{}](
		[]KV[string, struct{}]{
			{
				Key: "1",
			},
			{
				Key: "2",
			},
		},
	)

	require.True(t,
		m.Has("1"),
	)

	require.True(t,
		m.Has("2"),
	)

	require.False(t,
		m.Has("3"),
	)
}
