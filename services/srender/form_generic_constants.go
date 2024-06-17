package srender

import (
	"github.com/gofiber/fiber/v2"
	g "github.com/maragudk/gomponents"
	html "github.com/maragudk/gomponents/html"
)

type ParamsNewFormSearchTickets struct {
	TextForm          string
	ActionForm        string
	ClassEnclosingDiv string
	ClassButtonSubmit string
	LabelButtonSubmit string
}

func NewFormSearchTickets(params *ParamsNewFormSearchTickets) g.Node {
	return newFormGeneric(
		&paramsNewFormGeneric{
			TextForm: params.TextForm,

			IDForm:            "searchForm",
			ActionForm:        params.ActionForm,
			HTTPMethodForm:    fiber.MethodGet,
			ClassEnclosingDiv: params.ClassEnclosingDiv,

			ButtonSubmit: html.Div(
				g.Attr(
					"class",
					params.ClassButtonSubmit,
				),

				html.Button(
					g.Attr(
						"type",
						"submit",
					),

					g.Text(
						params.LabelButtonSubmit,
					),
				),
			),

			Elements: []*ElementForm{
				{
					CSSClass:    "form-group",
					ElementName: "ID",
					TypeInput:   "text",
				},
				{
					CSSClass:    "form-group",
					ElementName: "Status",
					TypeInput:   "text",
				},
			},
		},
	)
}
