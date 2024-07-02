package ticket

import (
	"fmt"

	"github.com/TudorHulban/authentication/apperrors"
	"github.com/TudorHulban/authentication/helpers"
)

type (
	EventType  uint8
	EventLevel uint8
)

func NewEventType(eventType string) (EventType, error) {
	var result EventType

	var entryFound bool

	setEventTypeUS.Iter(
		func(k EventType, v string) (stop bool) {
			if v == eventType {
				result = k

				entryFound = true

				return true
			}

			return false
		},
	)

	if entryFound {
		return result,
			nil
	}

	return 0,
		apperrors.ErrEntryNotFound{
			Key: eventType,
		}
}

func (ev EventType) String() string {
	fmt.Println("EventType", ev) // TODO: crashing

	value, exists := setEventTypeUS.Get(ev)
	if !exists {
		return msgUnknownEventType
	}

	return value
}

// for all ticket types
var setEventTypeUS = helpers.NewImmutableSetFrom[EventType, string](
	[]helpers.KV[EventType, string]{
		{
			Key:   1,
			Value: "Open",
		},
		{
			Key:   2,
			Value: "AssignTo",
		},
		{
			Key:   3,
			Value: "WIP",
		},
		{
			Key:   4,
			Value: "Analysis",
		},
		{
			Key:   5,
			Value: "NI",
		},
		{
			Key:   6,
			Value: "WFI",
		},
		{
			Key:   7,
			Value: "Resolution",
		},
		{
			Key:   8,
			Value: "With 3rd party",
		},
		{
			Key:   9,
			Value: "Blocking",
		},
		{
			Key:   10,
			Value: "Unblocking",
		},
		{
			Key:   11,
			Value: "Escalation Internal",
		},
		{
			Key:   12,
			Value: "Escalation Customer",
		},
		{
			Key:   13,
			Value: "Close",
		},
	},
)

var SetEventType = setEventTypeUS

type ParamsGetNextEventTypes struct {
	TicketKind TicketKind
	EventType  EventType
}

func GetNextEventTypesFor(params *ParamsGetNextEventTypes) ([]string, error) {

	ticketEventTypeInfo, existsTicketKind := TicketKindToEventType[params.TicketKind]
	if !existsTicketKind {
		return nil,
			apperrors.ErrInvalidInput{
				InputName:  "params.TicketKind",
				InputValue: params.TicketKind,
			}
	}

	fmt.Println("xxxxxxxxxxxxxxxxxxxxxx", params.TicketKind, params.EventType)

	nextEventTypes, existsTicketEventType := ticketEventTypeInfo[params.EventType]
	if !existsTicketEventType {
		return nil,
			apperrors.ErrInvalidInput{
				InputName:  "params.EventType",
				InputValue: params.EventType,
			}
	}

	fmt.Println("yyyyyyyyyyyyyyyyyyyyy")

	result := make(
		[]string,
		len(nextEventTypes.AllowedNextEventTypes),
		len(nextEventTypes.AllowedNextEventTypes),
	)

	for ix, value := range nextEventTypes.AllowedNextEventTypes {
		result[ix] = value.String()
	}

	return result, nil
}
