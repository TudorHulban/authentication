package stask

import (
	"context"
	"fmt"
	"testing"

	"github.com/TudorHulban/authentication/domain/task"
	storememory "github.com/TudorHulban/authentication/infra/stores/store-memory"
	"github.com/stretchr/testify/require"
)

func TestTask(t *testing.T) {
	service := NewService(
		storememory.NewStoreTask(),
	)

	paramTask := ParamsCreateTask{
		TaskName:       "xxx",
		OpenedByUserID: 1,
		TaskKind:       task.KindUndefined,
	}

	ctx := context.Background()

	pkTask1, errCr := service.CreateTask(ctx, &paramTask)
	require.NoError(t, errCr)
	require.NotZero(t, pkTask1)

	reconstructedTask, errGet := service.GetTaskByID(ctx, pkTask1)
	require.NoError(t, errGet)
	require.NotNil(t, reconstructedTask)
	require.Equal(t,
		paramTask.TaskName,
		reconstructedTask.Name,
	)
	require.NotZero(t, reconstructedTask.TimestampOfLastUpdate, "timestamp")
	require.NotZero(t, reconstructedTask.OpenedByUserID)
	require.EqualValues(t,
		task.StatusNew,
		reconstructedTask.Status,
	)

	reconstructedTasks, errGetTasks := service.SearchTasks(ctx, &task.ParamsSearchTasks{})
	require.NoError(t, errGetTasks)
	require.NotZero(t, reconstructedTasks)

	fmt.Println(
		reconstructedTasks[0].PrimaryKeyTask,
		reconstructedTasks[0].TaskInfo,
	)

	e1 := task.EventInfo{
		Content: "lorem ipsum 1",
	}

	require.NoError(t,
		service.AddEvent(
			ctx,
			pkTask1,
			&ParamsAddEvent{
				EventContent:   e1.Content,
				OpenedByUserID: 1,
			},
		),
	)

	e2 := task.EventInfo{
		Content: "lorem ipsum 2",
	}

	require.NoError(t,
		service.AddEvent(
			ctx,
			pkTask1,
			&ParamsAddEvent{
				EventContent:   e2.Content,
				OpenedByUserID: 1,
			},
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
