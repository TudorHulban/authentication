package sticket

import (
	"context"

	appuser "github.com/TudorHulban/authentication/domain/app-user"
	"github.com/TudorHulban/authentication/domain/ticket"
	"github.com/TudorHulban/authentication/helpers"
)

type ParamsGetTicketStatus struct {
	TicketID helpers.PrimaryKey
	UserInfo appuser.UserInfo
}

func (s *Service) GetTicketStatus(ctx context.Context, params *ParamsGetTicketStatus) (ticket.EventType, error) {
	reconstructedTicketEvents, errGetEvents := s.GetEventsForTicketID(
		ctx,
		params.TicketID,
	)
	if errGetEvents != nil {
		return 0,
			errGetEvents
	}

	lastVisibleEvent, errGetEvent := reconstructedTicketEvents.GetLastEventFor(
		params.UserInfo.Level,
	)
	if errGetEvent != nil {
		return 0,
			errGetEvent
	}

	return lastVisibleEvent.ActualEventTypeLevel,
		nil
}
