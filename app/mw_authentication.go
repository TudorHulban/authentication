package app

import (
	"fmt"
	"strconv"

	appuser "github.com/TudorHulban/authentication/domain/app-user"
	"github.com/TudorHulban/authentication/services/srender"
	"github.com/gofiber/fiber/v2"
)

func (a *App) MwAuthentication() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		sessionID, errConvert := strconv.Atoi(
			c.Cookies(CookieLoggedUser),
		)
		if errConvert != nil {
			c.Set("Content-Type", "text/html")

			return srender.PageLogin(
				fmt.Sprintf("MwAuthentication - SessionID: %d", sessionID),
			).
				Render(c)
		}

		cachedUser, errGet := a.serviceSessions.GetUser(
			int64(sessionID),
		)
		if errGet != nil {
			c.Set("Content-Type", "text/html")

			return srender.PageLogin(
				fmt.Sprintf("MwAuthentication - cachedUser: %s", cachedUser.Name),
			).
				Render(c)
		}

		c.Locals(
			appuser.KeyLoggedUser{},
			cachedUser,
		)

		return c.Next()
	}
}
