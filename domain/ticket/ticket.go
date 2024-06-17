package ticket

import (
	"encoding/json"
	"fmt"

	"github.com/TudorHulban/authentication/helpers"
)

type TicketMetadata struct {
	Status         TicketStatus       `json:",omitempty"`
	OpenedByUserID helpers.PrimaryKey `json:",omitempty"`
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

type ParamsTicketAsHTML struct {
	RouteTicket string
	Index       int
}

func (ti Ticket) AsHTMLTRow(params *ParamsTicketAsHTML) string {
	return fmt.Sprintf(
		`<tr><td>%d</td><td>%d</td><td><a href="%s/%d">%s</a><td>%s</td><td>%s</td><td>%s</td>`,

		params.Index,
		ti.PrimaryKey,
		params.RouteTicket,
		ti.PrimaryKey,
		ti.Name,
		ti.Status,
		ti.Name,
		helpers.UnixNanoTo(
			ti.UpdatedAt,
		),
	)

}
