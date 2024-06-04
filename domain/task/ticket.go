package task

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
