package srender

import (
	"github.com/TudorHulban/authentication/app/constants"
	"github.com/TudorHulban/authentication/helpers"
	g "github.com/maragudk/gomponents"
	html "github.com/maragudk/gomponents/html"
)

type ParamsNewFormSearchCreateTickets struct {
	TextForm string

	ActionButtonCreate string
	ActionButtonSearch string

	IDEnclosingDiv    string
	IDInputTicketName string

	ClassButtonSearch string
	LabelButtonSearch string
	TargetsSwapSearch []string

	ClassButtonCreate string
	LabelButtonCreate string
	TargetsSwapCreate []string
}

func (s *Service) NewFormSearchCreateTickets(params *ParamsNewFormSearchCreateTickets) g.Node {
	return newFormGeneric(
		&paramsNewFormGeneric{
			TextForm: params.TextForm,

			IDForm:         constants.IDSearchItems,
			IDEnclosingDiv: params.IDEnclosingDiv,

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

					g.Attr(
						"hx-require",
						helpers.SanitizeCSSId(
							params.IDInputTicketName,
						),
					),

					g.Attr(
						"hx-swap",
						NewMultiswap(
							params.TargetsSwapCreate,
						),
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

					g.Attr(
						"hx-swap",
						NewMultiswap(
							params.TargetsSwapCreate,
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

					ElementName: "ID",
					TypeInput:   "text",
				},
				{
					CSSClassDiv: "form-group",
					CSSIDInput:  constants.IDSearchItemsInputStatus,

					ElementName: "Status",
					TypeInput:   "text",
				},
				{
					CSSClassDiv: "form-group",
					CSSIDInput:  constants.IDSearchItemsInputName,

					ElementName: "Name",
					TypeInput:   "text",
				},
			},
		},
	)
}
