package app

import (
	"errors"

	"github.com/TudorHulban/authentication/app/constants"
	"github.com/TudorHulban/authentication/apperrors"
	"github.com/TudorHulban/authentication/domain/ticket"
	"github.com/TudorHulban/authentication/helpers"
	"github.com/TudorHulban/authentication/services/srender"
	"github.com/gofiber/fiber/v2"
)

// HandlerHTMLTicketsTableBody should be used for search,
// when change of table header is not needed.
func (a *App) HandlerHTMLTicketsTableBody(c *fiber.Ctx) error {
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
					"handler": "HandlerHTMLTicketsTableBody - helpers.ParseMultipartForm", // development only
				},
			)
	}

	reconstructedTickets, errGetTickets := a.ServiceTicket.SearchTickets(
		c.Context(),
		ticket.NewParamsSearchTicketsFromMap(
			responseForm,
		),
	)
	if errGetTickets != nil {
		if errors.As(
			errGetTickets,
			&apperrors.ErrNoEntriesFound{},
		) {
			return a.serviceRender.
				TableItemsBodyForTickets(
					c.Context(),
					&srender.ParamsRenderTickets{
						Tickets: reconstructedTickets,

						RouteTicket:     constants.RouteTickets,
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
		TableItemsBodyForTickets(
			c.Context(),
			&srender.ParamsRenderTickets{
				Tickets: reconstructedTickets,

				RouteTicket:     constants.RouteTickets,
				CSSIDTicketBody: constants.IDItemsTableBody,
			},
		).
		Render(c)
}

// should be called by sidebar menu entry.
func (a *App) HandlerHTMLTicketsTable(c *fiber.Ctx) error {
	reconstructedTickets, errGetTickets := a.ServiceTicket.SearchTickets(
		c.Context(),
		nil,
	)
	if errGetTickets != nil {
		if errors.As(
			errGetTickets,
			&apperrors.ErrNoEntriesFound{},
		) {
			return c.Send(
				srender.RenderNodes(
					a.serviceRender.HTMLTableItemsForTickets(
						c.Context(),
						&srender.ParamsHTMLTableItemsForTickets{
							IDItemsTableHead: constants.IDItemsTableHead,
							RouteTickets:     constants.RouteTickets,
							CSSIDTicketBody:  constants.IDItemsTableBody,
							Tickets:          reconstructedTickets,
						},
					)...,
				),
			)
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

	responseBytes := srender.RenderNodes(
		a.HTMLWithTickets(
			c.Context(),
			reconstructedTickets,
		)...,
	)

	return c.Send(
		responseBytes,
	)
}
