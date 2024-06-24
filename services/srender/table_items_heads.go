package srender

import (
	g "github.com/maragudk/gomponents"
	html "github.com/maragudk/gomponents/html"
)

func (s *Service) TableItemsHeadForTickets(cssID string) g.Node {
	return html.THead(
		html.Tr(
			g.Attr(
				"id",
				cssID,
			),

			html.Th(
				g.Text("#"),
			),
			html.Th(
				g.Text("Ticket ID"),
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
	)
}

func (s *Service) TableItemsHeadForTicketEvents(cssID string) g.Node {
	return html.THead(
		html.Tr(
			g.Attr(
				"id",
				cssID,
			),

			html.Th(
				g.Text("#"),
			),
			html.Th(
				g.Text("Ticket ID"),
			),
			html.Th(
				g.Text("Event ID"),
			),
			html.Th(
				g.Text("Event Type"),
			),
			html.Th(
				g.Text("Opened By"),
			),
			html.Th(
				g.Text("Opened time"),
			),
		),
	)
}
