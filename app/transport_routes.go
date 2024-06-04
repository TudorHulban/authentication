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

	application.Transport.Post(
		RouteTicket,
		application.HandlerAddTicket,
	)

	application.Transport.Get(
		RouteTickets,
		application.HandlerTickets,
	)

	application.Transport.Get(
		RouteTicket+"/:id",
		application.HandlerTicketID,
	)
}
