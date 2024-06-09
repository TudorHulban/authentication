package pages

import (
	g "github.com/maragudk/gomponents"
	html "github.com/maragudk/gomponents/html"
)

type ParamsNavigation struct {
	WhereTo        string
	LabelToDisplay string
}

func Navigation(params *ParamsNavigation) g.Node {
	return html.A(
		html.Href(params.WhereTo),
		g.Text(params.LabelToDisplay),
	)
}
