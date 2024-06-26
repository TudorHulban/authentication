package ticket

import (
	"encoding/json"

	"github.com/TudorHulban/authentication/helpers"
)

type TicketMetadata struct {
	// Status         TicketStatus       `json:",omitempty"`
	OpenedByUserID helpers.PrimaryKey `json:",omitempty"`
	CurrentOwner   helpers.PrimaryKey `json:",omitempty"`
	Kind           TicketKind         `json:",omitempty"`
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

func NewTicketFrom(text string) (*Ticket, error) {
	var result Ticket

	if errUnmarshal := json.Unmarshal(
		[]byte(text),
		&result,
	); errUnmarshal != nil {
		return nil,
			errUnmarshal
	}

	return &result,
		nil
}
