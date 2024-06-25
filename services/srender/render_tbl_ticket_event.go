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

	RouteGetTicket   string
	TargetsMultiswap string

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
		`<tr><td>%d</td><td><a href="#" hx-get="%s/%d" hx-swap="%s">%d</td><td>%d<td>%d</td><td>%s</td><td>%s</td></tr>`,

		params.Index,

		params.RouteGetTicket,
		params.TicketEvent.TicketPK,

		params.TargetsMultiswap,

		params.TicketEvent.TicketPK,
		params.TicketEvent.PrimaryKey,

		params.TicketEvent.TicketEventType,

		userInfo.Name,
		helpers.UnixNanoTo(
			params.TicketEvent.TimestampOfAdd,
		),
	)
}

func (s *Service) RenderTicketEventWContentInTableRow(ctx context.Context, params *ParamsTicketEventAsHTML) g.Node {
	userInfo, errGetUserInfo := s.serviceUser.GetUserInfoByID(
		ctx,
		params.TicketEvent.OpenedByUserID,
	)
	if errGetUserInfo != nil {
		fmt.Println(errGetUserInfo) //TODO: proper log
	}

	return g.Rawf(
		`<tr><td>%d</td><td>%d<td>%d</td><td>%s</td><td>%s</td><td>%s</td></tr>`,

		params.Index,
		params.TicketEvent.PrimaryKey,

		params.TicketEvent.TicketEventType,

		userInfo.Name,
		helpers.UnixNanoTo(
			params.TicketEvent.TimestampOfAdd,
		),
		params.TicketEvent.Content,
	)
}

type ParamsRenderTicketEvents struct {
	Events ticket.Events

	RouteGetTicket        string
	CSSIDTicketEventsBody string

	TargetsSwapSearch []string
}

func (s *Service) TableItemsBodyForTicketEvents(ctx context.Context, params *ParamsRenderTicketEvents) g.Node {
	return s.tableBodyForTicketEvents(
		ctx,
		&paramsTableBodyForTicketEvents{
			RenderInfo: params,
			Renderer:   s.RenderTicketEventInTableRow,
		},
	)
}

func (s *Service) TableItemsBodyForTicketEventsWContent(ctx context.Context, params *ParamsRenderTicketEvents) g.Node {
	return s.tableBodyForTicketEvents(
		ctx,
		&paramsTableBodyForTicketEvents{
			RenderInfo: params,
			Renderer:   s.RenderTicketEventWContentInTableRow,
		},
	)
}

type paramsTableBodyForTicketEvents struct {
	RenderInfo *ParamsRenderTicketEvents
	Renderer   func(ctx context.Context, params *ParamsTicketEventAsHTML) g.Node
}

func (s *Service) tableBodyForTicketEvents(ctx context.Context, params *paramsTableBodyForTicketEvents) g.Node {
	result := make([]g.Node, len(params.RenderInfo.Events), len(params.RenderInfo.Events))

	targetsMultiswap := NewMultiswap(
		params.RenderInfo.TargetsSwapSearch,
	)

	for ix, item := range params.RenderInfo.Events {
		result[ix] = params.Renderer(
			ctx,
			&ParamsTicketEventAsHTML{
				TicketEvent: item,
				Index:       ix + 1,

				RouteGetTicket:   params.RenderInfo.RouteGetTicket,
				TargetsMultiswap: targetsMultiswap,
			},
		)
	}

	return html.TBody(
		append(
			result,

			g.If(
				len(params.RenderInfo.CSSIDTicketEventsBody) > 0,
				g.Attr(
					"id",
					params.RenderInfo.CSSIDTicketEventsBody,
				),
			),
		)...,
	)
}
