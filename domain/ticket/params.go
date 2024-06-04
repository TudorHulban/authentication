package ticket

import "github.com/TudorHulban/authentication/helpers"

type ParamsSearchTasks struct {
	helpers.ParamsPagination

	WithStatus TicketStatus
	WithKind   TicketKind

	WithLastUpdateBefore int64
	WithLastUpdatedAfter int64

	WithOpenedByUserID uint
}
