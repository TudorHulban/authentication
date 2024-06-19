package ticket

import (
	"github.com/TudorHulban/authentication/helpers"
)

type ParamsSearchTickets struct {
	helpers.ParamsPagination

	WithID     *string `json:",omitempty" form:"id"`
	WithStatus string  `json:",omitempty" form:"status"`
	WithKind   string

	WithLastUpdateBefore string
	WithLastUpdatedAfter string

	WithOpenedByUserID uint
}
