package app

import (
	"github.com/TudorHulban/authentication/helpers"
	"github.com/TudorHulban/authentication/services/suser"
)

type App struct {
	ServiceUser suser.Service
}

type PiersApp struct {
	ServiceUser *suser.Service
}

func NewApp(piers *PiersApp) (*App, error) {
	if errValidate := helpers.ValidatePiers(piers); errValidate != nil {
		return nil,
			errValidate
	}

	return &App{
			ServiceUser: *piers.ServiceUser,
		},
		nil
}
