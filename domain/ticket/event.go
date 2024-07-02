package ticket

import (
	"fmt"
	"strings"

	"github.com/TudorHulban/authentication/helpers"
)

type EventInfo struct {
	Content        string
	TimestampOfAdd int64
	OpenedByUserID helpers.PrimaryKey
}

type TicketEventTypeInfo struct {
	DefaultEventTypeLevel EventLevel
	ActualEventTypeLevel  EventLevel
	Dependencies          []uint8
	AllowedNextEventTypes []EventType
}

type TicketEventType struct {
	EvType EventType

	*TicketEventTypeInfo
}

type Event struct {
	helpers.PrimaryKey

	TicketPK helpers.PrimaryKey

	TicketEventType
	*EventInfo
}

func (ev Event) String() string {
	result := []string{
		fmt.Sprintf(
			"Event %s belonging to ticket ID: %s",
			ev.PrimaryKey.String(),
			ev.TicketPK.String(),
		),
	}

	result = append(
		result,
		fmt.Sprintf(
			"Event type: %s",
			ev.EvType.String(),
		),
	)

	result = append(
		result,
		fmt.Sprintf(
			"Event opened by user ID: %s",
			ev.OpenedByUserID.String(),
		),
	)

	return strings.Join(result, "\n")
}

func GetIDEvent(item *Event) helpers.PrimaryKey {
	return item.PrimaryKey
}

func GetIDEventTicket(item *Event) uint64 {
	return uint64(item.TicketPK)
}

var CriteriaIDOfTicket = func(pk helpers.PrimaryKey) func(item *Event) bool {
	return func(item *Event) bool {
		return GetIDEvent(item) == pk
	}
}

var CriteriaEventsWithTicketID = func(pk helpers.PrimaryKey) func(item *Event) bool {
	return func(item *Event) bool {
		return GetIDEventTicket(item) == uint64(pk)
	}
}
