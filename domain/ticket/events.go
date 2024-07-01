package ticket

import (
	"github.com/TudorHulban/authentication/apperrors"
	"github.com/TudorHulban/authentication/helpers"
)

type Events []*Event

func (evs Events) GetLastEventFor(level EventLevel) (*Event, error) {
	for ix := len(evs) - 1; ix >= 0; ix-- {
		if helpers.Sanitize(evs[ix].TicketEventTypeInfo).ActualEventTypeLevel >= level {
			return evs[ix],
				nil
		}
	}

	return nil,
		apperrors.ErrEntryNotFound{
			Key: level,
		}
}
