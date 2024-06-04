package sticket

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/TudorHulban/authentication/domain/ticket"
	storememory "github.com/TudorHulban/authentication/infra/stores/store-memory"
	"github.com/stretchr/testify/require"
)

func TestTask(t *testing.T) {
	service := NewService(
		storememory.NewStoreTask(),
	)

	paramTicket := ParamsCreateTicket{
		TicketName:     "xxx",
		OpenedByUserID: 1,
		TicketKind:     ticket.KindUndefined,
	}

	ctx := context.Background()

	pkTask1, errCr := service.CreateTicket(ctx, &paramTicket)
	require.NoError(t, errCr)
	require.NotZero(t, pkTask1)

	reconstructedTicket, errGet := service.GetTicketByID(
		ctx,
		&ParamsGetTicketByID{
			TicketID:     strconv.Itoa(int(pkTask1)),
			UserLoggedID: 1,
		},
	)
	require.NoError(t, errGet)
	require.NotNil(t, reconstructedTicket)
	require.Equal(t,
		paramTicket.TicketName,
		reconstructedTicket.Name,
	)
	require.NotZero(t, reconstructedTicket.TimestampOfLastUpdate, "timestamp")
	require.NotZero(t, reconstructedTicket.OpenedByUserID)
	require.EqualValues(t,
		ticket.StatusNew,
		reconstructedTicket.Status,
	)

	reconstructedTasks, errGetTasks := service.SearchTasks(ctx, &ticket.ParamsSearchTasks{})
	require.NoError(t, errGetTasks)
	require.NotZero(t, reconstructedTasks)

	fmt.Println(
		reconstructedTasks[0].PrimaryKeyTicket,
		reconstructedTasks[0].TicketInfo,
	)

	e1 := ticket.EventInfo{
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

	e2 := ticket.EventInfo{
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
