package app

import (
	"context"

	"github.com/TudorHulban/authentication/app/constants"
	"github.com/TudorHulban/authentication/domain/ticket"
	"github.com/TudorHulban/authentication/services/srender"
	g "github.com/maragudk/gomponents"
)

func (a *App) formSearchCreateTicketEvents() g.Node {
	return a.serviceRender.NewFormSearchTicketEvents(
		&srender.ParamsNewFormSearchTicketEvents{
			TextForm: "Search Ticket Events",

			ActionButtonSearch: constants.RouteTicketEvents,
			LabelButtonSearch:  "Search",

			TargetsSwapSearch: []string{
				constants.IDItemsTableBody,
			},

			IDEnclosingDiv: constants.IDContainerSearchItems,
		},
	)
}

func (a *App) HTMLWithTicketEvents(ctx context.Context, ticketEvents ticket.Events) []g.Node {
	return []g.Node{
		a.formSearchCreateTicketEvents(),

		a.serviceRender.TableItemsHeadForTicketEvents(
			constants.IDItemsTableHead,
		),

		a.serviceRender.TableItemsBodyForTicketEvents(
			ctx,
			&srender.ParamsRenderTicketEvents{
				Events: ticketEvents,

				RouteGetTicket:        constants.RouteGetTicket,
				CSSIDTicketEventsBody: constants.IDItemsTableBody,

				TargetsSwapSearch: []string{
					constants.IDContainerSearchItems,
					constants.IDItemsTableHead,
					constants.IDItemsTableBody,
				},
			},
		),
	}
}
