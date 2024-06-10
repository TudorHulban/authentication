package app

import (
	appuser "github.com/TudorHulban/authentication/domain/app-user"
	"github.com/TudorHulban/authentication/domain/ticket"
	"github.com/TudorHulban/authentication/helpers"
	"github.com/TudorHulban/authentication/pages"
	"github.com/TudorHulban/authentication/services/sticket"

	"github.com/gofiber/fiber/v2"
	g "github.com/maragudk/gomponents"
	co "github.com/maragudk/gomponents/components"
)

func (a *App) HandlerAddTicket(c *fiber.Ctx) error {
	userLogged, errGetUser := appuser.ExtractLoggedUserFrom(c.Context())
	if errGetUser != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(
				&fiber.Map{
					"success": false,
					"error":   errGetUser,
					"handler": "HandlerAddTicket - ExtractLoggedUserFrom", // development only
				},
			)
	}

	var params sticket.ParamsCreateTicket

	if errValidateBody := c.BodyParser(&params); errValidateBody != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(
				&fiber.Map{
					"success": false,
					"error":   errValidateBody,
					"handler": "HandlerAddTicket - c.BodyParser", // development only
				},
			)
	}

	pkConstructedTicket, errGetTicket := a.serviceTicket.CreateTicket(
		c.Context(),
		&sticket.ParamsCreateTicket{
			OpenedByUserID: userLogged.PrimaryKey,
			TicketName:     params.TicketName,
			TicketKind:     params.TicketKind,
		},
	)
	if errGetTicket != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(
				&fiber.Map{
					"success": false,
					"error":   errGetTicket,
					"handler": "HandlerAddTicket - serviceTask.CreateTicket", // development only
				},
			)
	}

	reconstructedTicket, errGet := a.serviceTicket.GetTicketByID(
		c.Context(),
		&sticket.ParamsGetTicketByID{
			TicketID:     pkConstructedTicket.String(),
			UserLoggedID: userLogged.PrimaryKey,
		},
	)
	if errGet != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(
				&fiber.Map{
					"success": false,
					"error":   errGet,
					"handler": "HandlerAddTask - serviceTask.GetTicketByID", // development only
				},
			)
	}

	return c.Status(fiber.StatusOK).JSON(
		reconstructedTicket.PrimaryKey,
	)
}

func (a *App) HandlerTickets(c *fiber.Ctx) error {
	_, errGetUser := appuser.ExtractLoggedUserFrom(c.Context())
	if errGetUser != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(
				&fiber.Map{
					"success": false,
					"error":   errGetUser,
				},
			)
	}

	reconstructedTasks, errGetTasks := a.serviceTicket.SearchTasks(
		c.Context(),
		&ticket.ParamsSearchTasks{
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
			Title:       "Tickets",
			Description: "HTMX Login",
			Language:    "English",
			Head: []g.Node{
				pages.ScriptHTMX,
				pages.ScriptCommonJS,
				pages.LinkCSSWater,
				pages.LinkCSSCommon,
			},
			Body: []g.Node{
				pages.Header(),
				pages.TableTickets(
					&pages.ParamsTableTickets{
						Tickets:   reconstructedTasks,
						URLTicket: a.baseURL() + RouteTicket,
					},
				),
				pages.ButtonCreateTicket("Create Ticket"),
				pages.ModalCreateTicket(
					&pages.ParamsModalCreateTicket{
						URLAddTicket: RouteTicket,
					},
				),
				pages.ScriptCreateTicket(),
				pages.Footer(),
			},
		},
	)

	c.Set("Content-Type", "text/html")

	return page.Render(c)
}

func (a *App) HandlerTicketID(c *fiber.Ctx) error {
	userLogged, errGetUser := appuser.ExtractLoggedUserFrom(c.Context())
	if errGetUser != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(
				&fiber.Map{
					"success": false,
					"error":   errGetUser,
				},
			)
	}

	reconstructedTask, errGetTask := a.serviceTicket.GetTicketByID(
		c.Context(),
		&sticket.ParamsGetTicketByID{
			TicketID:     c.Params("id"),
			UserLoggedID: userLogged.PrimaryKey,
		},
	)
	if errGetTask != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(
				&fiber.Map{
					"success": false,
					"error":   errGetTask,
					"level":   "GetTaskByID",
				},
			)
	}

	reconstructedEvents, errGetEvents := a.serviceTicket.GetEventsForTicketID(
		c.Context(),
		helpers.PrimaryKey(reconstructedTask.PrimaryKey),
	)
	if errGetEvents != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(
				&fiber.Map{
					"success": false,
					"error":   errGetEvents,
				},
			)
	}

	return c.Render(
		"pages"+RouteTicket,

		fiber.Map{
			"title":  "T" + reconstructedTask.PrimaryKey.String(),
			"name":   userLogged.Name,
			"ticket": reconstructedTask,
			"events": reconstructedEvents,

			"routeAddEvent": RouteEvent,

			"UnixNanoTo": helpers.UnixNanoTo,
		},

		"layouts/base",
	)
}
