package app

import (
	appuser "github.com/TudorHulban/authentication/domain/app-user"
	"github.com/TudorHulban/authentication/services/suser"
	"github.com/gofiber/fiber/v2"
)

func (a *App) MwPassThrough(withUserCredentials *suser.ParamsGetUser) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		reconstructedUser, errGet := a.ServiceUser.GetUser(
			c.Context(),
			&suser.ParamsGetUser{
				Email:    withUserCredentials.Email,
				Password: withUserCredentials.Password,
			},
		)
		if errGet != nil {
			return c.Render(
				"pages/login",
				nil,
				"layouts/base",
			)
		}

		c.Locals(
			appuser.KeyLoggedUser{},
			reconstructedUser,
		)

		return c.Next()
	}
}
