package ticket

type TicketStatus uint8

const (
	StatusNew        = TicketStatus(0)
	StatusInProgress = TicketStatus(1)
	StatusClosed     = TicketStatus(2)
)

func (status TicketStatus) String() string {
	switch status {
	case 0:
		return "New"

	case 1:
		return "In Progress"

	case 2:
		return "Closed"

	default:
		return "Status Unknown"
	}
}
