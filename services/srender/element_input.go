package srender

import (
	"fmt"
	"strings"

	"github.com/TudorHulban/authentication/helpers"
	g "github.com/maragudk/gomponents"
)

type ElementInput struct {
	CSSClassDiv string
	CSSIDInput  string
	ElementName string
	TypeInput   string
	TextInput   string

	SelectValues []string

	IsTextArea bool
}

func option(value string) string {
	return fmt.Sprintf(
		`<option value="%s">%s</option>`,
		strings.ToLower(value),
		value,
	)
}

func (el ElementInput) Raw() g.Node {
	var result [3]string

	toLowerElementName := strings.ToLower(el.ElementName)

	if len(el.CSSClassDiv) == 0 {
		result[0] = `<div>`
	} else {
		result[0] = fmt.Sprintf(
			`<div class="%s">`,
			el.CSSClassDiv,
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
			el.ElementName,
		)
	}

	if el.IsTextArea {
		result[2] = fmt.Sprintf(
			`<textarea id="%s" name="%s"></textarea></div>`,
			el.CSSIDInput,
			toLowerElementName,
		)

		return g.Raw(
			strings.Join(result[:], "\n"),
		)
	}

	if len(el.SelectValues) > 0 {
		result[2] = fmt.Sprintf(
			`<select id="%s" name="%s">%s</select></div>`,
			el.CSSIDInput,
			toLowerElementName,

			strings.Join(
				helpers.ForEach(el.SelectValues, option),
				"\n",
			),
		)

		return g.Raw(
			strings.Join(result[:], "\n"),
		)
	}

	if len(el.TextInput) > 0 && el.TextInput != "0" {
		result[2] = fmt.Sprintf(
			`<input type="%s" id="%s" name="%s" value="%s"></div>`,
			el.TypeInput,
			el.CSSIDInput,
			toLowerElementName,
			el.TextInput,
		)
	} else {
		result[2] = fmt.Sprintf(
			`<input type="%s" id="%s" name="%s"></div>`,
			el.TypeInput,
			el.CSSIDInput,
			toLowerElementName,
		)
	}

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
