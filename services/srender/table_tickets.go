package srender

import (
	"context"
	"fmt"
	"strconv"

	"github.com/TudorHulban/authentication/app/constants"
	"github.com/TudorHulban/authentication/domain/ticket"
	"github.com/TudorHulban/authentication/helpers"
	g "github.com/maragudk/gomponents"
	html "github.com/maragudk/gomponents/html"
)

// table tickets
// create new ticket button

func (s *Service) TableTicketsHead(cssID string) g.Node {
	return html.THead(
		g.Attr(
			"id",
			cssID,
		),

		html.Th(
			g.Text("#"),
		),
		html.Th(
			g.Text("ID"),
		),
		html.Th(
			g.Text("Name"),
		),
		html.Th(
			g.Text("Status"),
		),
		html.Th(
			g.Text("Opened By"),
		),
		html.Th(
			g.Text("Last Update"),
		),
	)
}

type ParamsTableTickets struct {
	Tickets   ticket.Tickets
	URLTicket string

	CSSIDTableHead string
}

func (s *Service) TableTickets(ctx context.Context, params *ParamsTableTickets) g.Node {
	if len(params.Tickets) == 0 {
		return nil
	}

	var ix int
	var currentTicket *ticket.Ticket

	// TODO: delete and use RenderTicketInTableRow
	tableTicketsRow := func(item *ticket.Ticket) g.Node {
		userInfo, errGetUserInfo := s.serviceUser.GetUserInfoByID(ctx, item.OpenedByUserID)
		if errGetUserInfo != nil {
			fmt.Println(errGetUserInfo) //TODO: proper log
		}

		return html.Tr(
			html.Td(
				g.Text(
					strconv.Itoa(ix+1),
				),
			),
			html.Td(
				g.Text(
					item.PrimaryKey.String(),
				),
			),
			html.Td(
				Navigation(
					&ParamsNavigation{
						WhereTo:        params.URLTicket + "/" + item.PrimaryKey.String(),
						LabelToDisplay: item.Name,
					},
				),
			),
			html.Td(
				g.Text(
					item.Status.String(),
				),
			),
			html.Td(
				g.Text(
					helpers.Sanitize(userInfo).Name,
				),
			),
			html.Td(
				g.Text(
					helpers.UnixNanoTo(
						item.UpdatedAt,
					),
				),
			),
		)
	}

	rows := make([]g.Node, 0)

	for ix, currentTicket = range params.Tickets {
		rows = append(rows,
			tableTicketsRow(currentTicket),
		)
	}

	return html.Div(
		html.Table(
			[]g.Node{
				g.Attr(
					"id",
					constants.IDItemsTable,
				),

				g.Attr(
					"class",
					constants.ClassTableItems,
				),

				s.TableTicketsHead(
					params.CSSIDTableHead,
				),

				html.TBody(
					append(
						rows,

						g.Attr(
							"id",
							constants.IDItemsTableBody,
						),
					)...,
				),
			}...,
		),
	)
}
