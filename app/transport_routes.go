package app

import "github.com/TudorHulban/authentication/app/constants"

func InitializeTransportRoutes(application *App) {
	// Authentication

	application.Transport.Get(
		constants.RoutesAll,
		application.HandlerLoginPage,
	)

	application.Transport.Get(
		constants.RouteLogin,
		application.HandlerLoginPage,
	)

	application.Transport.Get(
		constants.RouteLogged,
		application.HandlerHomePage,
	)

	application.Transport.Post(
		constants.RouteLogin,
		application.HandlerLoginRequest,
	)

	// Ticket

	application.Transport.Post(
		constants.RouteTicket,
		application.HandlerAddTicket,
	)

	application.Transport.Post(
		constants.RouteTickets,
		application.HandlerHTMLTicketsTableBody,
	)

	application.Transport.Get(
		constants.RouteTicketEvents,
		application.HandlerHTMLTicketEventsTable,
	)

	application.Transport.Get(
		constants.RouteTickets,
		application.HandlerHTMLTicketsTable,
	)

	application.Transport.Get(
		constants.RouteTicket+"/:id",
		application.HandlerTicketID,
	)

	// Event

	application.Transport.Post(
		constants.RouteTicketEvent,
		application.HandlerAddEvent,
	)
}
