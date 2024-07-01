package ticket

import "github.com/TudorHulban/authentication/helpers"

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
