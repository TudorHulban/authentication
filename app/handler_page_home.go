package app

import (
	"github.com/TudorHulban/authentication/app/constants"
	appuser "github.com/TudorHulban/authentication/domain/app-user"
	"github.com/TudorHulban/authentication/domain/ticket"
	"github.com/TudorHulban/authentication/helpers"
	"github.com/TudorHulban/authentication/services/srender"
	"github.com/gofiber/fiber/v2"
	g "github.com/maragudk/gomponents"
	co "github.com/maragudk/gomponents/components"
)

func (a *App) HandlerHomePage(c *fiber.Ctx) error {
	userLogged, errGetUser := appuser.ExtractLoggedUserFrom(c.Context())
	if errGetUser != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	reconstructedTickets, errGetTickets := a.ServiceTicket.SearchTickets(
		c.Context(),
		&ticket.ParamsSearchTickets{
			ParamsPagination: helpers.ParamsPagination{
				First: 10,
			},
		},
	)
	if errGetTickets != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(
				&fiber.Map{
					"success": false,
					"error":   errGetTickets,
				},
			)
	}

	menu, errMenu := a.newMenuSidebar()
	if errMenu != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(
				&fiber.Map{
					"success": false,
					"error":   errMenu,
					"handler": "HandlerLoggedInPage - srender.NewMenuSidebar", // development only
				},
			)
	}

	page := co.HTML5(
		co.HTML5Props{
			Title:       "Home",
			Description: "HTMX Logged",
			Language:    "English",

			Head: append(
				srender.LinksFavicon,
				[]g.Node{
					srender.ScriptHTMX,
					srender.ScriptCommonJS,
					srender.LinkCSSMaterialSymbolOutlined,
					srender.LinkCSSCommon,
				}...,
			),

			Body: srender.Body(
				&srender.ParamsBody{
					EntriesHeader: []g.Node{
						srender.Header(),
						srender.UserSalutation(userLogged),
					},

					SidebarMenu: menu,

					EntriesMain: []g.Node{
						srender.NewSearchCreateTickets(
							&srender.ParamsNewSearchCreateTickets{
								LabelButtonSearch: "Search",
								LabelButtonCreate: "Create",

								IDEnclosingDiv: "container-search",
								IDTargetHTMX:   constants.IDItemsTableBody,
							},
						),

						a.serviceRender.TableTickets(
							c.Context(),
							&srender.ParamsTableTickets{
								Tickets:   reconstructedTickets,
								URLTicket: a.baseURL() + constants.RouteTicket,
							},
						),

						srender.ModalCreateTicket(
							&srender.ParamsModalCreateTicket{
								URLAddTicket: constants.RouteTicket,
							},
						),
					},
				},
			),
		},
	)

	c.Set("Content-Type", "text/html")

	return page.Render(c)
}
