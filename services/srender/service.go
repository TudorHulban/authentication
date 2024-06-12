package srender

import "github.com/TudorHulban/authentication/services/suser"

type ServiceRender struct {
	serviceUser *suser.Service
}

type PiersServiceRender struct {
	ServiceUser suser.Service
}

func NewServiceRender(piers *PiersServiceRender) (*ServiceRender, error) {
	// TODO: input validation

	return &ServiceRender{
			serviceUser: &piers.ServiceUser,
		},
		nil
}
