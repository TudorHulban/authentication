package pages

import (
	"strconv"

	"github.com/TudorHulban/authentication/domain/ticket"
	"github.com/TudorHulban/authentication/helpers"
	g "github.com/maragudk/gomponents"
	html "github.com/maragudk/gomponents/html"
)

// create new event button
// table ticket events

type ParamsTableTicketEvents struct {
	TicketEvents ticket.Events
}

func TableTicketEvents(params *ParamsTableTicketEvents) g.Node {
	var isHidden g.Node

	if len(params.TicketEvents) == 0 {
		isHidden = hidden
	}

	var ix int
	var currentTicketEvent *ticket.Event

	tableTicketEventsRow := func(item *ticket.Event) g.Node {
		return html.Tr(
			html.Td(
				g.Text(
					strconv.Itoa(ix+1),
				),
			),
			html.Td(
				g.Text(
					helpers.UnixNanoTo(item.TimestampOfAdd),
				),
			),
			html.Td(
				g.Text(
					item.Content,
				),
			),
		)
	}

	rows := make([]g.Node, 0)

	for ix, currentTicketEvent = range params.TicketEvents {
		rows = append(rows,
			tableTicketEventsRow(currentTicketEvent),
		)
	}

	return html.Div(
		html.Table(
			g.Attr(
				"id",
				"events-list",
			),

			isHidden,

			html.THead(
				html.Th(
					g.Text("#"),
				),
				html.Th(
					g.Text("Timestamp"),
				),
				html.Th(
					g.Text("Content"),
				),
			),

			g.Group(rows),
		),
	)
}
