package srender

import (
	"fmt"
	"strings"

	g "github.com/maragudk/gomponents"
)

type ElementInput struct {
	CSSClass    string
	ElementName string
	TypeInput   string
}

func (el ElementInput) Raw() g.Node {
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

type InputElements []*ElementInput

func (els InputElements) AsHTML(buttons ...g.Node) []g.Node {
	result := make([]g.Node, len(els), len(els))

	for ix, el := range els {
		result[ix] = el.Raw()
	}

	return append(
		result,
		buttons...,
	)
}
