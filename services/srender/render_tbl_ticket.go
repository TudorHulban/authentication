package srender

import (
	"context"
	"fmt"

	"github.com/maragudk/gomponents/html"

	"github.com/TudorHulban/authentication/domain/ticket"
	"github.com/TudorHulban/authentication/helpers"
	"github.com/TudorHulban/authentication/services/sticket"
	g "github.com/maragudk/gomponents"
)

type ParamsTicketAsHTML struct {
	Ticket *ticket.Ticket

	RouteGetTicket   string
	TargetsMultiswap string

	Index int
}

func (s *Service) RenderTicketInTableRow(ctx context.Context, params *ParamsTicketAsHTML) g.Node {
	userInfo, errGetUserInfo := s.serviceUser.GetUserInfoByID(
		ctx,
		params.Ticket.OpenedByUserID,
	)
	if errGetUserInfo != nil {
		fmt.Println(errGetUserInfo) //TODO: proper log
	}

	ticketEventType, errGetStatus := s.serviceTicket.GetTicketStatus(
		ctx,
		&sticket.ParamsGetTicketStatus{
			TicketID: params.Ticket.PrimaryKey,
			UserInfo: *userInfo,
		},
	)
	if errGetStatus != nil {
		fmt.Println(errGetStatus)
	}

	return g.Rawf(
		`<tr><td>%d</td><td><a href="#" hx-get="%s/%d" hx-swap="%s">%d</td><td>%s</a><td>%s</td><td>%s</td><td>%s</td></tr>`,

		params.Index,

		params.RouteGetTicket,
		params.Ticket.PrimaryKey,

		params.TargetsMultiswap,

		params.Ticket.PrimaryKey,

		params.Ticket.Name,

		ticketEventType.String(),
		userInfo.Name,
		helpers.UnixNanoTo(
			params.Ticket.UpdatedAt,
		),
	)
}

type ParamsRenderTickets struct {
	Tickets ticket.Tickets

	RouteGetTicket  string
	CSSIDTicketBody string

	TargetsSwapSearch []string
}

func (s *Service) TableItemsBodyForTickets(ctx context.Context, params *ParamsRenderTickets) g.Node {
	result := make([]g.Node, len(params.Tickets), len(params.Tickets))

	targetsMultiswap := NewMultiswap(
		params.TargetsSwapSearch,
	)

	for ix, item := range params.Tickets {
		result[ix] = s.RenderTicketInTableRow(
			ctx,
			&ParamsTicketAsHTML{
				Ticket: item,
				Index:  ix + 1,

				RouteGetTicket:   params.RouteGetTicket,
				TargetsMultiswap: targetsMultiswap,
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
