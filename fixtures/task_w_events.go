package fixtures

import (
	"context"

	"github.com/TudorHulban/authentication/domain/task"
	"github.com/TudorHulban/authentication/helpers"
	"github.com/TudorHulban/authentication/services/stask"
	"github.com/go-loremipsum/loremipsum"
)

type PiersFixtureTaskWEvents struct {
	ServiceTask *stask.Service
}

type ParamsFixtureTaskWEvents struct {
	TaskName           string
	TaskKind           task.TicketKind
	TaskOpenedByUserID uint

	NumberEvents uint
}

func FixtureTaskWEvents(ctx context.Context, piers *PiersFixtureTaskWEvents, params *ParamsFixtureTaskWEvents) (helpers.PrimaryKey, error) {
	idTask, errCr := piers.ServiceTask.CreateTask(
		ctx,
		&stask.ParamsCreateTask{
			TaskName: params.TaskName,
			TaskKind: params.TaskKind,

			OpenedByUserID: params.TaskOpenedByUserID,
		},
	)
	if errCr != nil {
		return helpers.PrimaryKeyZero,
			errCr
	}

	loremIpsumGenerator := loremipsum.New()

	for range params.NumberEvents {
		if errAddEvent := piers.ServiceTask.AddEvent(
			ctx,
			idTask,
			&stask.ParamsAddEvent{
				EventContent:   loremIpsumGenerator.Sentence(),
				OpenedByUserID: params.TaskOpenedByUserID,
			},
		); errAddEvent != nil {
			return helpers.PrimaryKeyZero,
				errAddEvent
		}
	}

	return idTask,
		nil
}
