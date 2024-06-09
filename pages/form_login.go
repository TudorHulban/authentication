package pages

import (
	"github.com/TudorHulban/authentication/fixtures"
	g "github.com/maragudk/gomponents"
	html "github.com/maragudk/gomponents/html"
)

func FormLogin() g.Node {
	return html.Div(
		html.H2(
			g.Text(
				"HTMX Login",
			),
		),

		html.Form(
			g.Attr("hx-post", "/login"),
			g.Attr("hx-swap", "none"),

			html.Label(
				g.Attr("for", "email"),
				g.Text("Email:"),
			),

			html.Input(
				g.Attr("type", "email"),
				g.Attr("id", "email"),
				g.Attr("name", "email"),
				g.Attr("value", fixtures.TestUser.Email),
			),

			html.Label(
				g.Attr("for", "password"),
				g.Text("Password:"),
			),

			html.Input(
				g.Attr("type", "password"),
				g.Attr("id", "password"),
				g.Attr("name", "password"),
				g.Attr("value", fixtures.TestUser.Password),
			),

			html.Input(
				g.Attr("type", "submit"),
				g.Attr("value", "Submit"),
			),
		),
	)
}
