package infra

import (
	"github.com/TudorHulban/authentication/app"
)

func InitializeTransportRoutes(application *app.App) {
	application.Transport.Get(
		app.RouteLogin,
		application.HandlerLoginPage,
	)

	application.Transport.Get(
		app.RouteLogged,
		application.HandlerLoggedInPage,
	)

	application.Transport.Post(
		app.RouteLogin,
		application.HandlerLoginRequest,
	)
}
