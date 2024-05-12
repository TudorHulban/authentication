package app

import (
	"strconv"

	"github.com/TudorHulban/authentication/services/suser"
	"github.com/gofiber/fiber/v2"
)

func (a *App) HandlerLoginPage(c *fiber.Ctx) error {
	return c.Render(
		"pages/login",
		nil,
		"layouts/base",
	)
}

func (a *App) HandlerLoggedInPage(c *fiber.Ctx) error {
	sessionID, errConvert := strconv.Atoi(
		c.Cookies(CookieLoggedUser),
	)
	if errConvert != nil {
		c.Set("HX-Redirect", RouteLogin)
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	cachedUser, errGet := a.ServiceSessions.GetUser(
		int64(sessionID),
	)
	if errGet != nil {
		c.Set("HX-Redirect", RouteLogin)
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.Render(
		"pages/logged",
		fiber.Map{
			"name": cachedUser.Name,
		},
		"layouts/base",
	)
}

func (a *App) HandlerLoginRequest(c *fiber.Ctx) error {
	var params suser.ParamsGetUser

	if err := c.BodyParser(&params); err != nil {
		return err
	}

	reconstructedUser, errGetItem := a.ServiceUser.GetUser(
		c.Context(),
		&params,
	)
	if errGetItem != nil {
		return c.Render(
			"components/form_input_error",
			fiber.Map{
				"Email":    "Username or email incorrect",
				"Password": "Username or email incorrect",

				"errors": errGetItem,
			},
		)
	}

	sessionID, errCacheLoggedUser := a.ServiceSessions.PutUserTTL(
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
			),
		},
	)

	c.Set("HX-Redirect", RouteLogged)
	return c.SendStatus(fiber.StatusOK)
}
