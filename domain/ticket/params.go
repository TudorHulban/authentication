package ticket

import (
	"database/sql"

	"github.com/TudorHulban/authentication/helpers"
)

type ParamsSearchTickets struct {
	helpers.ParamsPagination

	WithID     sql.NullString `json:",omitempty" form:"id"`
	WithStatus sql.NullString `json:",omitempty" form:"status"`
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

	var withStatus sql.NullString

	if value, exists := responseParams["status"]; exists {
		withStatus = sql.NullString{
			Valid:  true,
			String: value,
		}
	}

	return &ParamsSearchTickets{
		WithID:     withID,
		WithStatus: withStatus,
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

	var withStatus sql.NullString

	if value, exists := responseForm["status"]; exists {
		withStatus = sql.NullString{
			Valid:  true,
			String: value,
		}
	}

	return &ParamsSearchTickets{
		WithID:     withID,
		WithStatus: withStatus,
	}
}

type ParamsSearchTicketEvents struct {
	helpers.ParamsPagination

	WithTicketID sql.NullString `json:",omitempty" form:"ticketid"`

	WithLastUpdateBefore string
	WithLastUpdatedAfter string

	WithOpenedByUserID uint
}

func NewParamsSearchTicketEvents(responseForm []byte) *ParamsSearchTicketEvents {
	responseParams := helpers.ProcessFormURLEncoded(responseForm)

	var withID sql.NullString

	if value, exists := responseParams["id"]; exists {
		withID = sql.NullString{
			Valid:  true,
			String: value,
		}
	}

	return &ParamsSearchTicketEvents{
		WithTicketID: withID,
	}
}
