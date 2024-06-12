package app

import (
	"fmt"

	appuser "github.com/TudorHulban/authentication/domain/app-user"
	"github.com/TudorHulban/authentication/services/srender"
	"github.com/TudorHulban/authentication/services/suser"
	"github.com/gofiber/fiber/v2"
)

func (a *App) MwPassThrough(withUserCredentials *suser.ParamsGetUser) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		reconstructedUser, errGet := a.ServiceUser.GetUserByCredentials(
			c.Context(),
			&suser.ParamsGetUser{
				Email:    withUserCredentials.Email,
				Password: withUserCredentials.Password,
			},
		)
		if errGet != nil {
			c.Set("Content-Type", "text/html")

			fmt.Println(
				"MwPassThrough", errGet,
			)

			return srender.PageLogin("MwPassThrough - a.ServiceUser.GetUser").Render(c)
		}

		c.Locals(
			appuser.KeyLoggedUser{},
			reconstructedUser,
		)

		return c.Next()
	}
}
