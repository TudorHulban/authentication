package ticket

import (
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

// func GetEventTypeFor(eventType string) EventType {
// 	var result EventType

// 	setEventTypeUS.Iter(
// 		func(k EventType, v string) (stop bool) {
// 			if v == eventType {
// 				result = k

// 				return true
// 			}

// 			return false
// 		},
// 	)

// 	return result
// }

func (ev EventType) String() string {
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
			Value: "WIP",
		},
		{
			Key:   3,
			Value: "Analysis",
		},
		{
			Key:   4,
			Value: "NI",
		},
		{
			Key:   5,
			Value: "WFI",
		},
		{
			Key:   6,
			Value: "Resolution",
		},
		{
			Key:   7,
			Value: "With 3rd party",
		},
		{
			Key:   8,
			Value: "Blocking",
		},
		{
			Key:   9,
			Value: "Unblocking",
		},
		{
			Key:   10,
			Value: "Escalation Internal",
		},
		{
			Key:   11,
			Value: "Escalation Customer",
		},
		{
			Key:   12,
			Value: "Close",
		},
	},
)

var SetEventType = setEventTypeUS

type ParamsGetNextEventTypes struct {
	EventKind TicketKind
	EventType string
}

func GetNextEventTypesFor(params *ParamsGetNextEventTypes) ([]string, error) {
	evType, errCr := NewEventType(params.EventType)
	if errCr != nil {
		return nil,
			errCr
	}

	nextEventTypes, exists := TicketKindToEventType[params.EventKind][evType]
	if !exists {
		return nil,
			apperrors.ErrInvalidInput{
				InputName: params.EventType,
			}
	}

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
