package fixtures

import (
	"github.com/TudorHulban/authentication/domain/task"
	"github.com/TudorHulban/authentication/services/stask"
)

type PiersFixtureTaskWEvents struct {
	ServiceTask *stask.Service
}

type ParamsFixtureTaskWEvents struct {
	TaskName           string
	TaskKind           task.TaskKind
	TaskOpenedByUserID uint

	NumberEvents uint
}

// func FixtureTaskWEvents(ctx context.Context, piers *PiersFixtureTaskWEvents, params *ParamsFixtureTaskWEvents, t *testing.T) helpers.PrimaryKey {
// 	idTask, errCr := piers.ServiceTask.CreateTask(
// 		ctx,
// 		&stask.ParamsCreateTask{
// 			TaskName: params.TaskName,
// 			TaskKind: params.TaskKind,

// 			OpenedByUserID: params.TaskOpenedByUserID,
// 		},
// 	)
// 	require.NoError(t, errCr)

// 	for ix := range params.NumberEvents {

// 	}

// 	piers.ServiceTask.AddEvent()

// 	return helpers.PrimaryKeyZero
// }
