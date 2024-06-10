package pages

import (
	g "github.com/maragudk/gomponents"
	html "github.com/maragudk/gomponents/html"
)

func ButtonCreateTicket(label string) g.Node {
	return html.Div(
		html.Button(
			g.Attr(
				"hx-target",
				"#modal-content",
			),
			g.Attr(
				"onclick",
				"openModal()",
			),
			g.Text(
				label,
			),
		),
	)
}
