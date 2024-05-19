package app

func InitializeTransportRoutes(application *App) {
	application.Transport.Get(
		RoutesAll,
		application.HandlerLoginPage,
	)

	application.Transport.Get(
		RouteLogin,
		application.HandlerLoginPage,
	)

	application.Transport.Get(
		RouteLogged,
		application.HandlerLoggedInPage,
	)

	application.Transport.Post(
		RouteLogin,
		application.HandlerLoginRequest,
	)

	application.Transport.Get(
		RouteTasks,
		application.HandlerTasksPage,
	)
}
