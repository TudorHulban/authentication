package srender

import (
	"github.com/TudorHulban/authentication/app/constants"
	"github.com/TudorHulban/authentication/helpers"
	g "github.com/maragudk/gomponents"
	"github.com/maragudk/gomponents/html"
)

type ObjSearch struct {
	TextTitle string

	IDEnclosingDiv string

	InputElements InputElements
	Buttons       []g.Node
}

func (obj ObjSearch) AsHTML() g.Node {
	return html.Div(
		append(
			[]g.Node{
				g.If(
					len(obj.IDEnclosingDiv) > 0,
					g.Attr(
						"id",
						obj.IDEnclosingDiv,
					),
				),

				html.H3(
					g.Text(
						obj.TextTitle,
					),
				),
			},

			append(
				obj.InputElements.AsHTML(),
				obj.Buttons...,
			)...,
		)...,
	)
}

type ParamsNewSearchCreateTickets struct {
	IDEnclosingDiv string

	IDTargetHTMX      string
	ClassButtonSearch string
	LabelButtonSearch string

	ClassButtonCreate string
	LabelButtonCreate string
}

func NewSearchCreateTickets(params *ParamsNewSearchCreateTickets) g.Node {
	return ObjSearch{
		IDEnclosingDiv: params.IDEnclosingDiv,

		TextTitle: "Search / Create Tickets",

		InputElements: []*ElementInput{
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

		Buttons: []g.Node{
			html.Button(
				g.Attr(
					"hx-get",
					constants.RouteTickets,
				),

				g.Attr(
					"hx-swap",
					"outerHTML",
				),

				g.If(
					len(params.IDTargetHTMX) > 0,

					g.Attr(
						"hx-target",
						helpers.SanitizeCSSId(
							params.IDTargetHTMX,
						),
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
					"hx-post",
					constants.RouteTickets,
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
		},
	}.
		AsHTML()
}
