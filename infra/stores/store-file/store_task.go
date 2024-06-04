package storefile

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/TudorHulban/authentication/domain/ticket"
)

type StoreTickets struct {
	pathCacheTicket string
	pathCacheEvent  string

	mu sync.RWMutex
}

type ParamsNewStoreTickets struct {
	PathCacheTickets string
	PathCacheEvent   string
}

// TODO: make it generic
func NewStoreTicket(params *ParamsNewStoreTickets) *StoreTickets {
	return &StoreTickets{
		pathCacheTicket: params.PathCacheTickets,
		pathCacheEvent:  params.PathCacheEvent,
	}
}

func (s *StoreTickets) readFile() (ticket.Tickets, error) {
	file, err := os.Open(s.pathCacheTicket)
	if err != nil {
		if os.IsNotExist(err) {
			return nil,
				os.WriteFile(s.pathCacheTicket, nil, 0644)
		}
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var result ticket.Tickets

	if len(bytes) > 0 {
		if err := json.Unmarshal(bytes, &result); err != nil {
			return nil, err
		}
	}
	return result, nil
}

func (s *StoreTickets) writeFile(tasks ticket.Tickets) error {
	bytes, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.pathCacheTicket, bytes, 0644)
}

func (s *StoreTickets) CreateTicket(ctx context.Context, ticket *ticket.Ticket) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	items, err := s.readFile()
	if err != nil {
		return err
	}

	_, errGetTaskByID := items.GetTaskByID(ticket.PrimaryKeyTicket)
	if errGetTaskByID == nil {
		return nil
	}

	items = append(items, ticket)

	return s.writeFile(items)
}

func (s *StoreTickets) GetTicketByID(ctx context.Context, taskID ticket.PrimaryKeyTicket, result *ticket.TicketInfo) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	tasks, errGetTasks := s.readFile()
	if errGetTasks != nil {
		return errGetTasks
	}

	ticket, errGetTask := tasks.GetTaskByID(taskID)
	if errGetTask != nil {
		return errGetTask
	}

	*result = ticket.TicketInfo

	return nil
}

func (s *StoreTickets) SearchTasks(ctx context.Context, params *ticket.ParamsSearchTasks) (ticket.Tickets, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.readFile()
}

func (s *StoreTickets) UpdateTask(ctx context.Context, ticket *ticket.Ticket) {
	s.mu.Lock()
	defer s.mu.Unlock()

	items, err := s.readFile()
	if err != nil {
		fmt.Println(err)
	}

	for i, item := range items {
		if item.PrimaryKeyTicket == ticket.PrimaryKeyTicket {
			items[i] = ticket

			fmt.Println(s.writeFile(items))
		}
	}

	fmt.Printf("item with ID %v not found", ticket.PrimaryKeyTicket)
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
