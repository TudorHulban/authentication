package storefile

import (
	"context"

	"github.com/TudorHulban/authentication/apperrors"
	"github.com/TudorHulban/authentication/domain/ticket"
	"github.com/TudorHulban/authentication/helpers"
	genericfile "github.com/TudorHulban/authentication/infra/generic-file"
	paramsstores "github.com/TudorHulban/authentication/infra/stores/params-stores"
)

type StoreTickets struct {
	storeTickets      *genericfile.GenericStoreFile[ticket.Ticket]
	storeTicketEvents *genericfile.GenericStoreFile[ticket.Event]
}

type ParamsNewStoreTickets struct {
	PathCacheTickets string
	PathCacheEvents  string
}

func NewStoreTicket(params *ParamsNewStoreTickets) *StoreTickets {
	return &StoreTickets{
		storeTickets: genericfile.
			NewGenericStoreFile[ticket.Ticket](
			&genericfile.ParamsNewGenericStoreFile{
				PathStoreFile: params.PathCacheTickets,
			},
		),

		storeTicketEvents: genericfile.
			NewGenericStoreFile[ticket.Event](
			&genericfile.ParamsNewGenericStoreFile{
				PathStoreFile: params.PathCacheEvents,
			},
		),
	}
}

func (s *StoreTickets) CreateTicket(ctx context.Context, item *ticket.Ticket, force ...bool) error {
	return s.storeTickets.CreateItem(item, ticket.GetIDTicket, force...)
}

func (s *StoreTickets) GetTicketByID(ctx context.Context, ticketID helpers.PrimaryKey, result *ticket.TicketInfo) error {
	reconstructedItem, errGet := s.storeTickets.SearchItem(ticket.CriteriaPK(ticketID))
	if errGet != nil {
		return apperrors.ErrInfrastructure{
			Issue:              errGet,
			NameInfrastructure: "StoreTickets",
			Caller:             "GetTicketByID",
		}
	}

	*result = reconstructedItem.TicketInfo

	return nil
}

func (s *StoreTickets) SearchTickets(ctx context.Context, params *paramsstores.ParamsSearchTickets) (ticket.Tickets, error) {
	if params.WithID == helpers.PrimaryKeyZero {
		return s.storeTickets.SearchItems(
			helpers.CriteriaTrue,
		)
	}

	return s.storeTickets.SearchItems(
		ticket.CriteriaPK(
			params.WithID,
		),
	)
}

func (s *StoreTickets) UpdateTicket(ctx context.Context, item *ticket.Ticket) error {
	return s.storeTickets.UpdateItem(uint64(item.PrimaryKey), item, ticket.GetIDTicket)
}

func (s *StoreTickets) CloseTicket(ctx context.Context, ticketID helpers.PrimaryKey, status ticket.TicketStatus) error {
	return nil
}

func (s *StoreTickets) AddEvent(ctx context.Context, ticketID helpers.PrimaryKey, event *ticket.Event) error {
	return s.storeTicketEvents.CreateItem(event, ticket.GetIDEvent)
}

func (s *StoreTickets) GetEventsForTicketID(ctx context.Context, ticketID helpers.PrimaryKey) ([]*ticket.Event, error) {
	return s.storeTicketEvents.SearchItems(
		ticket.CriteriaEventsOfTicket(ticketID),
	)
}
