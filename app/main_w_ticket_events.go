package app

import (
	"github.com/TudorHulban/authentication/app/constants"
	"github.com/TudorHulban/authentication/services/srender"
	g "github.com/maragudk/gomponents"
)

func (a *App) formSearchCreateTicketEvents() g.Node {
	return a.serviceRender.NewFormSearchCreateTicketEvents(
		&srender.ParamsNewFormSearchTicketEvents{
			TextForm: "Search / Create Ticket Events",

			ActionButtonCreate: constants.RouteTicketEvent,
			ActionButtonSearch: constants.RouteTicketEvents,

			LabelButtonSearch: "Search",
			LabelButtonCreate: "Create",

			IDEnclosingDiv: constants.IDContainerSearchItems,
		},
	)
}
