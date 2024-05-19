package storememory

import (
	"context"
	"fmt"
	"sync"

	"github.com/TudorHulban/authentication/domain/task"
)

// TODO: move to concurrent safe maps.
type StoreTask struct {
	cacheTask  map[task.PrimaryKeyTask]*task.TaskInfo
	cacheEvent map[task.PrimaryKeyTask][]*task.Event

	mu sync.RWMutex
}

func NewStoreTask() *StoreTask {
	return &StoreTask{
		cacheTask: make(
			map[task.PrimaryKeyTask]*task.TaskInfo,
		),

		cacheEvent: make(
			map[task.PrimaryKeyTask][]*task.Event,
		),
	}
}

func (s *StoreTask) CreateTask(ctx context.Context, task *task.Task) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.cacheTask[task.PrimaryKeyTask]; exists {
		return fmt.Errorf(
			"task with ID %d already exists",
			task.PrimaryKeyTask,
		)
	}

	s.cacheTask[task.PrimaryKeyTask] = task.TaskInfo

	return nil
}

func (s *StoreTask) GetTaskByID(ctx context.Context, taskID task.PrimaryKeyTask, result *task.TaskInfo) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	reconstructedTask, exists := s.cacheTask[taskID]
	if !exists {
		return fmt.Errorf(
			"task with ID %d not found",
			taskID,
		)
	}

	*result = *reconstructedTask

	return nil
}

func (s *StoreTask) SearchTasks(ctx context.Context, params *task.ParamsSearchTasks) ([]*task.Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	result := make([]*task.Task, 0)

	for pk, taskInfo := range s.cacheTask {
		// TODO: apply filtering

		result = append(
			result,
			&task.Task{
				PrimaryKeyTask: pk,
				TaskInfo:       taskInfo,
			},
		)
	}

	return result,
		nil
}

func (s *StoreTask) UpdateTask(ctx context.Context, task *task.Task) {
	s.mu.Lock()

	s.cacheTask[task.PrimaryKeyTask] = task.TaskInfo

	s.mu.Unlock()
}

func (s *StoreTask) CloseTask(ctx context.Context, taskID task.PrimaryKeyTask, status task.TaskStatus) error {
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
		&task.Task{
			PrimaryKeyTask: taskID,

			TaskInfo: cachedTask,
		},
	)

	return nil
}

func (s *StoreTask) AddEvent(ctx context.Context, taskID task.PrimaryKeyTask, event *task.Event) error {
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

func (s *StoreTask) GetEventsForTaskID(ctx context.Context, taskID task.PrimaryKeyTask) ([]*task.Event, error) {
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
