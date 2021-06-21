package auth

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHash(t *testing.T) {
	passRaw := "xxx"
	passSalt := "12345"
	hashCost := 14

	hash, errHash := HASHPassword(passRaw, passSalt, hashCost)
	require.Nil(t, errHash)

	require.True(t, CheckPasswordHash(passRaw, passSalt, string(hash)))
}
