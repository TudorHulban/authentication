package app

import (
	"github.com/TudorHulban/authentication/app/constants"
	appuser "github.com/TudorHulban/authentication/domain/app-user"
	"github.com/TudorHulban/authentication/domain/ticket"
	"github.com/TudorHulban/authentication/helpers"
	"github.com/gofiber/fiber/v2"
)

func (a *App) HandlerHTMLTickets(c *fiber.Ctx) error {
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

	reconstructedTasks, errGetTasks := a.ServiceTicket.SearchTickets(
		c.Context(),
		&ticket.ParamsSearchTickets{
			ParamsPagination: helpers.ParamsPagination{
				First: 10,
			},
		},
	)
	if errGetTasks != nil {
		return c.Status(
			fiber.StatusInternalServerError).
			JSON(
				&fiber.Map{
					"success": false,
					"error":   errGetTasks,
				},
			)
	}

	return c.SendString(
		reconstructedTasks.AsHTMLTBody(
			ticket.ParamsAsHTMLTBody{
				RouteTicket:     a.baseURL() + constants.RouteTickets,
				CSSIDTicketBody: constants.IDItemsTableBody,
			},
		),
	)
}
