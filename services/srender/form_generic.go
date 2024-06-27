package srender

import (
	g "github.com/maragudk/gomponents"
	html "github.com/maragudk/gomponents/html"
)

type paramsNewFormGeneric struct {
	TextForm string

	IDForm         string
	IDEnclosingDiv string

	Elements InputElements

	Buttons []g.Node
}

func newFormGeneric(params *paramsNewFormGeneric) g.Node {
	return html.Div(
		g.If(
			len(params.IDEnclosingDiv) > 0,
			g.Attr(
				"id",
				params.IDEnclosingDiv,
			),
		),

		g.If(
			len(params.TextForm) > 0,
			html.H3(
				g.Text(
					params.TextForm,
				),
			),
		),

		html.Form(
			append(
				[]g.Node{
					g.Attr(
						"id",
						params.IDForm,
					),
				},
				params.Elements.AsHTML(params.Buttons...)...,
			)...,
		),
	)
}
