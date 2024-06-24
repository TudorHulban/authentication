package srender

import (
	"context"

	"github.com/TudorHulban/authentication/app/constants"
	"github.com/TudorHulban/authentication/domain/ticket"
	g "github.com/maragudk/gomponents"
	html "github.com/maragudk/gomponents/html"
)

type ParamsTableItems struct {
	TableHead g.Node
	TableBody g.Node
}

func (s *Service) TableItems(ctx context.Context, params *ParamsTableItems) g.Node {
	return html.Div(
		html.Table(
			[]g.Node{
				g.Attr(
					"class",
					constants.ClassTableItems,
				),

				params.TableHead,
				params.TableBody,
			}...,
		),
	)
}

type ParamsHTMLTableItemsForTickets struct {
	IDItemsTableHead string
	RouteTickets     string
	CSSIDTicketBody  string

	Tickets ticket.Tickets
}

func (s *Service) HTMLTableItemsForTickets(ctx context.Context, params *ParamsHTMLTableItemsForTickets) []g.Node {
	return []g.Node{
		s.TableItemsHeadForTickets(
			constants.IDItemsTableHead,
		),

		s.TableItemsBodyForTickets(
			ctx,
			&ParamsRenderTickets{
				Tickets: params.Tickets,

				RouteTicket:     params.RouteTickets,
				CSSIDTicketBody: params.CSSIDTicketBody,
			},
		),
	}
}
