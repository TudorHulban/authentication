package app

import (
	"context"

	"github.com/TudorHulban/authentication/app/constants"
	"github.com/TudorHulban/authentication/domain/ticket"
	"github.com/TudorHulban/authentication/services/srender"
	g "github.com/maragudk/gomponents"
)

func (a *App) formSearchTicket() g.Node {
	return a.serviceRender.NewFormSearchTicket(
		&srender.ParamsNewFormSearchTicket{
			TextForm: "Search Ticket",

			ActionButtonSearch: constants.RouteGetTicket,
			LabelButtonSearch:  "Search",

			TargetsSwapSearch: []string{
				constants.IDItemsTableBody,
			},

			IDEnclosingDiv:  constants.IDContainerSearchItems,
			IDInputTicketID: constants.IDSearchItemsInputID,
		},
	)
}

func (a *App) HTMLWithTicketEventsWContent(ctx context.Context, ticketEvents ticket.Events) []g.Node {
	return []g.Node{
		a.formSearchTicket(),

		a.serviceRender.TableItemsHeadForTicket(
			constants.IDItemsTableHead,
		),

		a.serviceRender.TableItemsBodyForTicketEventsWContent(
			ctx,
			&srender.ParamsRenderTicketEvents{
				Events: ticketEvents,

				CSSIDTicketEventsBody: constants.IDItemsTableBody,
			},
		),
	}
}
