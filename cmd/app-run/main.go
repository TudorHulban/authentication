package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/TudorHulban/authentication/app"
	"github.com/TudorHulban/authentication/apperrors"
	"github.com/TudorHulban/authentication/domain/ticket"
	"github.com/TudorHulban/authentication/fixtures"
)

func main() {
	app, errInitialize := app.InitializeApp(
		&configuration,
	)
	if errInitialize != nil {
		fmt.Println(errInitialize)

		os.Exit(
			apperrors.OSExitForApplicationIssues,
		)
	}

	ctx := context.Background()

	pkUser, errCr := fixtures.FixtureAddTestUser(
		ctx,
		&fixtures.PiersFixtureAddTestUser{
			ServiceUser: app.ServiceUser,
		},
	)
	if errCr != nil && !errors.As(errCr, &apperrors.ErrEntryAlreadyExists{}) {
		fmt.Println(errCr)

		os.Exit(
			apperrors.OSExitForApplicationIssues,
		)
	}

	if pkUser > 0 {
		_, errFixture := fixtures.FixtureTicketWEvents(
			context.Background(),
			&fixtures.PiersFixtureTicketWEvents{
				ServiceTicket: app.ServiceTicket,
			},
			&fixtures.ParamsFixtureTaskWEvents{
				TicketName: fmt.Sprintf(
					"Ticket %d%d",
					time.Now().Minute(),
					time.Now().Second(),
				),

				TicketKind:           ticket.KindTicket,
				TicketOpenedByUserID: pkUser,
				NumberEvents:         3,
			},
		)
		if errFixture != nil {
			fmt.Println(errFixture)

			os.Exit(
				apperrors.OSExitForFixtureIssues,
			)
		}
	}

	fmt.Println(
		app.Start(),
	)
}
