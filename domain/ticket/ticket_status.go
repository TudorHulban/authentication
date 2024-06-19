package ticket

import "errors"

type TicketStatus uint8

const (
	StatusNew        = TicketStatus(0)
	StatusInProgress = TicketStatus(1)
	StatusClosed     = TicketStatus(2)
)

const (
	StringStatusNew        = "New"
	StringStatusInProgress = "In Progress"
	StringStatusClosed     = "Closed"
)

const msgUnknownStatus = "unknown status"

func NewTicketStatus(value string) (TicketStatus, error) {
	switch value {
	case StringStatusNew:
		return StatusNew, nil

	case StringStatusInProgress:
		return StatusInProgress, nil

	case StringStatusClosed:
		return StatusClosed, nil

	default:
		return 0,
			errors.New(msgUnknownStatus)
	}
}

func (status TicketStatus) String() string {
	switch status {
	case 0:
		return StringStatusNew

	case 1:
		return "In Progress"

	case 2:
		return "Closed"

	default:
		return "Status Unknown"
	}
}
