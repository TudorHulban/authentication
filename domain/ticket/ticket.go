package ticket

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
	PrimaryKeyTicket

	TicketInfo
}

func GetID(t *Ticket) uint64 {
	return uint64(t.PrimaryKeyTicket)
}

var CriteriaPK = func(pk PrimaryKeyTicket) func(item *Ticket) bool {
	return func(item *Ticket) bool {
		return GetID(item) == uint64(pk)
	}
}
