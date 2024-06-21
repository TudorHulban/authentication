package app

import (
	"github.com/TudorHulban/authentication/app/constants"
	testuser "github.com/TudorHulban/authentication/fixtures/test-user"
	"github.com/TudorHulban/authentication/services/suser"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func (a *App) root() string {
	return ":" + a.port
}

func (a *App) baseURL() string {
	return a.host + a.root()
}

func (a *App) Start() error {
	var mw func(c *fiber.Ctx) error

	if a.authenticationDisabled {
		mw = a.MwPassThrough(
			&suser.ParamsGetUser{
				Email:    testuser.TestUser.Email,
				Password: testuser.TestUser.Password,
			},
		)
	} else {
		mw = a.MwAuthentication()
	}

	a.Transport.Use(
		logger.New(),

		[]string{
			constants.RouteHome,
			constants.RouteTicket,
			constants.RouteTickets,
			constants.RouteTicketEvent,
		},

		mw,
	)

	a.Transport.Static("/public", "../public")

	InitializeTransportRoutes(a)

	return a.Transport.Listen(a.root())
}
