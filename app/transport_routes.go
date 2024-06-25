package app

import "github.com/TudorHulban/authentication/app/constants"

func InitializeTransportRoutes(application *App) {
	// Authentication

	application.Transport.Get(
		constants.RoutesAll,
		application.HandlerLoginPage,
	)

	application.Transport.Get(
		constants.RouteHome,
		application.HandlerHomePage,
	)

	application.Transport.Post(
		constants.RouteLogin,
		application.HandlerLoginRequest,
	)

	// Grouped by sidebar menu w sidebar call first
	// Tickets

	application.Transport.Get(
		constants.RouteTickets,
		application.HandlerHTMLTicketsTable,
	)

	application.Transport.Post(
		constants.RouteTickets,
		application.HandlerHTMLTicketsTableBody,
	)

	application.Transport.Post(
		constants.RouteTicket,
		application.HandlerAddTicket,
	)

	// Ticket Events

	application.Transport.Get(
		constants.RouteTicketEvents,
		application.HandlerHTMLTicketEventsTable,
	)

	application.Transport.Post(
		constants.RouteTicketEvents,
		application.HandlerHTMLTicketEventsTableBody,
	)

	// Ticket

	application.Transport.Get(
		constants.RouteGetTicket,
		application.HandlerHTMLTicketEventsWContentTable,
	)

	application.Transport.Post(
		constants.RouteGetTicket,
		application.HandlerHTMLTicketTableBody,
	)

	application.Transport.Get(
		constants.RouteGetTicket+"/:id",
		application.HandlerHTMLTicketIDFull,
	)

	// Unused

	application.Transport.Post(
		constants.RouteTicketEvent,
		application.HandlerAddEvent,
	)
}
