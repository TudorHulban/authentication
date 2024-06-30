package srender

import (
	"github.com/TudorHulban/authentication/helpers"
	"github.com/TudorHulban/authentication/services/sticket"
	"github.com/TudorHulban/authentication/services/suser"
)

type Service struct {
	serviceUser   *suser.Service
	serviceTicket *sticket.Service
}

type PiersServiceRender struct {
	ServiceUser   *suser.Service
	ServiceTicket *sticket.Service
}

func NewServiceRender(piers *PiersServiceRender) (*Service, error) {
	if errVa := helpers.ValidatePiers(piers); errVa != nil {
		return nil, errVa
	}

	return &Service{
			serviceUser:   piers.ServiceUser,
			serviceTicket: piers.ServiceTicket,
		},
		nil
}
