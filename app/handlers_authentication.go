package app

import (
	"strconv"
	"time"

	appuser "github.com/TudorHulban/authentication/domain/app-user"
	"github.com/TudorHulban/authentication/pages"
	"github.com/TudorHulban/authentication/services/suser"
	"github.com/gofiber/fiber/v2"

	g "github.com/maragudk/gomponents"
	co "github.com/maragudk/gomponents/components"
)

func (a *App) HandlerLoginPage(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/html")

	return pageLogin("HandlerLoginPage").Render(c)
}

func (a *App) HandlerLoggedInPage(c *fiber.Ctx) error {
	userLogged, errGetUser := appuser.ExtractLoggedUserFrom(c.Context())
	if errGetUser != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	page := co.HTML5(
		co.HTML5Props{
			Title:       "Home",
			Description: "HTMX Logged",
			Language:    "English",
			Head: []g.Node{
				pages.ScriptHTMX,
				pages.LinkCSSWater,
			},
			Body: []g.Node{
				pages.Header(),
				pages.UserSalutation(userLogged),
				pages.Navigation(
					&pages.ParamsNavigation{
						WhereTo:        a.baseURL() + RouteTickets,
						LabelToDisplay: "Tickets",
					},
				),
				pages.Footer(),
			},
		},
	)

	c.Set("Content-Type", "text/html")

	return page.Render(c)
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

		return pageLogin("HandlerLoginRequest - a.ServiceUser.GetUser").Render(c)
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

	c.Set("HX-Redirect", RouteLogged)

	return c.SendStatus(fiber.StatusOK)
}
