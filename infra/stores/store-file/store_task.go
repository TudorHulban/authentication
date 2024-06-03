package storefile

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/TudorHulban/authentication/domain/task"
)

type StoreTask struct {
	pathCacheTask  string
	pathCacheEvent string

	mu sync.RWMutex
}

type ParamsNewStoreTask struct {
	PathCacheTask  string
	PathCacheEvent string
}

// TODO: make it portable
func NewStoreTask(params *ParamsNewStoreTask) *StoreTask {
	return &StoreTask{
		pathCacheTask:  params.PathCacheTask,
		pathCacheEvent: params.PathCacheEvent,
	}
}

func (s *StoreTask) readFileTask() (task.Tasks, error) {
	file, err := os.Open(s.pathCacheTask)
	if err != nil {
		if os.IsNotExist(err) {
			return nil,
				os.WriteFile(s.pathCacheTask, nil, 0644)
		}
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var result task.Tasks

	if len(bytes) > 0 {
		if err := json.Unmarshal(bytes, &result); err != nil {
			return nil, err
		}
	}
	return result, nil
}

func (s *StoreTask) writeFile(tasks task.Tasks) error {
	bytes, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.pathCacheTask, bytes, 0644)
}

func (s *StoreTask) CreateTask(ctx context.Context, task *task.Task) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	items, err := s.readFileTask()
	if err != nil {
		return err
	}

	_, errGetTaskByID := items.GetTaskByID(task.PrimaryKeyTask)
	if errGetTaskByID == nil {
		return nil
	}

	items = append(items, task)

	return s.writeFile(items)
}

func (s *StoreTask) GetTaskByID(ctx context.Context, taskID task.PrimaryKeyTask, result *task.TaskInfo) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	tasks, errGetTasks := s.readFileTask()
	if errGetTasks != nil {
		return errGetTasks
	}

	task, errGetTask := tasks.GetTaskByID(taskID)
	if errGetTask != nil {
		return errGetTask
	}

	*result = task.TaskInfo

	return nil
}

func (s *StoreTask) SearchTasks(ctx context.Context, params *task.ParamsSearchTasks) (task.Tasks, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.readFileTask()
}

func (s *StoreTask) UpdateTask(ctx context.Context, task *task.Task) {
	s.mu.Lock()
	defer s.mu.Unlock()

	items, err := s.readFileTask()
	if err != nil {
		fmt.Println(err)
	}

	for i, item := range items {
		if item.PrimaryKeyTask == task.PrimaryKeyTask {
			items[i] = task

			fmt.Println(s.writeFile(items))
		}
	}

	fmt.Printf("item with ID %v not found", task.PrimaryKeyTask)
}

func (s *StoreTask) CloseTask(ctx context.Context, taskID task.PrimaryKeyTask, status task.TaskStatus) error {
	return nil
}

func (s *StoreTask) AddEvent(ctx context.Context, taskID task.PrimaryKeyTask, event *task.Event) error {
	return nil
}

func (s *StoreTask) GetEventsForTaskID(ctx context.Context, taskID task.PrimaryKeyTask) ([]*task.Event, error) {
	return nil, nil
}
