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

func TestStoreTicket(t *testing.T) {
	store := IStoreTicket(
		storefile.NewStoreTicket(
			&storefile.ParamsNewStoreTickets{
				PathCacheTickets: ".local_test_cache_tickets.json",
				PathCacheEvents:  ".local_test_cache_events.json",
			},
		),
	)

	t1 := ticket.Ticket{
		PrimaryKey: 1,

		TicketInfo: ticket.TicketInfo{
			Name: "t1",
		},
	}

	t2 := ticket.Ticket{
		PrimaryKey: 2,

		TicketInfo: ticket.TicketInfo{
			Name: "t2",
		},
	}

	ctx := context.Background()

	require.NoError(t,
		store.CreateTicket(ctx, &t1, true),
	)

	require.NoError(t,
		store.CreateTicket(ctx, &t2, true),
	)

	var reconstructedTaskInfo1 ticket.TicketInfo

	require.NoError(t,
		store.GetTicketByID(
			ctx,
			t1.PrimaryKey,
			&reconstructedTaskInfo1,
		),
	)
	require.NotZero(t, reconstructedTaskInfo1)

	tasks, erGetTasks := store.SearchTickets(
		ctx,
		&ticket.ParamsSearchTickets{
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
