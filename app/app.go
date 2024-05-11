package app

import (
	"fmt"

	"github.com/TudorHulban/authentication/apperrors"
	"github.com/TudorHulban/authentication/helpers"
	"github.com/TudorHulban/authentication/services/suser"
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/jet/v2"
)

type App struct {
	ServiceUser suser.Service

	port string

	Transport *fiber.App
}

type ParamsNewApp struct {
	TemplateFolder         string `valid:"required"`
	TemplateFilesExtension string `valid:"required"`

	Port string `valid:"required"`
}

type PiersApp struct {
	ServiceUser *suser.Service
}

func NewApp(params *ParamsNewApp, piers *PiersApp) (*App, error) {
	if _, errValidation := govalidator.ValidateStruct(params); errValidation != nil {
		return nil,
			apperrors.ErrValidation{
				Caller: "NewApp",
				Issue:  errValidation,
			}
	}

	if errValidate := helpers.ValidatePiers(piers); errValidate != nil {
		return nil,
			errValidate
	}

	return &App{
			ServiceUser: *piers.ServiceUser,

			Transport: fiber.New(
				fiber.Config{
					Views: jet.New(
						params.TemplateFolder,
						params.TemplateFilesExtension,
					),
				},
			),

			port: params.Port,
		},
		nil
}

func (a *App) Start() error {
	fmt.Printf(
		"listening on localhost:%s\n",
		a.port,
	)
	fmt.Printf(
		"routes:%d\n",
		len(a.Transport.GetRoutes()),
	)

	return a.Transport.Listen(":" + a.port)
}
