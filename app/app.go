package app

import (
	"github.com/TudorHulban/authentication/apperrors"
	"github.com/TudorHulban/authentication/helpers"
	"github.com/TudorHulban/authentication/services/srender"
	"github.com/TudorHulban/authentication/services/ssessions"
	"github.com/TudorHulban/authentication/services/sticket"
	"github.com/TudorHulban/authentication/services/suser"
	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	ServiceUser     *suser.Service
	serviceSessions *ssessions.Service
	ServiceTicket   *sticket.Service
	serviceRender   *srender.Service

	host string
	port string

	Transport *fiber.App

	authenticationDisabled bool
}

type ParamsNewApp struct {
	Port string `valid:"required"`

	AuthenticationDisabled bool
}

type PiersApp struct {
	ServiceUser     *suser.Service
	ServiceSessions *ssessions.Service

	ServiceTicket *sticket.Service
	ServiceRender *srender.Service
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
			ServiceUser:     piers.ServiceUser,
			serviceSessions: piers.ServiceSessions,
			ServiceTicket:   piers.ServiceTicket,
			serviceRender:   piers.ServiceRender,

			Transport: fiber.New(
				fiber.Config{
					Prefork: false,

					// Views: jet.New(
					// 	params.TemplateFolder,
					// 	params.TemplateFilesExtension,
					// ),
				},
			),

			host: "http://localhost",
			port: params.Port,

			authenticationDisabled: params.AuthenticationDisabled,
		},
		nil
}
