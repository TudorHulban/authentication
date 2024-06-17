package srender

import (
	"fmt"
	"strings"

	"github.com/TudorHulban/authentication/helpers"
	"github.com/gofiber/fiber/v2"
	g "github.com/maragudk/gomponents"
	html "github.com/maragudk/gomponents/html"
)

type ElementForm struct {
	CSSClass    string
	ElementName string
	TypeInput   string
}

func (el ElementForm) Raw() g.Node {
	var result [3]string

	toLowerElementName := strings.ToLower(el.ElementName)

	if len(el.CSSClass) == 0 {
		result[0] = `<div>`
	} else {
		result[0] = fmt.Sprintf(
			`<div class="%s">`,
			el.CSSClass,
		)
	}

	if toLowerElementName == "id" {
		result[1] = fmt.Sprintf(
			`<label for="%s">%s:</label>`,
			toLowerElementName,
			"ID",
		)
	} else {
		result[1] = fmt.Sprintf(
			`<label for="%s">%s:</label>`,
			toLowerElementName,
			strings.ToUpper(toLowerElementName[:1])+toLowerElementName[1:],
		)
	}

	result[2] = fmt.Sprintf(
		`<input type="%s" id="%s" name="%s"></div>`,
		el.TypeInput,
		toLowerElementName,
		toLowerElementName,
	)

	return g.Raw(
		strings.Join(result[:], "\n"),
	)
}

type FormElements []*ElementForm

func (els FormElements) Raw(buttonSubmit g.Node) []g.Node {
	result := make([]g.Node, len(els), len(els))

	for ix, el := range els {
		result[ix] = el.Raw()
	}

	return append(
		result,
		buttonSubmit,
	)
}

type paramsNewFormGeneric struct {
	TextForm string

	IDForm         string
	ActionForm     string
	HTTPMethodForm string
	IDEnclosingDiv string

	IDTarget string

	Elements     FormElements
	ButtonSubmit g.Node
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
				params.Elements.Raw(params.ButtonSubmit)...,
			)...,
		),
	)
}
