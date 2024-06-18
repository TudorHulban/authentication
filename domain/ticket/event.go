package ticket

import "github.com/TudorHulban/authentication/helpers"

type EventInfo struct {
	Content        string
	TimestampOfAdd int64
	OpenedByUserID helpers.PrimaryKey
}

type Event struct {
	helpers.PrimaryKey

	TicketPK helpers.PrimaryKey

	*EventInfo
}

func GetIDEvent(item *Event) uint64 {
	return uint64(item.PrimaryKey)
}

func GetIDEventTicket(item *Event) uint64 {
	return uint64(item.TicketPK)
}

var CriteriaIDOfTicket = func(pk helpers.PrimaryKey) func(item *Event) bool {
	return func(item *Event) bool {
		return GetIDEvent(item) == uint64(pk)
	}
}

var CriteriaEventsOfTicket = func(pk helpers.PrimaryKey) func(item *Event) bool {
	return func(item *Event) bool {
		return GetIDEventTicket(item) == uint64(pk)
	}
}
