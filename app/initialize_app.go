package app

import (
	"log"
	"os"

	"github.com/TudorHulban/authentication/apperrors"
	storefilefixtures "github.com/TudorHulban/authentication/fixtures/store-file-fixtures"
	storefile "github.com/TudorHulban/authentication/infra/stores/store-file"
	"github.com/TudorHulban/authentication/services/srender"
	"github.com/TudorHulban/authentication/services/ssessions"
	"github.com/TudorHulban/authentication/services/sticket"
	"github.com/TudorHulban/authentication/services/suser"
)

func InitializeApp(config *ParamsNewApp) (*App, error) {
	piers := PiersApp{
		ServiceUser: suser.NewService(
			storefile.NewStoreUser(&storefilefixtures.ParamsStoreFileUsers),
		),

		ServiceTicket: sticket.NewService(
			storefile.NewStoreTicket(&storefilefixtures.ParamsStoreFileTickets),
		),

		ServiceSessions: ssessions.NewService(),
	}

	serviceRender, errCr := srender.NewServiceRender(
		&srender.PiersServiceRender{
			ServiceUser:   piers.ServiceUser,
			ServiceTicket: piers.ServiceTicket,
		},
	)
	if errCr != nil {
		return nil, errCr
	}

	piers.ServiceRender = serviceRender

	app, errCr := NewApp(
		config,
		&piers,
	)
	if errCr != nil {
		log.Printf(
			"InitializeApp: %v",
			errCr,
		)

		os.Exit(
			apperrors.OSExitForApplicationIssues,
		)
	}

	return app,
		nil
}
