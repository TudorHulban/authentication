package app

import (
	"errors"

	"github.com/TudorHulban/authentication/app/constants"
	"github.com/TudorHulban/authentication/apperrors"
	appuser "github.com/TudorHulban/authentication/domain/app-user"
	"github.com/TudorHulban/authentication/domain/ticket"
	"github.com/TudorHulban/authentication/services/srender"
	"github.com/gofiber/fiber/v2"
)

// HandlerHTMLTicketsTableBody should be used for search,
// when change of table header is not needed.
func (a *App) HandlerHTMLTicketsTableBody(c *fiber.Ctx) error {
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
		ticket.NewParamsSearchTickets(
			c.BodyRaw(),
		),
	)
	if errGetTickets != nil {
		if errors.As(
			errGetTickets,
			&apperrors.ErrNoEntriesFound{},
		) {
			return a.serviceRender.
				RenderTicketsTableBody(
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
		RenderTicketsTableBody(
			c.Context(),
			&srender.ParamsRenderTickets{
				Tickets: reconstructedTickets,

				RouteTicket:     a.baseURL() + constants.RouteTickets,
				CSSIDTicketBody: constants.IDItemsTableBody,
			},
		).
		Render(c)
}

func (a *App) HandlerHTMLTicketsTable(c *fiber.Ctx) error {
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
		nil,
	)
	if errGetTickets != nil {
		if errors.As(
			errGetTickets,
			&apperrors.ErrNoEntriesFound{},
		) {
			return a.serviceRender.
				RenderTicketsTableBody(
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

	return a.serviceRender.RenderTicketsTableBody(
		c.Context(),
		&srender.ParamsRenderTickets{
			Tickets: reconstructedTickets,

			RouteTicket:     a.baseURL() + constants.RouteTickets,
			CSSIDTicketBody: constants.IDItemsTableBody,
		},
	).
		Render(c)
}
