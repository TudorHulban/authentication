package app

import (
	"context"

	"github.com/TudorHulban/authentication/app/constants"
	"github.com/TudorHulban/authentication/domain/ticket"
	"github.com/TudorHulban/authentication/services/srender"
	g "github.com/maragudk/gomponents"
)

func (a *App) formSearchCreateTickets() g.Node {
	return a.serviceRender.NewFormSearchCreateTickets(
		&srender.ParamsNewFormSearchTickets{
			TextForm: "Search / Create Tickets",

			ActionButtonCreate: constants.RouteTicket,
			ActionButtonSearch: constants.RouteTickets,

			LabelButtonSearch: "Search",
			LabelButtonCreate: "Create",

			TargetsSwapCreate: []string{
				constants.IDItemsTableBody,
			},

			IDEnclosingDiv: constants.IDContainerSearchItems,
		},
	)
}

func (a *App) mainWithTickets(ctx context.Context, tickets ticket.Tickets) []g.Node {
	return []g.Node{
		a.formSearchCreateTickets(),

		a.serviceRender.TableItems(
			ctx,

			&srender.ParamsTableItems{
				TableHead: a.serviceRender.TableItemsHeadForTickets(
					constants.IDItemsTableHead,
				),

				TableBody: a.serviceRender.TableItemsBodyForTickets(
					ctx,

					&srender.ParamsRenderTickets{
						Tickets: tickets,

						RouteTicket:     constants.RouteTicket,
						CSSIDTicketBody: constants.IDItemsTableBody,
					},
				),
			},
		),
	}
}
