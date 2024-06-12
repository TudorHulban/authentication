package sticket

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/TudorHulban/authentication/apperrors"
	"github.com/TudorHulban/authentication/domain/ticket"
	"github.com/TudorHulban/authentication/helpers"
	storefile "github.com/TudorHulban/authentication/infra/stores/store-file"
	"github.com/stretchr/testify/require"
)

const (
	_nameFileTickets = ".local_test_tickets.json"
	_nameFileEvents  = ".local_test_events.json"
)

func TestErrorsTicket(t *testing.T) {
	service := NewService(
		storefile.NewStoreTicket(
			&storefile.ParamsNewStoreTickets{
				PathCacheTickets: _nameFileTickets,
				PathCacheEvents:  _nameFileEvents,
			},
		),
	)

	ctx := context.Background()

	_, errGetNonExistentTicketID := service.GetTicketByID(
		ctx,
		&ParamsGetTicketByID{
			TicketID:     "1",
			UserLoggedID: 1,
		},
	)
	require.Error(t, errGetNonExistentTicketID)

	events, errGetEventsNonExistentTicketID := service.GetEventsForTicketID(
		ctx,
		helpers.PrimaryKeyZero,
	)
	require.ErrorAs(t, errGetEventsNonExistentTicketID, &apperrors.ErrNoEntriesFound{})
	require.Empty(t, events)

	require.NoError(t,
		os.Remove(_nameFileTickets),
	)

	require.NoError(t,
		os.Remove(_nameFileEvents),
	)
}

func TestTicket(t *testing.T) {
	service := NewService(
		storefile.NewStoreTicket(
			&storefile.ParamsNewStoreTickets{
				PathCacheTickets: _nameFileTickets,
				PathCacheEvents:  _nameFileEvents,
			},
		),
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
			TicketID:     pkTask1.String(),
			UserLoggedID: 1,
		},
	)
	require.NoError(t, errGet)
	require.NotNil(t, reconstructedTicket)
	require.Equal(t,
		paramTicket.TicketName,
		reconstructedTicket.Name,
	)
	require.NotZero(t, reconstructedTicket.CreatedAt, "created at timestamp")
	require.NotZero(t, reconstructedTicket.OpenedByUserID)
	require.EqualValues(t,
		ticket.StatusNew,
		reconstructedTicket.Status,
	)

	reconstructedTickets, errGetTasks := service.SearchTickets(ctx, &ticket.ParamsSearchTickets{})
	require.NoError(t, errGetTasks)
	require.NotZero(t, reconstructedTickets)

	fmt.Println(
		reconstructedTickets[0].PrimaryKey,
		"\n",
		reconstructedTickets[0].TicketInfo,
	)

	e1 := ticket.EventInfo{
		Content: "lorem ipsum 1",
	}

	require.NoError(t,
		service.AddEvent(
			ctx,

			&ParamsAddEvent{
				EventContent:   e1.Content,
				OpenedByUserID: 1,

				TicketID: pkTask1,
			},
		),
	)

	e2 := ticket.EventInfo{
		Content: "lorem ipsum 2",
	}

	require.NoError(t,
		service.AddEvent(
			ctx,

			&ParamsAddEvent{
				EventContent:   e2.Content,
				OpenedByUserID: 1,

				TicketID: pkTask1,
			},
		),
	)

	events, errGetEvents := service.GetEventsForTicketID(
		ctx,
		pkTask1,
	)
	require.NoError(t, errGetEvents)
	require.GreaterOrEqual(t,
		len(events),
		2,
	)
	require.NotZero(t,
		events[0].TimestampOfAdd,
		"timestamp event",
	)
	require.NotZero(t,
		events[0].TicketPK,
		"ticket PK event",
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

	require.NoError(t,
		os.Remove(_nameFileTickets),
	)

	require.NoError(t,
		os.Remove(_nameFileEvents),
	)
}
