package srender

import (
	g "github.com/maragudk/gomponents"
	html "github.com/maragudk/gomponents/html"
)

type ParamsBody struct {
	Header      []g.Node
	SidebarMenu *MenuSidebar
	Main        []g.Node
}

func (s *Service) Body(params *ParamsBody) []g.Node {
	return []g.Node{
		html.Div(
			append(
				params.Header,
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

			g.If(
				params.SidebarMenu != nil,
				params.SidebarMenu.Render(),
			),
		),

		html.Div(
			append(
				params.Main,
				g.Attr(
					"class",
					"main",
				),
			)...,
		),
	}
}
