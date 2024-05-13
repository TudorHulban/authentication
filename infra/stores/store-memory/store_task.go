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
	cacheEvent map[helpers.PrimaryKey]task.EventInfo

	mu sync.RWMutex
}

func NewStoreTask() *StoreTask {
	return &StoreTask{
		cacheTask: make(
			map[helpers.PrimaryKey]task.TaskInfo,
		),

		cacheEvent: make(
			map[helpers.PrimaryKey]task.EventInfo,
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

func (s *StoreTask) GetTaskByID(ctx context.Context, taskID uint, result *task.TaskInfo) error {
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
