package srender

import (
	g "github.com/maragudk/gomponents"
	html "github.com/maragudk/gomponents/html"
)

func Body(menu *MenuSidebar) g.Node {
	return html.Div(
		g.Attr(
			"class",
			"container",
		),

		html.Div(
			g.Attr(
				"class",
				"sidebar",
			),

			menu.Render(),
		),

		html.Div(
			g.Attr(
				"class",
				"article",
			),
		),
	)
}
