package app

func InitializeTransportRoutes(application *App) {
	// Authentication

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
		application.HandlerHomePage,
	)

	application.Transport.Post(
		RouteLogin,
		application.HandlerLoginRequest,
	)

	// Ticket

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

	// Event

	application.Transport.Post(
		RouteEvent,
		application.HandlerAddEvent,
	)
}
