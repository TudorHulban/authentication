package ticket

type TicketStatus uint8

const (
	StatusNew        = TicketStatus(0)
	StatusInProgress = TicketStatus(1)
	StatusClosed     = TicketStatus(2)
)
