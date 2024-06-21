package paramsstores

import (
	"github.com/TudorHulban/authentication/domain/ticket"
	"github.com/TudorHulban/authentication/helpers"
)

type ParamsSearchTickets struct {
	helpers.ParamsPagination

	WithID     helpers.PrimaryKey
	WithStatus ticket.TicketStatus
	WithKind   ticket.TicketKind

	WithLastUpdateBefore int64
	WithLastUpdatedAfter int64

	WithOpenedByUserID uint
}

type ParamsSearchTicketEvents struct {
	helpers.ParamsPagination

	WithTicketID helpers.PrimaryKey

	WithLastUpdateBefore int64
	WithLastUpdatedAfter int64

	WithOpenedByUserID uint
}
