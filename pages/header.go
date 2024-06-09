package pages

import (
	g "github.com/maragudk/gomponents"
	html "github.com/maragudk/gomponents/html"
)

func Header() g.Node {
	return html.Div(
		html.H1(
			g.Text(
				"this is header",
			),
		),
	)
}