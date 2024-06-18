package srender

import (
	"github.com/TudorHulban/authentication/app/constants"
	"github.com/gofiber/fiber/v2"
	g "github.com/maragudk/gomponents"
	html "github.com/maragudk/gomponents/html"
)

type ParamsNewFormSearchTickets struct {
	TextForm string

	ActionButtonCreate string
	ActionButtonSearch string

	IDEnclosingDiv string

	ClassButtonSearch string
	LabelButtonSearch string

	ClassButtonCreate string
	LabelButtonCreate string
}

func NewFormSearchTickets(params *ParamsNewFormSearchTickets) g.Node {
	return newFormGeneric(
		&paramsNewFormGeneric{
			TextForm: params.TextForm,

			IDForm:         constants.IDSearchItems,
			HTTPMethodForm: fiber.MethodGet,
			IDEnclosingDiv: params.IDEnclosingDiv,

			IDTarget: constants.IDItemsTableBody,

			Buttons: []g.Node{
				html.Button(
					g.Attr(
						"type",
						"submit",
					),

					g.Attr(
						"hx-post",
						params.ActionButtonCreate,
					),

					g.If(
						len(params.ClassButtonCreate) > 0,
						g.Attr(
							"class",
							params.ClassButtonCreate,
						),
					),

					g.Text(
						params.LabelButtonCreate,
					),
				),

				html.Button(
					g.Attr(
						"type",
						"submit",
					),

					g.Attr(
						"hx-post",
						params.ActionButtonSearch,
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
				{
					CSSClass:    "form-group",
					ElementName: "Name",
					TypeInput:   "text",
				},
			},
		},
	)
}
