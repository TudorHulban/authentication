package app

import (
	"errors"
	"net/http"

	"github.com/TudorHulban/authentication/app/constants"
	"github.com/TudorHulban/authentication/apperrors"
	appuser "github.com/TudorHulban/authentication/domain/app-user"
	"github.com/TudorHulban/authentication/domain/ticket"
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

	var params ticket.ParamsSearchTickets

	if errBody := c.BodyParser(&params); errBody != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{
				"success": false,
				"error":   errBody.Error(),
			},
		)
	}

	reconstructedTickets, errGetTickets := a.ServiceTicket.SearchTickets(
		c.Context(),
		&params,
	)
	if errGetTickets != nil {
		if errors.As(
			errGetTickets,
			&apperrors.ErrNoEntriesFound{},
		) {
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

		return c.Status(
			fiber.StatusInternalServerError).
			JSON(
				&fiber.Map{
					"success": false,
					"error":   errGetTickets.Error(),
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
