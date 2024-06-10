package pages

import (
	"strconv"

	"github.com/TudorHulban/authentication/domain/ticket"
	"github.com/TudorHulban/authentication/helpers"
	g "github.com/maragudk/gomponents"
	html "github.com/maragudk/gomponents/html"
)

// table tickets
// create new ticket button

type ParamsTableTickets struct {
	Tickets   ticket.Tickets
	URLTicket string
}

func TableTickets(params *ParamsTableTickets) g.Node {
	if len(params.Tickets) == 0 {
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
					item.OpenedByUserID.String(),
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
