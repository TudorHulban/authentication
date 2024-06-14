package app

import "github.com/TudorHulban/authentication/services/srender"

func (a *App) newMenuSidebar() (*srender.MenuSidebar, error) {
	return srender.NewMenuSidebar(
		srender.ParamsMenuSidebarToUse(
			&srender.ParamsCurrentMenuSidebar{
				TextLogo:      CompanyName,
				PathImageLogo: PathImageLogo,

				TextSection1: "Work",

				TextSection1Entry1:   "Tickets",
				SymbolSection1Entry1: "call",
				URLSection1Entry1:    a.baseURL() + RouteTickets,

				TextSection1Entry2:   "Ticket Events",
				SymbolSection1Entry2: "comment",
			},
		),
	)
}
