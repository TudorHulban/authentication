package stask

import (
	"context"
	"time"

	"github.com/TudorHulban/authentication/domain/task"
	"github.com/TudorHulban/authentication/helpers"
	"github.com/TudorHulban/authentication/infra/stores"
	"github.com/TudorHulban/epochid"
)

type Service struct {
	store stores.IStoreTask
}

func NewService(store stores.IStoreTask) *Service {
	return &Service{
		store: store,
	}
}

type ParamsCreateTask struct {
	TaskName       string
	OpenedByUserID uint
	TaskKind       task.TaskKind
}

func (s *Service) CreateTask(ctx context.Context, params *ParamsCreateTask) (helpers.PrimaryKey, error) {
	pk := task.PrimaryKeyTask(
		epochid.NewIDIncremental10KWCoCorrection(),
	)

	if errCr := s.store.CreateTask(
		ctx,
		&task.Task{
			PrimaryKeyTask: pk,

			TaskInfo: &task.TaskInfo{
				Name: params.TaskName,

				TaskMetadata: &task.TaskMetadata{
					TimestampOfLastUpdate: time.Now().UnixNano(),
					Status:                task.StatusNew,
					OpenedByUserID:        params.OpenedByUserID,
					Kind:                  params.TaskKind,
				},
			},
		},
	); errCr != nil {
		return helpers.PrimaryKeyZero,
			errCr
	}

	return helpers.PrimaryKey(pk),
		nil
}

func (s *Service) GetTaskByID(ctx context.Context, taskID helpers.PrimaryKey) (*task.TaskInfo, error) {
	var result task.TaskInfo

	if errGet := s.store.GetTaskByID(
		ctx,
		task.PrimaryKeyTask(taskID),
		&result,
	); errGet != nil {
		return nil,
			errGet
	}

	return &result,
		nil
}

func (s *Service) SearchTasks(ctx context.Context, params *task.ParamsSearchTasks) ([]*task.Task, error) {
	return s.store.SearchTasks(
		ctx,
		params,
	)
}

func (s *Service) CloseTask(ctx context.Context, taskID helpers.PrimaryKey, status task.TaskStatus) error {
	return s.store.CloseTask(
		ctx,
		task.PrimaryKeyTask(taskID),
		status,
	)
}

type ParamsAddEvent struct {
	EventContent   string
	OpenedByUserID uint
}

func (s *Service) AddEvent(ctx context.Context, taskID helpers.PrimaryKey, params *ParamsAddEvent) error {
	return s.store.AddEvent(
		ctx,
		task.PrimaryKeyTask(taskID),
		&task.Event{
			PrimaryKey: helpers.PrimaryKey(
				epochid.NewIDIncremental10KWCoCorrection(),
			),

			EventInfo: &task.EventInfo{
				Content:        params.EventContent,
				TimestampOfAdd: time.Now().UnixNano(),
				OpenedByUserID: params.OpenedByUserID,
			},
		},
	)
}

func (s *Service) GetEventsForTaskID(ctx context.Context, taskID helpers.PrimaryKey) ([]*task.Event, error) {
	return s.store.GetEventsForTaskID(
		ctx,
		task.PrimaryKeyTask(taskID),
	)
}
