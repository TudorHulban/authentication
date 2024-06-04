package stask

import (
	"context"
	"strconv"
	"time"

	"github.com/TudorHulban/authentication/apperrors"
	"github.com/TudorHulban/authentication/domain/task"
	"github.com/TudorHulban/authentication/helpers"
	"github.com/TudorHulban/authentication/infra/stores"
	"github.com/TudorHulban/epochid"
	"github.com/asaskevich/govalidator"
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
	TaskName       string `valid:"required" json:"taskname"`
	OpenedByUserID uint
	TaskKind       task.TicketKind
}

func (s *Service) CreateTask(ctx context.Context, params *ParamsCreateTask) (helpers.PrimaryKey, error) {
	if _, errValidation := govalidator.ValidateStruct(params); errValidation != nil {
		return 0,
			apperrors.ErrValidation{
				Caller: "CreateTask",
				Issue:  errValidation,
			}
	}

	pk := task.PrimaryKeyTicket(
		epochid.NewIDIncremental10KWCoCorrection(),
	)

	if errCr := s.store.CreateTask(
		ctx,
		&task.Ticket{
			PrimaryKeyTicket: pk,

			TicketInfo: task.TicketInfo{
				Name: params.TaskName,

				TicketMetadata: task.TicketMetadata{
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

type ParamsGetTaskByID struct {
	TaskID       string
	UserLoggedID uint
}

func (s *Service) GetTaskByID(ctx context.Context, params *ParamsGetTaskByID) (*task.Ticket, error) {
	numericPK, errConv := strconv.ParseUint(params.TaskID, 10, 64)
	if errConv != nil {
		return nil,
			apperrors.ErrServiceValidation{
				Issue:  errConv,
				Caller: "GetTaskByID",
			}
	}

	var result task.TicketInfo

	if errGet := s.store.GetTaskByID(
		ctx,
		task.PrimaryKeyTicket(numericPK),
		&result,
	); errGet != nil {
		return nil,
			errGet
	}

	return &task.Ticket{
			PrimaryKeyTicket: task.PrimaryKeyTicket(numericPK),
			TicketInfo:       result,
		},
		nil
}

func (s *Service) SearchTasks(ctx context.Context, params *task.ParamsSearchTasks) (task.Tickets, error) {
	return s.store.SearchTasks(
		ctx,
		params,
	)
}

func (s *Service) CloseTask(ctx context.Context, taskID helpers.PrimaryKey, status task.TicketStatus) error {
	return s.store.CloseTask(
		ctx,
		task.PrimaryKeyTicket(taskID),
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
		task.PrimaryKeyTicket(taskID),
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
		task.PrimaryKeyTicket(taskID),
	)
}
