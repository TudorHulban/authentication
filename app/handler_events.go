package app

import (
	appuser "github.com/TudorHulban/authentication/domain/app-user"
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

	var params sticket.ParamsAddEvent

	if errValidateBody := c.BodyParser(&params); errValidateBody != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(
				&fiber.Map{
					"success": false,
					"error":   errValidateBody,
					"handler": "HandlerAddEvent - c.BodyParser", // development only
				},
			)
	}

	params.OpenedByUserID = userLogged.PrimaryKey

	errAddEvent := a.ServiceTicket.AddEvent(
		c.Context(),
		&params,
	)
	if errAddEvent != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(
				&fiber.Map{
					"success": false,
					"error":   errAddEvent,
					"handler": "HandlerAddEvent - serviceTask.CreateTask", // development only
				},
			)
	}

	return c.SendStatus(fiber.StatusOK)
}
