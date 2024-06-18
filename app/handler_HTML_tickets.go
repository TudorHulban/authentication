package app

import (
	"github.com/TudorHulban/authentication/app/constants"
	appuser "github.com/TudorHulban/authentication/domain/app-user"
	"github.com/TudorHulban/authentication/domain/ticket"
	"github.com/TudorHulban/authentication/helpers"
	"github.com/TudorHulban/authentication/services/srender"
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

	reconstructedTickets, errGetTickets := a.ServiceTicket.SearchTickets(
		c.Context(),
		&ticket.ParamsSearchTickets{
			ParamsPagination: helpers.ParamsPagination{
				First: 10,
			},
		},
	)
	if errGetTickets != nil {
		return c.Status(
			fiber.StatusInternalServerError).
			JSON(
				&fiber.Map{
					"success": false,
					"error":   errGetTickets,
				},
			)
	}

	return a.serviceRender.
		RenderTickets(
			c.Context(),
			&srender.ParamsRenderTickets{
				Tickets: reconstructedTickets,

				RouteTicket:     a.baseURL() + constants.RouteTickets,
				CSSIDTicketBody: constants.IDItemsTableBody,
			},
		).
		Render(c)
}
