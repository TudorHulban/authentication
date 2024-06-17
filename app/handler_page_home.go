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

	reconstructedTasks, errGetTasks := a.ServiceTicket.SearchTickets(
		c.Context(),
		&ticket.ParamsSearchTickets{
			ParamsPagination: helpers.ParamsPagination{
				First: 10,
			},
		},
	)
	if errGetTasks != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(
				&fiber.Map{
					"success": false,
					"error":   errGetTasks,
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
						srender.NewFormSearchTickets(
							&srender.ParamsNewFormSearchTickets{
								TextForm:          "Search Tickets",
								ActionForm:        constants.RouteTickets,
								LabelButtonSubmit: "Submit",

								IDEnclosingDiv: "container-search",
							},
						),

						srender.ButtonCreateTicket("Create Ticket"),

						a.serviceRender.TableTickets(
							c.Context(),
							&srender.ParamsTableTickets{
								Tickets:   reconstructedTasks,
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
