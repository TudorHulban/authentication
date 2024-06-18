package app

import (
	"fmt"
	"net/http"

	"github.com/TudorHulban/authentication/app/constants"
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

	fmt.Println(
		string(c.BodyRaw()), // id=&status=&name=aaaaaaaaaa
	)

	if errBody := c.BodyParser(&params); errBody != nil {
		return c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{
				"success": false,
				"error":   errBody.Error(),
			},
		)
	}

	// fmt.Printf(
	// 	"%#v",
	// 	params,
	// )

	reconstructedTickets, errGetTickets := a.ServiceTicket.SearchTickets(
		c.Context(),
		&params,
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
