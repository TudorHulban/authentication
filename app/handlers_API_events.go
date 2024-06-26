package app

import (
	appuser "github.com/TudorHulban/authentication/domain/app-user"
	"github.com/TudorHulban/authentication/helpers"
	"github.com/TudorHulban/authentication/services/srender"
	"github.com/TudorHulban/authentication/services/sticket"
	"github.com/gofiber/fiber/v2"
)

func (a *App) HandlerAddEvent(c *fiber.Ctx) error {
	userLogged, errGetUser := appuser.ExtractLoggedUserFrom(c.Context())
	if errGetUser != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(
				&fiber.Map{
					"success": false,
					"error":   errGetUser,
					"handler": "HandlerAddEvent - ExtractLoggedUserFrom", // development only
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
					"handler": "HandlerAddEvent - helpers.ParseMultipartForm", // development only
				},
			)
	}

	params, errCrParams := sticket.NewParamsAddEvent(
		responseForm,
	)
	if errCrParams != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(
				&fiber.Map{
					"success": false,
					"error":   errCrParams.Error(),
					"handler": "HandlerAddEvent - sticket.NewParamsAddEvent", // development only
				},
			)
	}

	params.OpenedByUserID = userLogged.PrimaryKey

	errAddEvent := a.ServiceTicket.AddEvent(
		c.Context(),
		params,
	)
	if errAddEvent != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(
				&fiber.Map{
					"success": false,
					"error":   errAddEvent,
					"handler": "HandlerAddEvent - ServiceTicket.AddEvent", // development only
				},
			)
	}

	reconstructedTicketEvents, errGetTicketEvents := a.ServiceTicket.GetEventsForTicketID(
		c.Context(),
		params.TicketID,
	)
	if errGetTicketEvents != nil {
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
			a.HTMLWithTicketEventsWContent(
				c.Context(),
				&ParamsHTMLWithTicketEventsWContent{
					TicketEvents: reconstructedTicketEvents,
					TicketID:     params.TicketID,
				},
			)...,
		),
	)
}
