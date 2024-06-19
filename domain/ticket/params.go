package ticket

import (
	"database/sql"

	"github.com/TudorHulban/authentication/helpers"
)

type ParamsSearchTickets struct {
	helpers.ParamsPagination

	WithID     sql.NullString `json:",omitempty" form:"id"`
	WithStatus string         `json:",omitempty" form:"status"`
	WithKind   string

	WithLastUpdateBefore string
	WithLastUpdatedAfter string

	WithOpenedByUserID uint
}

func NewParamsSearchTickets(responseForm []byte) *ParamsSearchTickets {
	responseParams := helpers.ProcessFormURLEncoded(responseForm)

	var withID sql.NullString

	if value, exists := responseParams["id"]; exists {
		withID = sql.NullString{
			Valid:  true,
			String: value,
		}
	}

	return &ParamsSearchTickets{
		WithID: withID,
	}
}
