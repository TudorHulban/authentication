package app

import (
	"github.com/TudorHulban/authentication/pages"
	g "github.com/maragudk/gomponents"
	co "github.com/maragudk/gomponents/components"
)

func pageLogin(caller string) g.Node {
	return co.HTML5(
		co.HTML5Props{
			Title:       "Login",
			Description: "HTMX Login",
			Language:    "English",
			Head: []g.Node{
				pages.ScriptHTMX,
				pages.LinkCSSWater,
			},
			Body: []g.Node{
				pages.Header(),
				pages.FormLogin(caller),
				pages.Footer(),
			},
		},
	)
}
