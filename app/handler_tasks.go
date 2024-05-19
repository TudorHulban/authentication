package app

import (
	appuser "github.com/TudorHulban/authentication/domain/app-user"
	"github.com/TudorHulban/authentication/domain/task"
	"github.com/gofiber/fiber/v2"
)

func (a *App) HandlerTasksPage(c *fiber.Ctx) error {
	user, errGetUser := appuser.ExtractLoggedUserFrom(c.Context())
	if errGetUser != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	reconstructedTasks, errGetTasks := a.serviceTask.SearchTasks(
		c.Context(),
		&task.ParamsSearchTasks{},
	)
	if errGetTasks != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Render(
		"pages/tasks",
		fiber.Map{
			"name":  user.Name,
			"tasks": reconstructedTasks,
		},
		"layouts/base",
	)
}
