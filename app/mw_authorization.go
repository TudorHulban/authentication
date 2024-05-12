package app

import (
	"strconv"

	appuser "github.com/TudorHulban/authentication/domain/app-user"
	"github.com/gofiber/fiber/v2"
)

func (a *App) MwAuthorization() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		sessionID, errConvert := strconv.Atoi(
			c.Cookies(CookieLoggedUser),
		)
		if errConvert != nil {
			return c.Render(
				"pages/login",
				nil,
				"layouts/base",
			)
		}

		cachedUser, errGet := a.ServiceSessions.GetUser(
			int64(sessionID),
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
			cachedUser,
		)

		return c.Next()
	}
}
