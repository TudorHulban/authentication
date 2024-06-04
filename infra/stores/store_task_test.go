package stores

import (
	"context"
	"fmt"
	"testing"

	"github.com/TudorHulban/authentication/domain/task"
	"github.com/TudorHulban/authentication/helpers"
	storefile "github.com/TudorHulban/authentication/infra/stores/store-file"
	"github.com/stretchr/testify/require"
)

func TestStoreTask(t *testing.T) {
	store := IStoreTask(storefile.NewStoreTask(&storefile.ParamsNewStoreTask{}))

	t1 := task.Ticket{
		PrimaryKeyTicket: 1,

		TicketInfo: task.TicketInfo{
			Name: "t1",
		},
	}

	t2 := task.Ticket{
		PrimaryKeyTicket: 2,

		TicketInfo: task.TicketInfo{
			Name: "t2",
		},
	}

	ctx := context.Background()

	require.NoError(t,
		store.CreateTask(ctx, &t1),
	)

	require.NoError(t,
		store.CreateTask(ctx, &t2),
	)

	var reconstructedTaskInfo1 task.TicketInfo

	require.NoError(t,
		store.GetTaskByID(
			ctx,
			t1.PrimaryKeyTicket,
			&reconstructedTaskInfo1,
		),
	)
	require.NotZero(t, reconstructedTaskInfo1)

	tasks, erGetTasks := store.SearchTasks(
		ctx,
		&task.ParamsSearchTasks{
			ParamsPagination: helpers.ParamsPagination{
				First: 10,
			},
		},
	)
	require.NoError(t, erGetTasks)
	require.NotEmpty(t, tasks)
	require.Len(t, tasks, 2)
	require.NotEqual(t,
		tasks[0].Name,
		tasks[1].Name,
	)

	fmt.Println(tasks)
}
