package app

import (
	"strconv"

	appuser "github.com/TudorHulban/authentication/domain/app-user"
	"github.com/gofiber/fiber/v2"
)

func (a *App) MwAuthentication() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		sessionID, errConvert := strconv.Atoi(
			c.Cookies(CookieLoggedUser),
		)
		if errConvert != nil {
			c.Set("Content-Type", "text/html")

			return _pageLogin.Render(c)
		}

		cachedUser, errGet := a.serviceSessions.GetUser(
			int64(sessionID),
		)
		if errGet != nil {
			c.Set("Content-Type", "text/html")

			return _pageLogin.Render(c)
		}

		c.Locals(
			appuser.KeyLoggedUser{},
			cachedUser,
		)

		return c.Next()
	}
}
