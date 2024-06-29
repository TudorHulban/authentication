package app

import (
	"context"

	"github.com/TudorHulban/authentication/app/constants"
	"github.com/TudorHulban/authentication/domain/ticket"
	"github.com/TudorHulban/authentication/helpers"
	"github.com/TudorHulban/authentication/services/srender"
	g "github.com/maragudk/gomponents"
)

func (a *App) formSearchTicketCreateEvent(ticket *ticket.Ticket) g.Node {
	return a.serviceRender.NewFormSearchTicketCreateEvent(
		&srender.ParamsNewFormSearchTicketCreateEvent{
			TextForm: "Search Ticket / Create Event",
			Ticket:   ticket,

			ActionButtonSearch:            constants.RouteGetTicket,
			ActionButtonCreateTicketEvent: constants.RouteTicketEvent,

			LabelButtonSearch:            "Search",
			LabelButtonCreateTicketEvent: constants.LabelTicketEventContent,

			TargetsSwapSearch:            swapTargetsBodyTicket,
			TargetsSwapCreateTicketEvent: swapTargetsBodyTicket,

			IDEnclosingDiv:      constants.IDContainerSearchItems,
			IDInputTicketID:     constants.IDSearchItemsInputID,
			IDInputEventContent: constants.IDTicketEventContent,
		},
	)
}

type ParamsHTMLWithTicketEventsWContent struct {
	TicketEvents ticket.Events
	Ticket       *ticket.Ticket
}

func (a *App) HTMLWithTicketEventsWContent(ctx context.Context, params *ParamsHTMLWithTicketEventsWContent) []g.Node {
	return []g.Node{
		a.formSearchTicketCreateEvent(
			helpers.Sanitize(params).Ticket,
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
