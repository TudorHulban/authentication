package appuser

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoggedUserInContext(t *testing.T) {
	ctx := context.Background()

	user := User{
		UserCredentials: UserCredentials{
			Email: "x@x.co",
		},
	}

	ctxUpdated := InjectLoggedUserIn(
		ctx,
		&user,
	)

	extractedUser, errGet := ExtractLoggedUserFrom(
		ctxUpdated,
	)
	require.NoError(t, errGet)
	require.NotNil(t, extractedUser)
	require.Equal(t,
		user.Email,
		extractedUser.Email,
	)
}
