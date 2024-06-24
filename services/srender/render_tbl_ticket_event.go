package srender

import (
	"context"
	"fmt"

	"github.com/TudorHulban/authentication/domain/ticket"
	"github.com/TudorHulban/authentication/helpers"
	g "github.com/maragudk/gomponents"
	html "github.com/maragudk/gomponents/html"
)

type ParamsTicketEventAsHTML struct {
	TicketEvent *ticket.Event

	Index int
}

func (s *Service) RenderTicketEventInTableRow(ctx context.Context, params *ParamsTicketEventAsHTML) g.Node {
	userInfo, errGetUserInfo := s.serviceUser.GetUserInfoByID(
		ctx,
		params.TicketEvent.OpenedByUserID,
	)
	if errGetUserInfo != nil {
		fmt.Println(errGetUserInfo) //TODO: proper log
	}

	return g.Rawf(
		`<tr><td>%d</td><td>%d</td><td>%d<td>%d</td><td>%s</td><td>%s</td></tr>`,

		params.Index,
		params.TicketEvent.TicketPK,
		params.TicketEvent.PrimaryKey,

		params.TicketEvent.TicketEventType,

		userInfo.Name,
		helpers.UnixNanoTo(
			params.TicketEvent.TimestampOfAdd,
		),
	)
}

type ParamsRenderTicketEvents struct {
	Events ticket.Events

	CSSIDTicketEventsBody string
}

func (s *Service) TableItemsBodyForTicketEvents(ctx context.Context, params *ParamsRenderTicketEvents) g.Node {
	result := make([]g.Node, len(params.Events), len(params.Events))

	for ix, item := range params.Events {
		result[ix] = s.RenderTicketEventInTableRow(
			ctx,
			&ParamsTicketEventAsHTML{
				TicketEvent: item,
				Index:       ix + 1,
			},
		)
	}

	return html.TBody(
		append(
			result,

			g.If(
				len(params.CSSIDTicketEventsBody) > 0,
				g.Attr(
					"id",
					params.CSSIDTicketEventsBody,
				),
			),
		)...,
	)
}
