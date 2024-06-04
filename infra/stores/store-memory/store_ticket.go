package storememory

import (
	"context"
	"fmt"
	"sync"

	"github.com/TudorHulban/authentication/domain/ticket"
)

// TODO: move to concurrent safe maps.
type StoreTicket struct {
	cacheTask  map[ticket.PrimaryKeyTicket]ticket.TicketInfo
	cacheEvent map[ticket.PrimaryKeyTicket][]*ticket.Event

	mu sync.RWMutex
}

func NewStoreTask() *StoreTicket {
	return &StoreTicket{
		cacheTask: make(
			map[ticket.PrimaryKeyTicket]ticket.TicketInfo,
		),

		cacheEvent: make(
			map[ticket.PrimaryKeyTicket][]*ticket.Event,
		),
	}
}

func (s *StoreTicket) CreateTicket(ctx context.Context, ticket *ticket.Ticket) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.cacheTask[ticket.PrimaryKeyTicket]; exists {
		return fmt.Errorf(
			"task with ID %d already exists",
			ticket.PrimaryKeyTicket,
		)
	}

	s.cacheTask[ticket.PrimaryKeyTicket] = ticket.TicketInfo

	return nil
}

func (s *StoreTicket) GetTicketByID(ctx context.Context, taskID ticket.PrimaryKeyTicket, result *ticket.TicketInfo) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	reconstructedTask, exists := s.cacheTask[taskID]
	if !exists {
		return fmt.Errorf(
			"task with ID %d not found",
			taskID,
		)
	}

	*result = reconstructedTask

	return nil
}

func (s *StoreTicket) SearchTasks(ctx context.Context, params *ticket.ParamsSearchTasks) (ticket.Tickets, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := make([]*ticket.Ticket, 0)

	for pk, taskInfo := range s.cacheTask {
		// TODO: apply filtering

		result = append(
			result,
			&ticket.Ticket{
				PrimaryKeyTicket: pk,
				TicketInfo:       taskInfo,
			},
		)
	}

	return result,
		nil
}

func (s *StoreTicket) UpdateTask(ctx context.Context, ticket *ticket.Ticket) {
	s.mu.Lock()

	s.cacheTask[ticket.PrimaryKeyTicket] = ticket.TicketInfo

	s.mu.Unlock()
}

func (s *StoreTicket) CloseTask(ctx context.Context, taskID ticket.PrimaryKeyTicket, status ticket.TicketStatus) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	cachedTicket, exists := s.cacheTask[taskID]
	if !exists {
		return fmt.Errorf(
			"task with ID %d not found",
			taskID,
		)
	}

	cachedTicket.Status = status

	s.UpdateTask(
		ctx,
		&ticket.Ticket{
			PrimaryKeyTicket: taskID,

			TicketInfo: cachedTicket,
		},
	)

	return nil
}

func (s *StoreTicket) AddEvent(ctx context.Context, ticketID ticket.PrimaryKeyTicket, event *ticket.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, existsTask := s.cacheTask[ticketID]
	if !existsTask {
		return fmt.Errorf(
			"task with ID %d not found",
			ticketID,
		)
	}

	s.cacheEvent[ticketID] = append(s.cacheEvent[ticketID], event)

	return nil
}

func (s *StoreTicket) GetEventsForTaskID(ctx context.Context, taskID ticket.PrimaryKeyTicket) ([]*ticket.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	cachedEvents, exists := s.cacheEvent[taskID]
	if !exists {
		return nil,
			fmt.Errorf(
				"task with ID %d not found",
				taskID,
			)
	}

	return cachedEvents,
		nil
}
