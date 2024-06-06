package ticket

import "github.com/TudorHulban/authentication/helpers"

type TicketMetadata struct {
	TimestampOfLastUpdate int64
	Status                TicketStatus
	OpenedByUserID        uint
	Kind                  TicketKind
}

type TicketInfo struct {
	Name string

	TicketMetadata
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
