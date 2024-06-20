package srender

import (
	"context"
	"fmt"

	"github.com/maragudk/gomponents/html"

	"github.com/TudorHulban/authentication/domain/ticket"
	"github.com/TudorHulban/authentication/helpers"
	g "github.com/maragudk/gomponents"
)

type ParamsTicketAsHTML struct {
	Ticket *ticket.Ticket

	RouteTicket string
	Index       int
}

func (s *Service) RenderTicketInTableRow(ctx context.Context, params *ParamsTicketAsHTML) g.Node {
	userInfo, errGetUserInfo := s.serviceUser.GetUserInfoByID(
		ctx,
		params.Ticket.OpenedByUserID,
	)
	if errGetUserInfo != nil {
		fmt.Println(errGetUserInfo) //TODO: proper log
	}

	return g.Rawf(
		`<tr><td>%d</td><td>%d</td><td><a href="%s/%d">%s</a><td>%s</td><td>%s</td><td>%s</td>`,

		params.Index,
		params.Ticket.PrimaryKey,

		params.RouteTicket,
		params.Ticket.PrimaryKey,
		params.Ticket.Name,

		params.Ticket.Status,
		userInfo.Name,
		helpers.UnixNanoTo(
			params.Ticket.UpdatedAt,
		),
	)
}

type ParamsRenderTickets struct {
	Tickets ticket.Tickets

	RouteTicket     string
	CSSIDTicketBody string
}

func (s *Service) RenderTicketsTableBody(ctx context.Context, params *ParamsRenderTickets) g.Node {
	result := make([]g.Node, len(params.Tickets), len(params.Tickets))

	for ix, item := range params.Tickets {
		result[ix] = s.RenderTicketInTableRow(
			ctx,
			&ParamsTicketAsHTML{
				Ticket: item,
				Index:  ix + 1,
			},
		)
	}

	return html.TBody(
		append(
			result,

			g.If(
				len(params.CSSIDTicketBody) > 0,
				g.Attr(
					"id",
					params.CSSIDTicketBody,
				),
			),
		)...,
	)
}
