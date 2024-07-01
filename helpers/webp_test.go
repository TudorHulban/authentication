package helpers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConvertToWebP(t *testing.T) {
	require.NoError(t,
		ConvertToWebP(
			"../cmd/public/images/pexels-pixabay-159298.jpg",
			"../cmd/public/images/pexels-pixabay-159298.webp",
		),
	)
}
