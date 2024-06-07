package ticket

import "github.com/TudorHulban/authentication/helpers"

type TicketMetadata struct {
	Status         TicketStatus
	OpenedByUserID helpers.PrimaryKey
	Kind           TicketKind
}

type TicketInfo struct {
	Name string

	TicketMetadata
	helpers.Timestamp
}

type Ticket struct {
	helpers.PrimaryKey

	TicketInfo
}

func GetIDTicket(item *Ticket) uint64 {
	return uint64(item.PrimaryKey)
}

var CriteriaPK = func(pk helpers.PrimaryKey) func(item *Ticket) bool {
	return func(item *Ticket) bool {
		return GetIDTicket(item) == uint64(pk)
	}
}
