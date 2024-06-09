package pages

import (
	"fmt"

	appuser "github.com/TudorHulban/authentication/domain/app-user"
	g "github.com/maragudk/gomponents"
	html "github.com/maragudk/gomponents/html"
)

func UserSalutation(user *appuser.User) g.Node {
	return html.Div(
		html.H3(
			g.Text(
				fmt.Sprintf(
					"Hi %s!",
					user.Name,
				),
			),
		),
	)
}
