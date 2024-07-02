package sticket

import (
	"context"
	"fmt"

	"github.com/TudorHulban/authentication/domain/ticket"
	"github.com/TudorHulban/authentication/helpers"
)

func (s *Service) GetTicketPossibleEventTypes(ctx context.Context, item *ticket.Ticket) ([]string, error) {
	reconstructedTicketEvents, errGetEvents := s.GetEventsForTicketID(
		ctx,
		item.PrimaryKey,
	)
	if errGetEvents != nil {
		return nil,
			errGetEvents
	}

	fmt.Printf("%#v", reconstructedTicketEvents[2].String())

	return ticket.GetNextEventTypesFor(
		&ticket.ParamsGetNextEventTypes{
			TicketKind: item.Kind,
			EventType:  helpers.High(reconstructedTicketEvents).EvType,
		},
	)
}
