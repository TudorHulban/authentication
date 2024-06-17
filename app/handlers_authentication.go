package app

import (
	"strconv"
	"time"

	"github.com/TudorHulban/authentication/app/constants"
	"github.com/TudorHulban/authentication/services/srender"
	"github.com/TudorHulban/authentication/services/suser"
	"github.com/gofiber/fiber/v2"
)

func (a *App) HandlerLoginPage(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/html")

	return srender.PageLogin("HandlerLoginPage").
		Render(c)
}

func (a *App) HandlerLoginRequest(c *fiber.Ctx) error {
	var params suser.ParamsGetUser

	if err := c.BodyParser(&params); err != nil {
		return err
	}

	reconstructedUser, errGetItem := a.ServiceUser.GetUserByCredentials(
		c.Context(),
		&params,
	)
	if errGetItem != nil {
		c.Set("Content-Type", "text/html")

		return srender.PageLogin("HandlerLoginRequest - a.ServiceUser.GetUser").
			Render(c)
	}

	sessionID, errCacheLoggedUser := a.serviceSessions.PutUserTTL(
		reconstructedUser,
	)
	if errCacheLoggedUser != nil {
		return errCacheLoggedUser
	}

	c.Cookie(
		&fiber.Cookie{
			Name: CookieLoggedUser,
			Value: strconv.Itoa(
				int(
					sessionID,
				),
			) + "|" + time.Now().String(),
		},
	)

	c.Set("HX-Redirect", constants.RouteLogged)

	return c.SendStatus(fiber.StatusOK)
}
