package infra

import (
	"github.com/TudorHulban/authentication/app"
)

func InitializeTransportRoutes(application *app.App) {
	application.Transport.Get(
		"/login",
		application.LoginPageHandler,
	)

	application.Transport.Get(
		"/logged-in",
		application.LoggedInPageHandler,
	)

	application.Transport.Post(
		"/login",
		application.LoginRequestHandler,
	)
}
