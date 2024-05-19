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
			Email:    testUser.Email,
			Password: testUser.Password,
			Name:     testUser.Name,
		},
	)
}
