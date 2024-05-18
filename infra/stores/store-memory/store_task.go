package storememory

import (
	"context"
	"fmt"
	"sync"

	"github.com/TudorHulban/authentication/domain/task"
	"github.com/TudorHulban/authentication/helpers"
)

type StoreTask struct {
	cacheTask  map[helpers.PrimaryKey]task.TaskInfo
	cacheEvent map[helpers.PrimaryKey][]*task.Event // key is task pk

	mu sync.RWMutex
}

func NewStoreTask() *StoreTask {
	return &StoreTask{
		cacheTask: make(
			map[helpers.PrimaryKey]task.TaskInfo,
		),

		cacheEvent: make(
			map[helpers.PrimaryKey][]*task.Event,
		),
	}
}

func (s *StoreTask) CreateTask(ctx context.Context, task *task.Task) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.cacheTask[task.PrimaryKey]; exists {
		return fmt.Errorf(
			"task with ID %d already exists",
			task.PrimaryKey,
		)
	}

	s.cacheTask[task.PrimaryKey] = task.TaskInfo

	return nil
}

func (s *StoreTask) GetTaskByID(ctx context.Context, taskID helpers.PrimaryKey, result *task.TaskInfo) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	cachedTask, exists := s.cacheTask[helpers.PrimaryKey(taskID)]
	if !exists {
		return fmt.Errorf(
			"task with ID %d not found",
			taskID,
		)
	}

	*result = cachedTask

	return nil
}

func (s *StoreTask) UpdateTask(ctx context.Context, task *task.Task) {
	s.mu.Lock()

	s.cacheTask[task.PrimaryKey] = task.TaskInfo

	s.mu.Unlock()
}

func (s *StoreTask) CloseTask(ctx context.Context, taskID helpers.PrimaryKey, status task.TaskStatus) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	cachedTask, exists := s.cacheTask[helpers.PrimaryKey(taskID)]
	if !exists {
		return fmt.Errorf(
			"task with ID %d not found",
			taskID,
		)
	}

	cachedTask.Status = status

	s.UpdateTask(
		ctx,
		&task.Task{
			PrimaryKey: taskID,
			TaskInfo:   cachedTask,
		},
	)

	return nil
}

func (s *StoreTask) AddEvent(ctx context.Context, taskID helpers.PrimaryKey, event *task.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	cachedEvents, exists := s.cacheEvent[taskID]
	if !exists {
		return fmt.Errorf(
			"task with ID %d not found",
			taskID,
		)
	}

	cachedEvents = append(cachedEvents, event)

	s.cacheEvent[taskID] = cachedEvents

	return nil
}

func (s *StoreTask) GetEventsForTaskID(ctx context.Context, taskID helpers.PrimaryKey) ([]*task.Event, error) {
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
