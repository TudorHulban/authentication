package srender

import (
	"github.com/TudorHulban/authentication/app/constants"
	"github.com/gofiber/fiber/v2"
	g "github.com/maragudk/gomponents"
	html "github.com/maragudk/gomponents/html"
)

type ParamsNewFormSearchTicketEvents struct {
	TextForm       string
	IDEnclosingDiv string

	ActionButtonSearch string
	ClassButtonSearch  string
	LabelButtonSearch  string
	TargetsSwapSearch  []string
}

func (s *Service) NewFormSearchTicketEvents(params *ParamsNewFormSearchTicketEvents) g.Node {
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
						params.ActionButtonSearch,
					),

					g.Attr(
						"hx-swap",
						NewMultiswap(
							params.TargetsSwapSearch,
						),
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
					CSSClassDiv: "form-group",
					CSSIDInput:  constants.IDSearchItemsInputID,

					ElementName: "TicketID",
					TypeInput:   "text",
				},
			},
		},
	)
}
