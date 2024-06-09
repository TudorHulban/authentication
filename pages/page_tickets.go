package pages

import (
	"strconv"

	"github.com/TudorHulban/authentication/domain/ticket"
	g "github.com/maragudk/gomponents"
	html "github.com/maragudk/gomponents/html"
)

// table tickets

func TableTickets(tickets ticket.Tickets) g.Node {
	if len(tickets) == 0 {
		return nil
	}

	var ix int
	var currentTicket *ticket.Ticket

	tableTicketsRow := func(item *ticket.Ticket) g.Node {
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
		)
	}

	rows := make([]g.Node, 0)

	for ix, currentTicket = range tickets {
		rows = append(rows,
			tableTicketsRow(currentTicket),
		)
	}

	return html.Div(
		html.Table(
			g.Attr(
				"id",
				"tickets-list",
			),

			html.THead(
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
			),

			g.Group(rows),
		),
	)
}
