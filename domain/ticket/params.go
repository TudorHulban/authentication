package ticket

import (
	"github.com/TudorHulban/authentication/helpers"
)

type ParamsSearchTickets struct {
	helpers.ParamsPagination

	WithID     string `json:"id,omitempty"`
	WithStatus string
	WithKind   string

	WithLastUpdateBefore string
	WithLastUpdatedAfter string

	WithOpenedByUserID uint
}
