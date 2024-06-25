package srender

import (
	"github.com/TudorHulban/authentication/app/constants"
	"github.com/TudorHulban/authentication/helpers"
	g "github.com/maragudk/gomponents"
	html "github.com/maragudk/gomponents/html"
)

type ParamsNewFormSearchTicketCreateEvent struct {
	TextForm  string
	TextInput string

	IDEnclosingDiv      string
	IDInputTicketID     string
	IDInputEventContent string

	ActionButtonSearch string
	ClassButtonSearch  string
	LabelButtonSearch  string
	TargetsSwapSearch  []string

	ActionButtonCreateTicketEvent string
	ClassButtonCreateTicketEvent  string
	LabelButtonCreateTicketEvent  string
	TargetsSwapCreateTicketEvent  []string
}

func (s *Service) NewFormSearchTicketCreateEvent(params *ParamsNewFormSearchTicketCreateEvent) g.Node {
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

					g.Attr(
						"hx-require",
						helpers.SanitizeCSSId(
							params.IDInputTicketID,
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

				html.Button(
					g.Attr(
						"type",
						"submit",
					),

					g.Attr(
						"hx-post",
						params.ActionButtonCreateTicketEvent,
					),

					g.Attr(
						"hx-swap",
						NewMultiswap(
							params.TargetsSwapCreateTicketEvent,
						),
					),

					g.Attr(
						"hx-require",
						NewMultiswap(
							[]string{
								params.IDInputTicketID,
								params.IDInputEventContent,
							},
						),
					),

					g.Text(
						params.LabelButtonCreateTicketEvent,
					),
				),
			},

			Elements: []*ElementInput{
				{
					CSSClassDiv: "form-group",
					CSSIDInput:  constants.IDSearchItemsInputID,

					ElementName: "TicketID",
					TypeInput:   "text",

					TextInput: params.TextInput,
				},
				{
					CSSClassDiv: "form-group",
					CSSIDInput:  constants.IDSearchItemsInputID,

					ElementName: "Event type",
					TypeInput:   "text",
				},
				{
					CSSClassDiv: "form-group",
					CSSIDInput:  constants.IDTicketEventContent,

					ElementName: "Event content",
					TypeInput:   "text",

					TextInput:  params.TextInput,
					IsTextArea: true,
				},
			},
		},
	)
}
