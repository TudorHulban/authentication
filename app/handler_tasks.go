package app

import (
	appuser "github.com/TudorHulban/authentication/domain/app-user"
	"github.com/TudorHulban/authentication/domain/task"
	"github.com/TudorHulban/authentication/helpers"
	"github.com/TudorHulban/authentication/services/stask"
	"github.com/gofiber/fiber/v2"
)

func (a *App) HandlerAddTask(c *fiber.Ctx) error {
	userLogged, errGetUser := appuser.ExtractLoggedUserFrom(c.Context())
	if errGetUser != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(
				&fiber.Map{
					"success": false,
					"error":   errGetUser,
					"handler": "HandlerAddTask - ExtractLoggedUserFrom", // development only
				},
			)
	}

	var params stask.ParamsCreateTask

	if errValidateBody := c.BodyParser(&params); errValidateBody != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(
				&fiber.Map{
					"success": false,
					"error":   errValidateBody,
					"handler": "HandlerAddTask - c.BodyParser", // development only
				},
			)
	}

	pkConstructedTask, errGetTask := a.serviceTask.CreateTask(
		c.Context(),
		&stask.ParamsCreateTask{
			OpenedByUserID: userLogged.ID,
			TaskName:       params.TaskName,
			TaskKind:       params.TaskKind,
		},
	)
	if errGetTask != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(
				&fiber.Map{
					"success": false,
					"error":   errGetTask,
					"handler": "HandlerAddTask - serviceTask.CreateTask", // development only
				},
			)
	}

	reconstructedTask, errGet := a.serviceTask.GetTaskByID(
		c.Context(),
		&stask.ParamsGetTaskByID{
			TaskID:       pkConstructedTask.String(),
			UserLoggedID: userLogged.ID,
		},
	)
	if errGet != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(
				&fiber.Map{
					"success": false,
					"error":   errGet,
					"handler": "HandlerAddTask - serviceTask.GetTaskByID", // development only
				},
			)
	}

	return c.Status(fiber.StatusOK).JSON(
		reconstructedTask.PrimaryKeyTicket,
	)
}

func (a *App) HandlerTasks(c *fiber.Ctx) error {
	userLogged, errGetUser := appuser.ExtractLoggedUserFrom(c.Context())
	if errGetUser != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(
				&fiber.Map{
					"success": false,
					"error":   errGetUser,
				},
			)
	}

	reconstructedTasks, errGetTasks := a.serviceTask.SearchTasks(
		c.Context(),
		&task.ParamsSearchTasks{
			ParamsPagination: helpers.ParamsPagination{
				First: 10,
			},
		},
	)
	if errGetTasks != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Render(
		"pages/tasks",
		fiber.Map{
			"name":         userLogged.Name,
			"tasks":        reconstructedTasks,
			"baseURL":      a.baseURL(),
			"route":        a.baseURL() + RouteTask,
			"routeAddTask": RouteTask,
			"routeTasks":   RouteTasks,
		},
		"layouts/base",
	)
}

func (a *App) HandlerTaskID(c *fiber.Ctx) error {
	userLogged, errGetUser := appuser.ExtractLoggedUserFrom(c.Context())
	if errGetUser != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(
				&fiber.Map{
					"success": false,
					"error":   errGetUser,
				},
			)
	}

	reconstructedTask, errGetTask := a.serviceTask.GetTaskByID(
		c.Context(),
		&stask.ParamsGetTaskByID{
			TaskID:       c.Params("id"),
			UserLoggedID: userLogged.ID,
		},
	)
	if errGetTask != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(
				&fiber.Map{
					"success": false,
					"error":   errGetTask,
					"level":   "GetTaskByID",
				},
			)
	}

	reconstructedEvents, errGetEvents := a.serviceTask.GetEventsForTaskID(
		c.Context(),
		helpers.PrimaryKey(reconstructedTask.PrimaryKeyTicket),
	)
	if errGetEvents != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(
				&fiber.Map{
					"success": false,
					"error":   errGetEvents,
				},
			)
	}

	return c.Render(
		"pages/task",
		fiber.Map{
			"name":   userLogged.Name,
			"task":   reconstructedTask,
			"events": reconstructedEvents,

			"UnixNanoTo": helpers.UnixNanoTo,
		},
		"layouts/base",
	)
}
