package fixtures

import (
	"context"

	"github.com/TudorHulban/authentication/services/suser"
)

type PiersFixtureAddTestUser struct {
	ServiceUser *suser.Service
}

func FixtureAddTestUser(ctx context.Context, piers *PiersFixtureAddTestUser) error {
	return piers.ServiceUser.CreateUser(
		ctx,
		&suser.ParamsCreateUser{
			Email:    TestUser.Email,
			Password: TestUser.Password,
			Name:     TestUser.Name,
		},
	)
}
