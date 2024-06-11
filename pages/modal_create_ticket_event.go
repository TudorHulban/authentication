package pages

import (
	"github.com/TudorHulban/authentication/helpers"
	g "github.com/maragudk/gomponents"
	html "github.com/maragudk/gomponents/html"
)

type ParamsModalCreateTicketEvent struct {
	URLAddTicketEvent string
	TicketID          helpers.PrimaryKey
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
				// g.Attr(
				// 	"hx-target",
				// 	"#events-list",
				// ),
				g.Attr(
					"hx-swap",
					"none",
				),

				html.Label(
					g.Attr(
						"for",
						"eventcontent",
					),
					g.Text("Event Content"),
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
						"eventcontent",
					),
				),

				html.Input(
					g.Attr(
						"type",
						"hidden",
					),
					g.Attr(
						"name",
						"ticketid",
					),
					g.Attr("value", params.TicketID.String()),
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
