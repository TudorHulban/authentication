package srender

import (
	"github.com/TudorHulban/authentication/app/constants"
	g "github.com/maragudk/gomponents"
	html "github.com/maragudk/gomponents/html"
)

type ParamsNewFormSearchTicket struct {
	TextForm       string
	IDEnclosingDiv string

	ActionButtonSearch string
	ClassButtonSearch  string
	LabelButtonSearch  string
	TargetsSwapSearch  []string
}

func (s *Service) NewFormSearchTicket(params *ParamsNewFormSearchTicket) g.Node {
	return newFormGeneric(
		&paramsNewFormGeneric{
			TextForm: params.TextForm,

			IDForm:         constants.IDSearchItems,
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

					IsRequired: true,
				},
			},
		},
	)
}
