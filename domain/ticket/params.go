package ticket

import (
	"database/sql"

	"github.com/TudorHulban/authentication/helpers"
)

type ParamsSearchTickets struct {
	helpers.ParamsPagination

	WithID     sql.NullString `json:",omitempty" form:"id"`
	WithStatus uint8          `json:",omitempty" form:"status"`
	WithKind   string

	WithLastUpdateBefore string
	WithLastUpdatedAfter string

	WithOpenedByUserID uint
}

func NewParamsSearchTicketsFromBytes(responseForm []byte) *ParamsSearchTickets {
	responseParams := helpers.ProcessFormURLEncoded(responseForm)

	var withID sql.NullString

	if value, exists := responseParams["id"]; exists {
		withID = sql.NullString{
			Valid:  true,
			String: value,
		}
	}

	// var withStatus sql.NullInt16

	// if value, exists := responseParams["status"]; exists {
	// 	withStatus = sql.NullInt16{
	// 		Valid:  true,
	// 		String: value,
	// 	}
	// }

	return &ParamsSearchTickets{
		WithID: withID,
		// WithStatus: withStatus,
	}
}

// TODO: move to service.
func NewParamsSearchTicketsFromMap(responseForm map[string]string) *ParamsSearchTickets {
	var withID sql.NullString

	if value, exists := responseForm["id"]; exists {
		withID = sql.NullString{
			Valid:  true,
			String: value,
		}
	}

	// var withStatus sql.NullString

	// if value, exists := responseForm["status"]; exists {
	// 	withStatus = sql.NullString{
	// 		Valid:  true,
	// 		String: value,
	// 	}
	// }

	return &ParamsSearchTickets{
		WithID: withID,
		// WithStatus: withStatus,
	}
}

type ParamsSearchTicketEvents struct {
	helpers.ParamsPagination

	WithTicketID sql.NullString `json:",omitempty" form:"ticketid"`

	WithLastUpdateBefore string
	WithLastUpdatedAfter string

	WithOpenedByUserID uint
}

func NewParamsSearchTicketEventsFromBytes(responseForm []byte) *ParamsSearchTicketEvents {
	responseParams := helpers.ProcessFormURLEncoded(responseForm)

	var withTicketID sql.NullString

	if value, exists := responseParams["ticketid"]; exists {
		withTicketID = sql.NullString{
			Valid:  true,
			String: value,
		}
	}

	return &ParamsSearchTicketEvents{
		WithTicketID: withTicketID,
	}
}

func NewParamsSearchTicketEventsFromMap(responseForm map[string]string) *ParamsSearchTicketEvents {
	var withTicketID sql.NullString

	if value, exists := responseForm["ticketid"]; exists {
		withTicketID = sql.NullString{
			Valid:  true,
			String: value,
		}
	}

	return &ParamsSearchTicketEvents{
		WithTicketID: withTicketID,
	}
}

func NewParamsSearchTicketFromMap(responseForm map[string]string) *ParamsSearchTicketEvents {
	var withTicketID sql.NullString

	if value, exists := responseForm["ticketid"]; exists {
		withTicketID = sql.NullString{
			Valid:  true,
			String: value,
		}
	}

	return &ParamsSearchTicketEvents{
		WithTicketID: withTicketID,
	}
}
