package fixtures

import (
	"context"
	"fmt"

	"github.com/TudorHulban/authentication/domain/ticket"
	"github.com/TudorHulban/authentication/helpers"
	"github.com/TudorHulban/authentication/services/sticket"
	"github.com/go-loremipsum/loremipsum"
)

type PiersFixtureTicketWEvents struct {
	ServiceTicket *sticket.Service
}

type ParamsFixtureTaskWEvents struct {
	TicketName           string
	TicketKind           ticket.TicketKind
	TicketOpenedByUserID helpers.PrimaryKey

	NumberEvents uint
}

func FixtureTicketWEvents(ctx context.Context, piers *PiersFixtureTicketWEvents, params *ParamsFixtureTaskWEvents) (helpers.PrimaryKey, error) {
	idTicket, errCr := piers.ServiceTicket.CreateTicket(
		ctx,
		&sticket.ParamsCreateTicket{
			TicketName: params.TicketName,
			TicketKind: params.TicketKind,

			OpenedByUserID: params.TicketOpenedByUserID,
		},
	)
	if errCr != nil {
		return helpers.PrimaryKeyZero,
			errCr
	}

	loremIpsumGenerator := loremipsum.New()

	for ix := range params.NumberEvents {
		if errAddEvent := piers.ServiceTicket.AddEvent(
			ctx,

			&sticket.ParamsAddEvent{
				EventContent: fmt.Sprintf(
					"Event %d\n%s",
					ix+1,
					loremIpsumGenerator.Sentence(),
				),
				OpenedByUserID: params.TicketOpenedByUserID,

				TicketID: idTicket,
			},
		); errAddEvent != nil {
			return helpers.PrimaryKeyZero,
				errAddEvent
		}
	}

	return idTicket,
		nil
}
