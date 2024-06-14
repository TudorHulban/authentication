package srender

import (
	g "github.com/maragudk/gomponents"
	html "github.com/maragudk/gomponents/html"
)

type ParamsBody struct {
	EntriesHeader []g.Node
	SidebarMenu   *MenuSidebar
	EntriesMain   []g.Node
}

func Body(params *ParamsBody) []g.Node {
	return []g.Node{
		html.Div(
			append(
				params.EntriesHeader,
				g.Attr(
					"class",
					"header",
				),
			)...,
		),

		html.Div(
			g.Attr(
				"class",
				"sidebar",
			),

			params.SidebarMenu.Render(),
		),

		html.Div(
			append(
				params.EntriesMain,
				g.Attr(
					"class",
					"main",
				),
			)...,
		),
	}
}
