package app

import (
	"context"

	"github.com/TudorHulban/authentication/app/constants"
	"github.com/TudorHulban/authentication/domain/ticket"
	"github.com/TudorHulban/authentication/helpers"
	"github.com/TudorHulban/authentication/services/srender"
	g "github.com/maragudk/gomponents"
)

func (a *App) formSearchTicket(ticketID helpers.PrimaryKey) g.Node {
	return a.serviceRender.NewFormSearchTicket(
		&srender.ParamsNewFormSearchTicket{
			TextForm:  "Search Ticket",
			TextInput: ticketID.String(),

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

type ParamsHTMLWithTicketEventsWContent struct {
	TicketEvents ticket.Events
	TicketID     helpers.PrimaryKey
}

func (a *App) HTMLWithTicketEventsWContent(ctx context.Context, params *ParamsHTMLWithTicketEventsWContent) []g.Node {
	return []g.Node{
		a.formSearchTicket(
			helpers.Sanitize(params).TicketID,
		),

		a.serviceRender.TableItemsHeadForTicket(
			constants.IDItemsTableHead,
		),

		a.serviceRender.TableItemsBodyForTicketEventsWContent(
			ctx,
			&srender.ParamsRenderTicketEvents{
				Events: helpers.Sanitize(params).TicketEvents,

				CSSIDTicketEventsBody: constants.IDItemsTableBody,
			},
		),
	}
}
