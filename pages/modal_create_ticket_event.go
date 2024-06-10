package pages

import (
	g "github.com/maragudk/gomponents"
	html "github.com/maragudk/gomponents/html"
)

type ParamsModalCreateTicketEvent struct {
	URLAddTicketEvent string
}

func ModalCreateTicketEvent(params *ParamsModalCreateTicketEvent) g.Node {
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
					params.URLAddTicketEvent,
				),
				g.Attr(
					"hx-target",
					"#events-list",
				),

				html.Label(
					g.Attr(
						"for",
						"eventname",
					),
					g.Text("Event Name"),
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
						"eventname",
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
