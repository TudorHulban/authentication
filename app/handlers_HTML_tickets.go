package app

import (
	"errors"
	"fmt"

	"github.com/TudorHulban/authentication/app/constants"
	"github.com/TudorHulban/authentication/apperrors"
	"github.com/TudorHulban/authentication/domain/ticket"
	"github.com/TudorHulban/authentication/services/srender"
	"github.com/gofiber/fiber/v2"
)

// HandlerHTMLTicketsTableBody should be used for search,
// when change of table header is not needed.
func (a *App) HandlerHTMLTicketsTableBody(c *fiber.Ctx) error {
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
				TableItemsBodyForTickets(
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
		TableItemsBodyForTickets(
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
					a.serviceRender.TableItemsHeadForTickets(
						constants.IDItemsTableHead,
					),

					a.serviceRender.TableItemsBodyForTickets(
						c.Context(),
						&srender.ParamsRenderTickets{
							Tickets: reconstructedTickets,

							RouteTicket:     a.baseURL() + constants.RouteTickets,
							CSSIDTicketBody: constants.IDItemsTableBody,
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
					"error":   errGetTickets.Error(),
				},
			)
	}

	responseBytes := srender.RenderNodes(
		a.formSearchCreateTickets(),

		a.serviceRender.TableItemsHeadForTickets(
			constants.IDItemsTableHead,
		),

		a.serviceRender.TableItemsBodyForTickets(
			c.Context(),
			&srender.ParamsRenderTickets{
				Tickets: reconstructedTickets,

				RouteTicket:     a.baseURL() + constants.RouteTickets,
				CSSIDTicketBody: constants.IDItemsTableBody,
			},
		),
	)

	fmt.Println(
		string(
			responseBytes,
		),
	)

	return c.Send(
		responseBytes,
	)
}
