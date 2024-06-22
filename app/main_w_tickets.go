package app

import (
	"context"

	"github.com/TudorHulban/authentication/app/constants"
	"github.com/TudorHulban/authentication/domain/ticket"
	"github.com/TudorHulban/authentication/services/srender"
	g "github.com/maragudk/gomponents"
)

func (a *App) mainWithTickets(ctx context.Context, tickets ticket.Tickets) []g.Node {
	return []g.Node{
		// 	a.serviceRender.NewFormSearchCreateTickets(
		// 		&srender.ParamsNewFormSearchTickets{
		// 			TextForm: "Search / Create Tickets",

		// 			ActionButtonCreate: constants.RouteTicket,
		// 			ActionButtonSearch: constants.RouteTickets,

		// 			LabelButtonSearch: "Search",
		// 			LabelButtonCreate: "Create",

		// 			IDEnclosingDiv: constants.IDContainerSearchItems,
		// 		},
		// 	),

		// 	a.serviceRender.TableTicketsHead(
		// 		constants.IDItemsTableHead,
		// 	),

		// 	a.serviceRender.RenderTicketsTableBody(
		// 		ctx,

		// 		&srender.ParamsRenderTickets{
		// 			Tickets: tickets,

		// 			RouteTicket:     constants.RouteTicket,
		// 			CSSIDTicketBody: constants.IDItemsTableBody,
		// 		},
		// 	),

		a.serviceRender.NewFormSearchCreateTickets(
			&srender.ParamsNewFormSearchTickets{
				TextForm: "Search / Create Tickets",

				ActionButtonCreate: constants.RouteTicket,
				ActionButtonSearch: constants.RouteTickets,

				LabelButtonSearch: "Search",
				LabelButtonCreate: "Create",

				IDEnclosingDiv: constants.IDContainerSearchItems,
			},
		),

		a.serviceRender.TableItems(
			ctx,

			&srender.ParamsTableItems{
				Tickets:        tickets,
				URLTicket:      constants.RouteTicket,
				CSSIDTableHead: constants.IDItemsTableHead,
			},
		),
	}
}
