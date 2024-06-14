package app

import (
	"strconv"
	"time"

	appuser "github.com/TudorHulban/authentication/domain/app-user"
	"github.com/TudorHulban/authentication/domain/ticket"
	"github.com/TudorHulban/authentication/helpers"
	"github.com/TudorHulban/authentication/services/srender"
	"github.com/TudorHulban/authentication/services/suser"
	"github.com/gofiber/fiber/v2"

	g "github.com/maragudk/gomponents"
	co "github.com/maragudk/gomponents/components"
)

func (a *App) HandlerLoginPage(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/html")

	return srender.PageLogin("HandlerLoginPage").
		Render(c)
}

func (a *App) HandlerLoggedInPage(c *fiber.Ctx) error {
	userLogged, errGetUser := appuser.ExtractLoggedUserFrom(c.Context())
	if errGetUser != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
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
						srender.ButtonCreateTicket("Create Ticket"),
						a.serviceRender.TableTickets(
							c.Context(),
							&srender.ParamsTableTickets{
								Tickets:   reconstructedTasks,
								URLTicket: a.baseURL() + RouteTicket,
							},
						),

						srender.ModalCreateTicket(
							&srender.ParamsModalCreateTicket{
								URLAddTicket: RouteTicket,
							},
						),
						srender.ScriptCreateTicket(RouteLogged),
					},
				},
			),
		},
	)

	c.Set("Content-Type", "text/html")

	return page.Render(c)
}

func (a *App) HandlerLoginRequest(c *fiber.Ctx) error {
	var params suser.ParamsGetUser

	if err := c.BodyParser(&params); err != nil {
		return err
	}

	reconstructedUser, errGetItem := a.ServiceUser.GetUserByCredentials(
		c.Context(),
		&params,
	)
	if errGetItem != nil {
		c.Set("Content-Type", "text/html")

		return srender.PageLogin("HandlerLoginRequest - a.ServiceUser.GetUser").
			Render(c)
	}

	sessionID, errCacheLoggedUser := a.serviceSessions.PutUserTTL(
		reconstructedUser,
	)
	if errCacheLoggedUser != nil {
		return errCacheLoggedUser
	}

	c.Cookie(
		&fiber.Cookie{
			Name: CookieLoggedUser,
			Value: strconv.Itoa(
				int(
					sessionID,
				),
			) + "|" + time.Now().String(),
		},
	)

	c.Set("HX-Redirect", RouteLogged)

	return c.SendStatus(fiber.StatusOK)
}
