package fixtures

import (
	"context"

	testuser "github.com/TudorHulban/authentication/fixtures/test-user"
	"github.com/TudorHulban/authentication/helpers"
	"github.com/TudorHulban/authentication/services/suser"
)

type PiersFixtureAddTestUser struct {
	ServiceUser *suser.Service
}

func FixtureAddTestUser(ctx context.Context, piers *PiersFixtureAddTestUser) (helpers.PrimaryKey, error) {
	return piers.ServiceUser.CreateUser(
		ctx,
		&suser.ParamsCreateUser{
			Email:    testuser.TestUser.Email,
			Password: testuser.TestUser.Password,
			Name:     testuser.TestUser.Name,
		},
	)
}
