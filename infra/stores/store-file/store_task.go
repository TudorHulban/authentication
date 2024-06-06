package storefile

import (
	"context"

	"github.com/TudorHulban/authentication/domain/ticket"
	"github.com/TudorHulban/authentication/helpers"
	genericfile "github.com/TudorHulban/authentication/infra/generic-file"
)

type StoreTickets struct {
	storeIickets      *genericfile.GenericStoreFile[ticket.Ticket]
	storeTicketEvents *genericfile.GenericStoreFile[ticket.Event]
}

type ParamsNewStoreTickets struct {
	PathCacheTickets string
	PathCacheEvent   string
}

func NewStoreTicket(params *ParamsNewStoreTickets) *StoreTickets {
	return &StoreTickets{
		storeIickets: genericfile.
			NewGenericStoreFile[ticket.Ticket](
			&genericfile.ParamsNewGenericStoreFile{
				PathStoreFile: params.PathCacheTickets,
			},
		),

		storeTicketEvents: genericfile.
			NewGenericStoreFile[ticket.Event](
			&genericfile.ParamsNewGenericStoreFile{
				PathStoreFile: params.PathCacheEvent,
			},
		),
	}
}

func (s *StoreTickets) CreateTicket(ctx context.Context, item *ticket.Ticket) error {
	return s.storeIickets.CreateItem(item, ticket.GetID)
}

func (s *StoreTickets) GetTicketByID(ctx context.Context, taskID ticket.PrimaryKeyTicket, result *ticket.TicketInfo) error {
	reconstructedItem, errGet := s.storeIickets.SearchItem(ticket.CriteriaPK(taskID))
	if errGet != nil {
		return errGet
	}

	*result = reconstructedItem.TicketInfo

	return nil
}

func (s *StoreTickets) SearchTasks(ctx context.Context, params *ticket.ParamsSearchTasks) (ticket.Tickets, error) {
	return s.storeIickets.SearchItems(helpers.CriteriaTrue)
}

func (s *StoreTickets) UpdateTask(ctx context.Context, item *ticket.Ticket) error {
	return s.storeIickets.UpdateItem(uint64(item.PrimaryKeyTicket), item, ticket.GetID)
}

func (s *StoreTickets) CloseTask(ctx context.Context, taskID ticket.PrimaryKeyTicket, status ticket.TicketStatus) error {
	return nil
}

func (s *StoreTickets) AddEvent(ctx context.Context, taskID ticket.PrimaryKeyTicket, event *ticket.Event) error {
	return nil
}

func (s *StoreTickets) GetEventsForTaskID(ctx context.Context, taskID ticket.PrimaryKeyTicket) ([]*ticket.Event, error) {
	return nil, nil
}
