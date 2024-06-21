package app

import (
	"errors"

	"github.com/TudorHulban/authentication/app/constants"
	"github.com/TudorHulban/authentication/apperrors"
	"github.com/TudorHulban/authentication/domain/ticket"
	"github.com/TudorHulban/authentication/services/srender"
	"github.com/gofiber/fiber/v2"
)

func (a *App) HandlerHTMLTicketEventsTableBody(c *fiber.Ctx) error {
	reconstructedTicketEvents, errGetTicketEvents := a.ServiceTicket.SearchTicketEvents(
		c.Context(),
		ticket.NewParamsSearchTicketEvents(
			c.BodyRaw(),
		),
	)
	if errGetTicketEvents != nil {
		if errors.As(
			errGetTicketEvents,
			&apperrors.ErrNoEntriesFound{},
		) {
			return a.serviceRender.
				RenderTicketEventsTableBody(
					c.Context(),
					&srender.ParamsRenderTicketEvents{
						Events: reconstructedTicketEvents,

						CSSIDTicketEventsBody: constants.IDItemsTableBody,
					},
				).
				Render(c)
		}

		return c.Status(
			fiber.StatusInternalServerError).
			JSON(
				&fiber.Map{
					"success": false,
					"error":   errGetTicketEvents.Error(),
				},
			)
	}

	return a.serviceRender.
		RenderTicketEventsTableBody(
			c.Context(),
			&srender.ParamsRenderTicketEvents{
				Events: reconstructedTicketEvents,

				CSSIDTicketEventsBody: constants.IDItemsTableBody,
			},
		).
		Render(c)
}

func (a *App) HandlerHTMLTicketEventsTable(c *fiber.Ctx) error {
	reconstructedTicketEvents, errGetTicketEvents := a.ServiceTicket.SearchTicketEvents(
		c.Context(),
		ticket.NewParamsSearchTicketEvents(
			c.BodyRaw(),
		),
	)
	if errGetTicketEvents != nil {
		if errors.As(
			errGetTicketEvents,
			&apperrors.ErrNoEntriesFound{},
		) {
			return c.Send(
				srender.RenderNodes(
					a.serviceRender.TableTicketEventsHead(
						constants.IDItemsTableHead,
					),

					a.serviceRender.RenderTicketEventsTableBody(
						c.Context(),
						&srender.ParamsRenderTicketEvents{
							Events: reconstructedTicketEvents,

							CSSIDTicketEventsBody: constants.IDItemsTableBody,
						},
					),
				),
			)
		}

		return c.Status(
			fiber.StatusInternalServerError).
			JSON(
				&fiber.Map{
					"success": false,
					"error":   errGetTicketEvents.Error(),
				},
			)
	}

	return c.Send(
		srender.RenderNodes(
			a.serviceRender.TableTicketEventsHead(
				constants.IDItemsTableHead,
			),

			a.serviceRender.RenderTicketEventsTableBody(
				c.Context(),
				&srender.ParamsRenderTicketEvents{
					Events: reconstructedTicketEvents,

					CSSIDTicketEventsBody: constants.IDItemsTableBody,
				},
			),
		),
	)
}
