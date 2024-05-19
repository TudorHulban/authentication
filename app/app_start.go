package app

import (
	"github.com/TudorHulban/authentication/fixtures"
	"github.com/TudorHulban/authentication/services/suser"
	"github.com/TudorHulban/log"
	"github.com/gofiber/fiber/v2"
	fiberlog "github.com/gofiber/fiber/v2/log"
)

func (a *App) root() string {
	return ":" + a.port
}

func (a *App) baseURL() string {
	return a.host + a.root()
}

func (a *App) Start() error {
	fiberlog.SetLogger(
		log.NewLogger(log.DEBUG, output, true),
	)

	var mw func(c *fiber.Ctx) error

	if a.authenticationDisabled {
		mw = a.MwPassThrough(
			&suser.ParamsGetUser{
				Email:    fixtures.TestUser.Email,
				Password: fixtures.TestUser.Password,
			},
		)
	} else {
		mw = a.MwAuthentication()
	}

	a.Transport.Use(
		[]string{
			RouteLogged,
			RouteTask,
			RouteTasks,
		},
		mw,
	)

	InitializeTransportRoutes(a)

	return a.Transport.Listen(a.root())
}
