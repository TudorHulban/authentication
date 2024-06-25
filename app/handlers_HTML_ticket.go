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
)

// should be called by POST by search in form.
func (a *App) HandlerHTMLTicketTableBody(c *fiber.Ctx) error {
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
					"handler": "HandlerHTMLTicketTableBody - helpers.ParseMultipartForm", // development only
				},
			)
	}

	params := ticket.NewParamsSearchTicketEventsFromMap(
		responseForm,
	)

	reconstructedTicketEvents, errGetTicketEvents := a.ServiceTicket.SearchTicketEvents(
		c.Context(),
		params,
	)
	if errGetTicketEvents != nil {
		if errors.As(
			errGetTicketEvents,
			&apperrors.ErrNoEntriesFound{},
		) {
			return a.serviceRender.
				TableItemsBodyForTicketEventsWContent(
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
		TableItemsBodyForTicketEventsWContent(
			c.Context(),
			&srender.ParamsRenderTicketEvents{
				Events: reconstructedTicketEvents,

				CSSIDTicketEventsBody: constants.IDItemsTableBody,
			},
		).
		Render(c)
}

func (a *App) HandlerHTMLTicketIDFull(c *fiber.Ctx) error {
	userLogged, errGetUser := appuser.ExtractLoggedUserFrom(c.Context())
	if errGetUser != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(
				&fiber.Map{
					"success": false,
					"error":   errGetUser.Error(),
					"handler": "HandlerHTMLTicketIDFull - ExtractLoggedUserFrom", // development only
				},
			)
	}

	reconstructedTicket, errGetTicket := a.ServiceTicket.GetTicketByID(
		c.Context(),
		&sticket.ParamsGetTicketByID{
			TicketID:     c.Params("id"),
			UserLoggedID: userLogged.PrimaryKey,
		},
	)
	if errGetTicket != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(
				&fiber.Map{
					"success": false,
					"error":   errGetTicket,
					"level":   "GetTicketByID",
				},
			)
	}

	reconstructedTicketEvents, errGetEvents := a.ServiceTicket.GetEventsForTicketID(
		c.Context(),
		helpers.PrimaryKey(reconstructedTicket.PrimaryKey),
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

	return c.Send(
		srender.RenderNodes(
			a.HTMLWithTicketEventsWContent(
				c.Context(),
				&ParamsHTMLWithTicketEventsWContent{
					TicketEvents: reconstructedTicketEvents,
					TicketID:     reconstructedTicket.PrimaryKey,
				},
			)...,
		),
	)
}

// should be called by GET in sidebar menu.
func (a *App) HandlerHTMLTicketEventsWContentTable(c *fiber.Ctx) error {
	return c.Send(
		srender.RenderNodes(
			a.HTMLWithTicketEventsWContent(
				c.Context(),
				nil,
			)...,
		),
	)
}
