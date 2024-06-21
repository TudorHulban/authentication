package srender

import (
	g "github.com/maragudk/gomponents"
	html "github.com/maragudk/gomponents/html"
)

// table ticket events
// create new ticket event button

func (s *Service) TableTicketEventsHead(cssID string) g.Node {
	return html.THead(
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
	)
}
