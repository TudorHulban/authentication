package appuser

import (
	"context"

	"github.com/TudorHulban/authentication/apperrors"
)

type KeyLoggedUser struct{}

func injectLoggedUserIn(ctx context.Context, user *User) context.Context {
	return context.WithValue(
		ctx,
		KeyLoggedUser{},
		user,
	)
}

func extractLoggedUserFrom(ctx context.Context) (*User, error) {
	loggedUser := ctx.Value(
		KeyLoggedUser{},
	)
	if loggedUser == nil {
		return nil,
			apperrors.ErrContextValueNotFound{
				Value: "logged user",
			}
	}

	return loggedUser.(*User),
		nil
}
