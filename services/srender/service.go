package srender

import (
	"github.com/TudorHulban/authentication/helpers"
	"github.com/TudorHulban/authentication/services/suser"
)

type Service struct {
	serviceUser *suser.Service
}

type PiersServiceRender struct {
	ServiceUser *suser.Service
}

func NewServiceRender(piers *PiersServiceRender) (*Service, error) {
	if errVa := helpers.ValidatePiers(piers); errVa != nil {
		return nil, errVa
	}

	return &Service{
			serviceUser: piers.ServiceUser,
		},
		nil
}
