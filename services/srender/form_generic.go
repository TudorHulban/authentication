package srender

import (
	"github.com/TudorHulban/authentication/helpers"
	"github.com/gofiber/fiber/v2"
	g "github.com/maragudk/gomponents"
	html "github.com/maragudk/gomponents/html"
)

type paramsNewFormGeneric struct {
	TextForm string

	IDForm         string
	ActionForm     string
	HTTPMethodForm string
	IDEnclosingDiv string

	IDTarget string

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

					g.If(
						params.HTTPMethodForm == fiber.MethodGet,

						g.Attr(
							"hx-get",
							params.ActionForm,
						),
					),

					g.If(
						params.HTTPMethodForm == fiber.MethodPost,

						g.Attr(
							"hx-post",
							params.ActionForm,
						),
					),

					g.If(
						len(params.IDTarget) > 0,

						g.Attr(
							"hx-target",
							helpers.SanitizeCSSId(
								params.IDTarget,
							),
						),
					),

					g.Attr(
						"hx-swap",
						"outerHTML",
					),
				},
				params.Elements.AsHTML(params.Buttons...)...,
			)...,
		),
	)
}
