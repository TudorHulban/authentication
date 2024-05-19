package task

import "github.com/TudorHulban/authentication/helpers"

type ParamsSearchTasks struct {
	helpers.ParamsPagination

	WithStatus TaskStatus
	WithKind   TaskKind

	WithLastUpdateBefore int64
	WithLastUpdatedAfter int64

	WithOpenedByUserID uint
}
