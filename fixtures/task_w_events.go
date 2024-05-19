package fixtures

import (
	"context"
	"testing"

	"github.com/TudorHulban/authentication/domain/task"
	"github.com/TudorHulban/authentication/helpers"
	"github.com/TudorHulban/authentication/services/stask"
	"github.com/go-loremipsum/loremipsum"
	"github.com/stretchr/testify/require"
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

func FixtureTaskWEvents(ctx context.Context, piers *PiersFixtureTaskWEvents, params *ParamsFixtureTaskWEvents, t *testing.T) helpers.PrimaryKey {
	idTask, errCr := piers.ServiceTask.CreateTask(
		ctx,
		&stask.ParamsCreateTask{
			TaskName: params.TaskName,
			TaskKind: params.TaskKind,

			OpenedByUserID: params.TaskOpenedByUserID,
		},
	)
	require.NoError(t, errCr)

	loremIpsumGenerator := loremipsum.New()

	for range params.NumberEvents {
		require.NoError(t,
			piers.ServiceTask.AddEvent(
				ctx,
				idTask,
				&stask.ParamsAddEvent{
					EventContent:   loremIpsumGenerator.Sentence(),
					OpenedByUserID: params.TaskOpenedByUserID,
				},
			),
		)
	}

	return idTask
}
