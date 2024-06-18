package sticket

import (
	"context"
	"strconv"
	"time"

	"github.com/TudorHulban/authentication/apperrors"
	"github.com/TudorHulban/authentication/domain/ticket"
	"github.com/TudorHulban/authentication/helpers"
	"github.com/TudorHulban/authentication/infra/stores"
	paramsstores "github.com/TudorHulban/authentication/infra/stores/params-stores"
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
	TicketName     string             `valid:"required" json:"ticketname"`
	OpenedByUserID helpers.PrimaryKey `valid:"required"`
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

	pk := helpers.PrimaryKey(
		epochid.NewIDIncremental10KWCoCorrection(),
	)

	var timestamp helpers.Timestamp

	if errCr := s.store.CreateTicket(
		ctx,
		&ticket.Ticket{
			PrimaryKey: pk,

			TicketInfo: ticket.TicketInfo{
				Name: params.TicketName,

				TicketMetadata: ticket.TicketMetadata{
					Status:         ticket.StatusNew,
					OpenedByUserID: params.OpenedByUserID,
					Kind:           params.TicketKind,
				},

				Timestamp: timestamp.WithCreateNow(),
			},
		},
	); errCr != nil {
		return helpers.PrimaryKeyZero,
			apperrors.ErrService{
				Issue:  errCr,
				Caller: "CreateTicket",
			}
	}

	return helpers.PrimaryKey(pk),
		nil
}

type ParamsGetTicketByID struct {
	TicketID     string
	UserLoggedID helpers.PrimaryKey
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
		helpers.PrimaryKey(numericPK),
		&result,
	); errGet != nil {
		return nil,
			apperrors.ErrService{
				Issue:  errGet,
				Caller: "GetTicketByID - s.store.GetTicketByID",
			}
	}

	return &ticket.Ticket{
			PrimaryKey: helpers.PrimaryKey(numericPK),
			TicketInfo: result,
		},
		nil
}

func (s *Service) SearchTickets(ctx context.Context, params *ticket.ParamsSearchTickets) (ticket.Tickets, error) {
	var withID int

	if len(params.WithID) > 0 {
		var errConv error

		withID, errConv = strconv.Atoi(
			params.WithID,
		)
		if errConv != nil {
			return nil, errConv
		}
	}

	return s.store.SearchTickets(
		ctx,
		&paramsstores.ParamsSearchTickets{
			WithID: helpers.PrimaryKey(withID),
		},
	)
}

func (s *Service) CloseTicket(ctx context.Context, taskID helpers.PrimaryKey, status ticket.TicketStatus) error {
	return s.store.CloseTicket(
		ctx,
		helpers.PrimaryKey(taskID),
		status,
	)
}

type ParamsAddEvent struct {
	EventContent string `json:"eventcontent" valid:"required"`

	TicketID       helpers.PrimaryKey `json:"ticketid" valid:"required"`
	OpenedByUserID helpers.PrimaryKey `valid:"required"`
}

func (s *Service) AddEvent(ctx context.Context, params *ParamsAddEvent) error {
	if _, errValidation := govalidator.ValidateStruct(params); errValidation != nil {
		return apperrors.ErrValidation{
			Caller: "AddEvent",
			Issue:  errValidation,
		}
	}

	return s.store.AddEvent(
		ctx,
		helpers.PrimaryKey(params.TicketID),
		&ticket.Event{
			PrimaryKey: helpers.PrimaryKey(
				epochid.NewIDIncremental10KWCoCorrection(),
			),

			TicketPK: params.TicketID,

			EventInfo: &ticket.EventInfo{
				Content:        params.EventContent,
				TimestampOfAdd: time.Now().UnixNano(),
				OpenedByUserID: params.OpenedByUserID,
			},
		},
	)
}

func (s *Service) GetEventsForTicketID(ctx context.Context, ticketID helpers.PrimaryKey) ([]*ticket.Event, error) {
	return s.store.GetEventsForTicketID(
		ctx,
		helpers.PrimaryKey(ticketID),
	)
}
