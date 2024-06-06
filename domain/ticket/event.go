package ticket

import "github.com/TudorHulban/authentication/helpers"

type EventInfo struct {
	Content        string
	TimestampOfAdd int64
	OpenedByUserID uint
}

type Event struct {
	helpers.PrimaryKey

	TicketPK helpers.PrimaryKey

	*EventInfo
}

func GetIDEvent(item *Event) uint64 {
	return uint64(item.PrimaryKey)
}

var CriteriaEventsOfTicket = func(pk helpers.PrimaryKey) func(item *Ticket) bool {
	return func(item *Ticket) bool {
		return GetIDTicket(item) == uint64(pk)
	}
}
