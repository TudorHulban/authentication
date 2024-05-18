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

func (s *Service) CreateTask(ctx context.Context, item *task.TaskInfo) (helpers.PrimaryKey, error) {
	pk := helpers.PrimaryKey(
		epochid.NewIDIncremental10KWCoCorrection(),
	)

	if errCr := s.store.CreateTask(
		ctx,
		&task.Task{
			PrimaryKey: pk,
			TaskInfo: &task.TaskInfo{
				Name:           item.Name,
				OpenedByUserID: item.OpenedByUserID,
				Kind:           item.Kind,

				TaskMetadata: &task.TaskMetadata{
					TimestampOfLastUpdate: time.Now().UnixNano(),
					Status:                task.StatusNew,
				},
			},
		},
	); errCr != nil {
		return helpers.PrimaryKeyZero,
			errCr
	}

	return pk,
		nil
}

func (s *Service) GetTaskByID(ctx context.Context, taskID helpers.PrimaryKey) (*task.TaskInfo, error) {
	var result task.TaskInfo

	if errGet := s.store.GetTaskByID(
		ctx,
		taskID,
		&result,
	); errGet != nil {
		return nil,
			errGet
	}

	return &result,
		nil
}

func (s *Service) CloseTask(ctx context.Context, taskID helpers.PrimaryKey, status task.TaskStatus) error {
	return s.store.CloseTask(
		ctx,
		taskID,
		status,
	)
}

func (s *Service) AddEvent(ctx context.Context, taskID helpers.PrimaryKey, event *task.EventInfo) error {
	return s.store.AddEvent(
		ctx,
		taskID,
		&task.Event{
			PrimaryKey: helpers.PrimaryKey(
				epochid.NewIDIncremental10KWCoCorrection(),
			),
			EventInfo: &task.EventInfo{
				Content:        event.Content,
				TimestampOfAdd: time.Now().UnixNano(),
			},
		},
	)
}

func (s *Service) GetEventsForTaskID(ctx context.Context, taskID helpers.PrimaryKey) ([]*task.Event, error) {
	return s.store.GetEventsForTaskID(
		ctx,
		taskID,
	)
}
