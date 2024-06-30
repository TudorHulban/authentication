package sticket

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/TudorHulban/authentication/app/constants"
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
	TicketName     string             `valid:"required" json:"ticketname" form:"name"`
	OpenedByUserID helpers.PrimaryKey `valid:"required"`
	TicketKind     ticket.TicketKind
}

func NewParamsCreateTicketFromBytes(responseForm []byte) *ParamsCreateTicket {
	responseParams := helpers.ProcessFormURLEncoded(responseForm)

	var ticketName string

	if value, exists := responseParams["name"]; exists {
		ticketName = value
	}

	return &ParamsCreateTicket{
		TicketName: ticketName,
	}
}

func NewParamsCreateTicketFromMap(responseForm map[string]string) *ParamsCreateTicket {
	ticketName, exists := responseForm["name"]
	if exists {
		return &ParamsCreateTicket{
			TicketName: ticketName,
		}
	}

	return &ParamsCreateTicket{}
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

type ParamsGetTicketByIDString struct {
	TicketID     string
	UserLoggedID helpers.PrimaryKey
}

func (s *Service) GetTicketByIDString(ctx context.Context, params *ParamsGetTicketByIDString) (*ticket.Ticket, error) {
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

type ParamsGetTicketByIDNumeric struct {
	TicketID     helpers.PrimaryKey
	UserLoggedID helpers.PrimaryKey
}

// TODO: inject in string version
func (s *Service) GetTicketByIDNumeric(ctx context.Context, params *ParamsGetTicketByIDNumeric) (*ticket.Ticket, error) {
	var result ticket.TicketInfo

	if errGet := s.store.GetTicketByID(
		ctx,
		params.TicketID,
		&result,
	); errGet != nil {
		return nil,
			apperrors.ErrService{
				Issue:  errGet,
				Caller: "GetTicketByIDNumeric - s.store.GetTicketByID",
			}
	}

	return &ticket.Ticket{
			PrimaryKey: params.TicketID,
			TicketInfo: result,
		},
		nil
}

func (s *Service) SearchTickets(ctx context.Context, params *ticket.ParamsSearchTickets) (ticket.Tickets, error) {
	if params == nil {
		return s.store.SearchTickets(
			ctx,
			nil,
		)
	}

	var withID helpers.PrimaryKey

	if params.WithID.Valid {
		var errConv error

		withID, errConv = helpers.NewPrimaryKey(
			params.WithID.String,
		)
		if errConv != nil {
			return nil, errConv
		}
	}

	// var withStatus ticket.TicketStatus

	// if params.WithStatus.Valid {
	// 	var errConv error

	// 	withStatus, errConv = ticket.NewTicketStatus(params.WithStatus.String)
	// 	if errConv != nil {
	// 		return nil, errConv
	// 	}
	// }

	return s.store.SearchTickets(
		ctx,
		&paramsstores.ParamsSearchTickets{
			WithID: withID,
			// WithStatus: withStatus,
		},
	)
}

func (s *Service) SearchTicketEvents(ctx context.Context, params *ticket.ParamsSearchTicketEvents) (ticket.Events, error) {
	if params == nil {
		return s.store.SearchTicketEvents(
			ctx,
			nil,
		)
	}

	var withID helpers.PrimaryKey

	if params.WithTicketID.Valid {
		var errConv error

		withID, errConv = helpers.NewPrimaryKey(
			params.WithTicketID.String,
		)
		if errConv != nil {
			return nil, errConv
		}
	}

	return s.store.SearchTicketEvents(
		ctx,
		&paramsstores.ParamsSearchTicketEvents{
			WithTicketID: withID,
		},
	)
}

func (s *Service) CloseTicket(ctx context.Context, ticketID helpers.PrimaryKey) error {
	return s.store.CloseTicket(
		ctx,
		helpers.PrimaryKey(ticketID),
	)
}

type ParamsAddEvent struct {
	EventContent         string `json:"eventcontent,omitempty" valid:"required"`
	EventType            uint8  `json:"eventtype,omitempty" valid:"required"`
	ActualEventTypeLevel uint8  `json:"omitempty"`

	TicketID       helpers.PrimaryKey `json:"ticketid" valid:"required"`
	OpenedByUserID helpers.PrimaryKey `valid:"required"`
}

func NewParamsAddEvent(responseForm map[string]string) (*ParamsAddEvent, error) {
	var withTicketID helpers.PrimaryKey

	ticketID, exists := responseForm["ticketid"]
	if exists {
		var errConv error

		withTicketID, errConv = helpers.NewPrimaryKey(ticketID)
		if errConv != nil {
			return nil,
				apperrors.ErrValidation{
					Issue:  errConv,
					Caller: "NewParamsAddEvent",
				}
		}
	}

	var withEventContent string

	eventContent, exists := responseForm[strings.ToLower(
		constants.LabelTicketEventContent,
	)]
	if exists {
		withEventContent = eventContent
	}

	return &ParamsAddEvent{
			TicketID:     withTicketID,
			EventContent: withEventContent,
		},
		nil
}

func (s *Service) AddEvent(ctx context.Context, params *ParamsAddEvent) error {
	if _, errValidation := govalidator.ValidateStruct(params); errValidation != nil {
		return apperrors.ErrValidation{
			Caller: "AddEvent",
			Issue:  errValidation,
		}
	}

	reconstructedTicket, errGetTicket := s.GetTicketByIDNumeric(
		ctx,
		&ParamsGetTicketByIDNumeric{
			TicketID:     params.TicketID,
			UserLoggedID: params.OpenedByUserID,
		},
	)
	if errGetTicket != nil {
		return errGetTicket
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

			TicketEventType: ticket.TicketEventType{
				EventType: params.EventType,
				TicketEventTypeInfo: &ticket.TicketEventTypeInfo{
					ActualEventTypeLevel: helpers.Coalesce(
						ticket.EventType(params.ActualEventTypeLevel),
						ticket.TicketKindToEventType[reconstructedTicket.Kind][ticket.EventType(params.EventType)].
							DefaultEventTypeLevel,
					),
				},
			},
		},
	)
}

func (s *Service) GetEventsForTicketID(ctx context.Context, ticketID helpers.PrimaryKey) (ticket.Events, error) {
	return s.store.SearchTicketEvents(
		ctx,
		&paramsstores.ParamsSearchTicketEvents{
			WithTicketID: ticketID,
		},
	)
}
