package app

import (
	"errors"

	"github.com/TudorHulban/authentication/app/constants"
	"github.com/TudorHulban/authentication/apperrors"
	appuser "github.com/TudorHulban/authentication/domain/app-user"
	"github.com/TudorHulban/authentication/domain/ticket"
	"github.com/TudorHulban/authentication/helpers"
	"github.com/TudorHulban/authentication/services/srender"
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

	responseForm, errCr := helpers.ParseMultipartForm(
		c.BodyRaw(),
		c.GetReqHeaders(),
	)
	if errCr != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(
				&fiber.Map{
					"success": false,
					"error":   errCr.Error(),
					"handler": "HandlerAddTicket - helpers.ParseMultipartForm", // development only
				},
			)
	}

	params := sticket.NewParamsCreateTicketFromMap(
		responseForm,
	)

	params.OpenedByUserID = userLogged.PrimaryKey

	pkConstructedTicket, errGetTicket := a.ServiceTicket.CreateTicket(
		c.Context(),
		params,
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

	reconstructedTicket, errGet := a.ServiceTicket.GetTicketByID(
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

	return a.serviceRender.
		TableItemsBodyForTickets(
			c.Context(),
			&srender.ParamsRenderTickets{
				Tickets: []*ticket.Ticket{
					reconstructedTicket,
				},

				RouteTicket:     a.baseURL() + constants.RouteTickets,
				CSSIDTicketBody: constants.IDItemsTableBody,
			},
		).
		Render(c)
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

	reconstructedTask, errGetTask := a.ServiceTicket.GetTicketByID(
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
					"level":   "GetTicketByID",
				},
			)
	}

	reconstructedEvents, errGetEvents := a.ServiceTicket.GetEventsForTicketID(
		c.Context(),
		helpers.PrimaryKey(reconstructedTask.PrimaryKey),
	)
	if errGetEvents != nil && !errors.As(errGetEvents, &apperrors.ErrNoEntriesFound{}) {
		return c.Status(fiber.StatusInternalServerError).
			JSON(
				&fiber.Map{
					"success": false,
					"error":   errGetEvents,
				},
			)
	}

	page := co.HTML5(
		co.HTML5Props{
			Title:       "T" + reconstructedTask.PrimaryKey.String(),
			Description: "Ticket Information",
			Language:    "English",
			Head: []g.Node{
				srender.ScriptCommonJS,
				srender.LinkCSSWater,
				srender.LinkCSSCommon,
			},
			Body: []g.Node{
				srender.Header(),
				srender.TableTicketEvents(
					&srender.ParamsTableTicketEvents{
						TicketEvents: reconstructedEvents,
					},
				),
				srender.ButtonCreateTicketEvent("Create Ticket Event"),
				srender.ModalCreateTicketEvent(
					&srender.ParamsModalCreateTicketEvent{
						URLAddTicketEvent: constants.RouteTicketEvent,
						TicketID:          reconstructedTask.PrimaryKey,
					},
				),
			},
		},
	)

	c.Set("Content-Type", "text/html")

	return page.Render(c)
}
