package pages

import (
	g "github.com/maragudk/gomponents"
	html "github.com/maragudk/gomponents/html"
)

type ParamsModalCreateTicket struct {
	URLAddTicket string
	Label        string
}

func ModalCreateTicket(params *ParamsModalCreateTicket) g.Node {
	return html.Div(
		html.Div(
			g.Attr(
				"id",
				"modal-content",
			),
			g.Attr(
				"class",
				"modal",
			),

			html.Form(
				g.Attr(
					"onSubmit",
					"closeModal()",
				),
				g.Attr(
					"hx-post",
					params.URLAddTicket,
				),
				g.Attr(
					"hx-target",
					"#tickets-list",
				),
				g.Attr(
					"hx-target",
					"#tickets-list",
				),

				html.Label(
					g.Attr(
						"for",
						"ticketname",
					),
					g.Text("Ticket Name"),
				),

				html.Input(
					g.Attr(
						"type",
						"text",
					),
					g.Attr(
						"id",
						"name",
					),
					g.Attr(
						"name",
						"ticketname",
					),
				),

				html.Button(
					g.Attr(
						"type",
						"submit",
					),
					g.Text("Submit"),
				),
			),
		),

		html.Div(
			g.Attr(
				"id",
				"htmx-alert",
			),
			g.Attr(
				"hidden",
			),
			g.Attr(
				"class",
				"alert alert-warning sticky-top",
			),
		),
	)
}
