package fixtures

import (
	"context"

	"github.com/TudorHulban/authentication/app"
	"github.com/TudorHulban/authentication/services/suser"
)

func InitializeAddTestUser(ctx context.Context, application *app.App) error {
	return application.ServiceUser.CreateUser(
		ctx,
		&suser.ParamsCreateUser{
			Email:    testUser.Email,
			Password: testUser.Password,
			Name:     testUser.Name,
		},
	)
}
