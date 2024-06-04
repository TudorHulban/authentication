package sticket

import (
	"context"
	"strconv"
	"time"

	"github.com/TudorHulban/authentication/apperrors"
	"github.com/TudorHulban/authentication/domain/ticket"
	"github.com/TudorHulban/authentication/helpers"
	"github.com/TudorHulban/authentication/infra/stores"
	"github.com/TudorHulban/epochid"
	"github.com/asaskevich/govalidator"
)

type Service struct {
	store stores.IStoreTicket
}

func NewService(store stores.IStoreTicket) *Service {
	return &Service{
		store: store,
	}
}

type ParamsCreateTicket struct {
	TicketName     string `valid:"required" json:"ticketname"`
	OpenedByUserID uint   //`valid:"required"`
	TicketKind     ticket.TicketKind
}

func (s *Service) CreateTicket(ctx context.Context, params *ParamsCreateTicket) (helpers.PrimaryKey, error) {
	if _, errValidation := govalidator.ValidateStruct(params); errValidation != nil {
		return 0,
			apperrors.ErrValidation{
				Caller: "CreateTicket",
				Issue:  errValidation,
			}
	}

	pk := ticket.PrimaryKeyTicket(
		epochid.NewIDIncremental10KWCoCorrection(),
	)

	if errCr := s.store.CreateTicket(
		ctx,
		&ticket.Ticket{
			PrimaryKeyTicket: pk,

			TicketInfo: ticket.TicketInfo{
				Name: params.TicketName,

				TicketMetadata: ticket.TicketMetadata{
					TimestampOfLastUpdate: time.Now().UnixNano(),
					Status:                ticket.StatusNew,
					OpenedByUserID:        params.OpenedByUserID,
					Kind:                  params.TicketKind,
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

type ParamsGetTicketByID struct {
	TicketID     string
	UserLoggedID uint
}

func (s *Service) GetTicketByID(ctx context.Context, params *ParamsGetTicketByID) (*ticket.Ticket, error) {
	numericPK, errConv := strconv.ParseUint(params.TicketID, 10, 64)
	if errConv != nil {
		return nil,
			apperrors.ErrServiceValidation{
				Issue:  errConv,
				Caller: "GetTicketByID",
			}
	}

	var result ticket.TicketInfo

	if errGet := s.store.GetTicketByID(
		ctx,
		ticket.PrimaryKeyTicket(numericPK),
		&result,
	); errGet != nil {
		return nil,
			errGet
	}

	return &ticket.Ticket{
			PrimaryKeyTicket: ticket.PrimaryKeyTicket(numericPK),
			TicketInfo:       result,
		},
		nil
}

func (s *Service) SearchTasks(ctx context.Context, params *ticket.ParamsSearchTasks) (ticket.Tickets, error) {
	return s.store.SearchTasks(
		ctx,
		params,
	)
}

func (s *Service) CloseTask(ctx context.Context, taskID helpers.PrimaryKey, status ticket.TicketStatus) error {
	return s.store.CloseTask(
		ctx,
		ticket.PrimaryKeyTicket(taskID),
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
		ticket.PrimaryKeyTicket(taskID),
		&ticket.Event{
			PrimaryKey: helpers.PrimaryKey(
				epochid.NewIDIncremental10KWCoCorrection(),
			),

			EventInfo: &ticket.EventInfo{
				Content:        params.EventContent,
				TimestampOfAdd: time.Now().UnixNano(),
				OpenedByUserID: params.OpenedByUserID,
			},
		},
	)
}

func (s *Service) GetEventsForTaskID(ctx context.Context, taskID helpers.PrimaryKey) ([]*ticket.Event, error) {
	return s.store.GetEventsForTaskID(
		ctx,
		ticket.PrimaryKeyTicket(taskID),
	)
}
