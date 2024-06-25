package app

import (
	"context"

	"github.com/TudorHulban/authentication/app/constants"
	"github.com/TudorHulban/authentication/domain/ticket"
	"github.com/TudorHulban/authentication/helpers"
	"github.com/TudorHulban/authentication/services/srender"
	g "github.com/maragudk/gomponents"
)

func (a *App) formSearchTicketCreateEvent(ticketID helpers.PrimaryKey) g.Node {
	return a.serviceRender.NewFormSearchTicketCreateEvent(
		&srender.ParamsNewFormSearchTicketCreateEvent{
			TextForm:  "Search Ticket / Create Event",
			TextInput: ticketID.String(),

			ActionButtonSearch:            constants.RouteGetTicket,
			ActionButtonCreateTicketEvent: constants.RouteTicketEvent,

			LabelButtonSearch:            "Search",
			LabelButtonCreateTicketEvent: "Create event",

			TargetsSwapSearch: []string{
				constants.IDItemsTableBody,
			},
			TargetsSwapCreateTicketEvent: []string{
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
		a.formSearchTicketCreateEvent(
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
