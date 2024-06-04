package stores

import (
	"context"
	"fmt"
	"testing"

	"github.com/TudorHulban/authentication/domain/ticket"
	"github.com/TudorHulban/authentication/helpers"
	storefile "github.com/TudorHulban/authentication/infra/stores/store-file"
	"github.com/stretchr/testify/require"
)

func TestStoreTask(t *testing.T) {
	store := IStoreTicket(storefile.NewStoreTicket(&storefile.ParamsNewStoreTickets{}))

	t1 := ticket.Ticket{
		PrimaryKeyTicket: 1,

		TicketInfo: ticket.TicketInfo{
			Name: "t1",
		},
	}

	t2 := ticket.Ticket{
		PrimaryKeyTicket: 2,

		TicketInfo: ticket.TicketInfo{
			Name: "t2",
		},
	}

	ctx := context.Background()

	require.NoError(t,
		store.CreateTicket(ctx, &t1),
	)

	require.NoError(t,
		store.CreateTicket(ctx, &t2),
	)

	var reconstructedTaskInfo1 ticket.TicketInfo

	require.NoError(t,
		store.GetTicketByID(
			ctx,
			t1.PrimaryKeyTicket,
			&reconstructedTaskInfo1,
		),
	)
	require.NotZero(t, reconstructedTaskInfo1)

	tasks, erGetTasks := store.SearchTasks(
		ctx,
		&ticket.ParamsSearchTasks{
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
