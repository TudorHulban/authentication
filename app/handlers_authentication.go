package app

import (
	"strconv"

	appuser "github.com/TudorHulban/authentication/domain/app-user"
	"github.com/TudorHulban/authentication/pages"
	"github.com/TudorHulban/authentication/services/suser"
	"github.com/gofiber/fiber/v2"

	g "github.com/maragudk/gomponents"
	co "github.com/maragudk/gomponents/components"
)

func (a *App) HandlerLoginPage(c *fiber.Ctx) error {
	page := co.HTML5(
		co.HTML5Props{
			Title:       "Login",
			Description: "HTMX Login",
			Language:    "English",
			Head: []g.Node{
				pages.ScriptHTMX,
				pages.LinkCSSWater,
			},
			Body: []g.Node{
				pages.Header(),
				pages.FormLogin(),
				pages.Footer(),
			},
		},
	)

	c.Set("Content-Type", "text/html")

	return page.Render(c)
}

func (a *App) HandlerLoggedInPage(c *fiber.Ctx) error {
	user, errGetUser := appuser.ExtractLoggedUserFrom(c.Context())
	if errGetUser != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Render(
		"pages/logged",
		fiber.Map{
			"name":  user.Name,
			"route": a.baseURL() + RouteTickets,
			"title": "Home",
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

				"title": "Login",
			},
		)
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
			),
		},
	)

	c.Set("HX-Redirect", RouteLogged)

	return c.SendStatus(fiber.StatusOK)
}
