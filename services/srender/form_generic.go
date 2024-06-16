package srender

import (
	"fmt"
	"strings"

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

	result[0] = fmt.Sprintf(
		`<div class="%s">`,
		el.CSSClass,
	)
	result[1] = fmt.Sprintf(
		`<label for="%s">%s:</label>`,
		toLowerElementName,
		strings.ToUpper(toLowerElementName[:1])+toLowerElementName[1:],
	)
	result[2] = fmt.Sprintf(
		`<input type="%s" id="%s" name="%s">`,
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
	IDForm            string
	ActionForm        string
	HTTPMethodForm    string
	ClassEnclosingDiv string

	Elements     FormElements
	ButtonSubmit g.Node
}

func newFormGeneric(params *paramsNewFormGeneric) g.Node {
	return html.Div(
		g.Attr(
			"class",
			params.ClassEnclosingDiv,
		),

		html.Form(
			append(
				[]g.Node{
					g.Attr(
						"id",
						params.IDForm,
					),

					g.Attr(
						"action",
						params.ActionForm,
					),

					g.Attr(
						"method",
						params.HTTPMethodForm,
					),
				},
				params.Elements.Raw(params.ButtonSubmit)...,
			)...,
		),
	)
}
