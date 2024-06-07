package app

import (
	"context"
	"log"
	"os"

	"github.com/TudorHulban/authentication/apperrors"
	"github.com/TudorHulban/authentication/domain/ticket"
	"github.com/TudorHulban/authentication/fixtures"
	storefilefixtures "github.com/TudorHulban/authentication/fixtures/store-file-fixtures"
	storefile "github.com/TudorHulban/authentication/infra/stores/store-file"
	"github.com/TudorHulban/authentication/services/ssessions"
	"github.com/TudorHulban/authentication/services/sticket"
	"github.com/TudorHulban/authentication/services/suser"
)

func InitializeApp(config *ParamsNewApp) *App {
	app, errCr := NewApp(
		config,
		&PiersApp{
			ServiceUser: suser.NewService(
				storefile.NewStoreUser(&storefilefixtures.ParamsStoreFileUsers),
			),

			ServiceTicket: sticket.NewService(
				storefile.NewStoreTicket(&storefilefixtures.ParamsStoreFileTickets),
			),

			ServiceSessions: ssessions.NewService(),
		},
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

	fixtures.FixtureTaskWEvents(
		context.Background(),
		&fixtures.PiersFixtureTicketWEvents{
			ServiceTicket: app.serviceTicket,
		},
		&fixtures.ParamsFixtureTaskWEvents{
			TicketName:           "Ticket 1",
			TicketKind:           ticket.KindUndefined,
			TicketOpenedByUserID: 1,
			NumberEvents:         10,
		},
	)

	return app
}
