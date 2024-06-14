package main

import (
	"log"

	"github.com/TudorHulban/authentication/services/srender"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	g "github.com/maragudk/gomponents"
	co "github.com/maragudk/gomponents/components"
)

func main() {
	app := fiber.New()

	app.Use(
		logger.New(),
	)

	app.Static("/public", "../public")

	app.Get("/",
		func(c *fiber.Ctx) error {
			menu, errMenu := srender.NewMenuSidebar(
				&srender.ParamsMenuSidebar{
					TextLogo:      "Logo",
					PathImageLogo: "../public/images/logo.png",

					Sections: []*srender.MenuSidebarSection{
						{
							TextSection: "Section 1",

							Entries: []*srender.MenuSidebarSectionEntry{
								{
									TextSectionEntry: "Entry 1",
									SymbolEntry:      "dashboard",
								},
								{
									TextSectionEntry: "Entry 2",
								},
							},
						},
					},
				},
			)
			if errMenu != nil {
				return errMenu
			}

			page := co.HTML5(
				co.HTML5Props{
					Title: "Menu Sidebar",

					Head: append(
						srender.LinksFavicon,
						[]g.Node{
							srender.LinkCSSMaterialSymbolOutlined,
							// srender.LinkCSSWater,
							srender.LinkCSSCommon,
						}...,
					),

					Body: srender.Body(menu),
				},
			)

			c.Set("Content-Type", "text/html")

			return page.Render(c)
		},
	)

	log.Fatal(app.Listen(":3000"))
}
