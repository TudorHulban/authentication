package srender

import (
	"github.com/TudorHulban/authentication/app/constants"
	"github.com/gofiber/fiber/v2"
	g "github.com/maragudk/gomponents"
	html "github.com/maragudk/gomponents/html"
)

type ParamsNewFormSearchTickets struct {
	TextForm       string
	ActionForm     string
	IDEnclosingDiv string

	ClassButtonSubmit string
	LabelButtonSubmit string

	ClassButtonSearch string
	LabelButtonSearch string
}

func NewFormSearchTickets(params *ParamsNewFormSearchTickets) g.Node {
	return newFormGeneric(
		&paramsNewFormGeneric{
			TextForm: params.TextForm,

			IDForm:         constants.IDSearchItems,
			ActionForm:     params.ActionForm,
			HTTPMethodForm: fiber.MethodGet,
			IDEnclosingDiv: params.IDEnclosingDiv,

			IDTarget: constants.IDItemsTableBody,

			Buttons: []g.Node{
				html.Button(
					g.Attr(
						"type",
						"submit",
					),

					g.If(
						len(params.ClassButtonSubmit) > 0,
						g.Attr(
							"class",
							params.ClassButtonSubmit,
						),
					),

					g.Text(
						params.LabelButtonSubmit,
					),
				),

				html.Button(
					g.Attr(
						"type",
						"submit",
					),

					g.If(
						len(params.ClassButtonSearch) > 0,
						g.Attr(
							"class",
							params.ClassButtonSearch,
						),
					),

					g.Text(
						params.LabelButtonSearch,
					),
				),
			},

			Elements: []*ElementInput{
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
