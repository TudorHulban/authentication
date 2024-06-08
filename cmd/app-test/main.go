package main

import (
	"fmt"
	"log"
	"os"

	"github.com/CloudyKit/jet/v6"
	"github.com/TudorHulban/authentication/domain/ticket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("../../views"),
	jet.InDevelopmentMode(),
)

func main() {
	app := fiber.New()

	app.Use(
		logger.New(),
	)

	app.Get("/",
		func(c *fiber.Ctx) error {
			view, err := views.GetTemplate("pages/ticket.jet")
			if err != nil {
				log.Println("Unexpected template err:", err.Error())
			}

			vars := jet.VarMap{}

			item, errCr := ticket.NewTicketFrom(
				`
				{
					"PrimaryKey": 17177738116406920000,
					"Name": "Ticket 1",
					"OpenedByUserID": 1,
					"CreatedAt": 1717773811682774586,
					"DeletedAt": {
						"Int64": 0,
						"Valid": false
					}
				}
				`,
			)
			if errCr != nil {
				fmt.Println(errCr)

				os.Exit(5)
			}

			vars.Set(
				"ticket",
				item,
			)

			vars.Set(
				"events",
				nil,
			)

			vars.Set(
				"routeAddEvent",
				"",
			)

			return view.Execute(
				c,
				vars,
				nil,
			)
		},
	)

	log.Fatal(app.Listen(":3000"))
}
