package app

import (
	"github.com/TudorHulban/authentication/app/constants"
	appuser "github.com/TudorHulban/authentication/domain/app-user"
	"github.com/TudorHulban/authentication/domain/ticket"
	"github.com/TudorHulban/authentication/helpers"
	"github.com/TudorHulban/authentication/services/srender"
	"github.com/TudorHulban/authentication/services/sticket"

	"github.com/gofiber/fiber/v2"
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

	reconstructedTicket, errGet := a.ServiceTicket.GetTicketByIDString(
		c.Context(),
		&sticket.ParamsGetTicketByIDString{
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

				RouteGetTicket:  constants.RouteGetTicket,
				CSSIDTicketBody: constants.IDItemsTableBody,

				TargetsSwapSearch: swapTargetsBodyTicket,
			},
		).
		Render(c)
}
