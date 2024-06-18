package stores

import (
	"context"
	"os"
	"testing"

	"github.com/TudorHulban/authentication/domain/ticket"
	"github.com/TudorHulban/authentication/helpers"
	paramsstores "github.com/TudorHulban/authentication/infra/stores/params-stores"
	storefile "github.com/TudorHulban/authentication/infra/stores/store-file"
	"github.com/stretchr/testify/require"
)

const (
	_nameFileTickets = ".local_test_cache_tickets.json"
	_nameFileEvents  = ".local_test_cache_events.json"
)

func TestStoreTicket(t *testing.T) {
	store := IStoreTicket(
		storefile.NewStoreTicket(
			&storefile.ParamsNewStoreTickets{
				PathCacheTickets: ".local_test_cache_tickets.json",
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

	require.Error(t,
		store.CreateTicket(ctx, &t1),
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
		&paramsstores.ParamsSearchTickets{
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

	require.NoError(t,
		os.Remove(_nameFileTickets),
	)
}
