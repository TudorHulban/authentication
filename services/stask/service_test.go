package stask

import (
	"context"
	"testing"

	"github.com/TudorHulban/authentication/domain/task"
	storememory "github.com/TudorHulban/authentication/infra/stores/store-memory"
	"github.com/stretchr/testify/require"
)

func TestTask(t *testing.T) {

	service := NewService(
		storememory.NewStoreTask(),
	)

	t1 := task.TaskInfo{
		Name: "xxx",
	}

	ctx := context.Background()

	pkTask1, errCr := service.CreateTask(ctx, &t1)
	require.NoError(t, errCr)
	require.NotZero(t, pkTask1)

	reconstructedTask, errGet := service.GetTaskByID(ctx, pkTask1)
	require.NoError(t, errGet)
	require.NotNil(t, reconstructedTask)
	require.Equal(t,
		t1.Name,
		reconstructedTask.Name,
	)
	require.NotZero(t, reconstructedTask.TimestampOfLastUpdate, "timestamp")
	require.EqualValues(t,
		task.StatusNew,
		reconstructedTask.Status,
	)

	e1 := task.EventInfo{
		Content: "lorem ipsum 1",
	}

	require.NoError(t,
		service.AddEvent(
			ctx,
			pkTask1,
			&e1,
		),
	)

	e2 := task.EventInfo{
		Content: "lorem ipsum 2",
	}

	require.NoError(t,
		service.AddEvent(
			ctx,
			pkTask1,
			&e2,
		),
	)

	events, errGetEvents := service.GetEventsForTaskID(
		ctx,
		pkTask1,
	)
	require.NoError(t, errGetEvents)
	require.Len(t,
		events,
		2,
	)
	require.NotZero(t,
		events[0].TimestampOfAdd,
		"timestamp event",
	)
	require.Equal(t,
		e1.Content,
		events[0].Content,
		"content event 1",
	)
	require.Equal(t,
		e2.Content,
		events[1].Content,
		"content event 2",
	)
}
