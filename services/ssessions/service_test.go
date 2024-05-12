package ssessions

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestService(t *testing.T) {
	s := NewService()
	require.NotNil(t, s)

	u, errGet := s.GetUser(1)
	require.Error(t, errGet)
	require.Nil(t, u)
}
