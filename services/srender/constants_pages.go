package srender

import (
	g "github.com/maragudk/gomponents"
	co "github.com/maragudk/gomponents/components"
)

func PageLogin(caller string) g.Node {
	return co.HTML5(
		co.HTML5Props{
			Title:       "Login",
			Description: "HTMX Login",
			Language:    "English",
			Head: []g.Node{
				LinkCSSWater,
			},
			Body: []g.Node{
				Header(),
				FormLogin(caller),
			},
		},
	)
}
