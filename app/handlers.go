package app

import (
	"fmt"

	"github.com/TudorHulban/authentication/services/suser"
	"github.com/gofiber/fiber/v2"
)

func (a *App) LoginPageHandler(c *fiber.Ctx) error {
	return c.Render(
		"pages/login",
		nil,
		"layouts/base",
	)
}

func (a *App) LoggedInPageHandler(c *fiber.Ctx) error {
	return c.Render(
		"pages/logged-in",
		nil,
		"layouts/base",
	)
}

func (a *App) LoginRequestHandler(c *fiber.Ctx) error {
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

	fmt.Println(reconstructedUser)

	c.Set("HX-Redirect", "/logged-in")
	return c.SendStatus(200)
}
