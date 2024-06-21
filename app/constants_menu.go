package app

import (
	"github.com/TudorHulban/authentication/app/constants"
	"github.com/TudorHulban/authentication/services/srender"
)

func (a *App) newMenuSidebar() (*srender.MenuSidebar, error) {
	return srender.NewMenuSidebar(
		srender.ParamsMenuSidebarToUse(
			&srender.ParamsCurrentMenuSidebar{
				TextLogo:      CompanyName,
				PathImageLogo: PathImageLogo,

				TextSection1: "Work",

				TextSection1Entry1:   "Tickets",
				SymbolSection1Entry1: "call",
				URLSection1Entry1:    constants.RouteTickets,

				HTMXTargetsSection1Entry1: []string{
					constants.IDItemsTableHead,
					constants.IDItemsTableBody,
				},

				TextSection1Entry2:   "Ticket Events",
				SymbolSection1Entry2: "comment",
				URLSection1Entry2:    constants.RouteTicketEvents,

				HTMXTargetsSection1Entry2: []string{
					constants.IDItemsTableHead,
					constants.IDItemsTableBody,
				},
			},
		),
	)
}
