package storememory

import (
	"context"
	"fmt"
	"sync"

	"github.com/TudorHulban/authentication/domain/task"
)

// TODO: move to concurrent safe maps.
type StoreTask struct {
	cacheTask  map[task.PrimaryKeyTicket]task.TicketInfo
	cacheEvent map[task.PrimaryKeyTicket][]*task.Event

	mu sync.RWMutex
}

func NewStoreTask() *StoreTask {
	return &StoreTask{
		cacheTask: make(
			map[task.PrimaryKeyTicket]task.TicketInfo,
		),

		cacheEvent: make(
			map[task.PrimaryKeyTicket][]*task.Event,
		),
	}
}

func (s *StoreTask) CreateTask(ctx context.Context, task *task.Ticket) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.cacheTask[task.PrimaryKeyTicket]; exists {
		return fmt.Errorf(
			"task with ID %d already exists",
			task.PrimaryKeyTicket,
		)
	}

	s.cacheTask[task.PrimaryKeyTicket] = task.TicketInfo

	return nil
}

func (s *StoreTask) GetTaskByID(ctx context.Context, taskID task.PrimaryKeyTicket, result *task.TicketInfo) error {
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

func (s *StoreTask) SearchTasks(ctx context.Context, params *task.ParamsSearchTasks) (task.Tickets, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := make([]*task.Ticket, 0)

	for pk, taskInfo := range s.cacheTask {
		// TODO: apply filtering

		result = append(
			result,
			&task.Ticket{
				PrimaryKeyTicket: pk,
				TicketInfo:       taskInfo,
			},
		)
	}

	return result,
		nil
}

func (s *StoreTask) UpdateTask(ctx context.Context, task *task.Ticket) {
	s.mu.Lock()

	s.cacheTask[task.PrimaryKeyTicket] = task.TicketInfo

	s.mu.Unlock()
}

func (s *StoreTask) CloseTask(ctx context.Context, taskID task.PrimaryKeyTicket, status task.TicketStatus) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	cachedTask, exists := s.cacheTask[taskID]
	if !exists {
		return fmt.Errorf(
			"task with ID %d not found",
			taskID,
		)
	}

	cachedTask.Status = status

	s.UpdateTask(
		ctx,
		&task.Ticket{
			PrimaryKeyTicket: taskID,

			TicketInfo: cachedTask,
		},
	)

	return nil
}

func (s *StoreTask) AddEvent(ctx context.Context, taskID task.PrimaryKeyTicket, event *task.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, existsTask := s.cacheTask[taskID]
	if !existsTask {
		return fmt.Errorf(
			"task with ID %d not found",
			taskID,
		)
	}

	s.cacheEvent[taskID] = append(s.cacheEvent[taskID], event)

	return nil
}

func (s *StoreTask) GetEventsForTaskID(ctx context.Context, taskID task.PrimaryKeyTicket) ([]*task.Event, error) {
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
