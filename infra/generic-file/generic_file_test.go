package genericfile

import (
	"fmt"
	"os"
	"testing"

	"github.com/TudorHulban/authentication/apperrors"
	"github.com/TudorHulban/authentication/domain/ticket"
	"github.com/TudorHulban/authentication/helpers"
	"github.com/stretchr/testify/require"
)

func TestCriteriaPK(t *testing.T) {
	nameFileTest := fmt.Sprintf(
		".local_test_%s.json",
		t.Name(),
	)

	store := NewGenericStoreFile[ticket.Ticket](
		&ParamsNewGenericStoreFile{
			PathStoreFile: nameFileTest,
		},
	)
	require.NotNil(t, store)

	recontructedNoEntries, errGetNoEntries := store.SearchItems(
		ticket.CriteriaPK(
			helpers.PrimaryKey(1),
		),
	)
	require.ErrorAs(t,
		errGetNoEntries,
		&apperrors.ErrNoEntriesFound{},
	)
	require.Empty(t, recontructedNoEntries)

	t1 := ticket.Ticket{
		PrimaryKey: helpers.PrimaryKey(1),

		TicketInfo: ticket.TicketInfo{
			Name: "T1",
		},
	}

	require.NoError(t,
		store.CreateFirstItem(&t1),
	)

	t2 := ticket.Ticket{
		PrimaryKey: helpers.PrimaryKey(2),

		TicketInfo: ticket.TicketInfo{
			Name: "T2",
		},
	}

	require.NoError(t,
		store.CreateItem(
			&t2,
			ticket.GetIDTicket,
		),
	)

	recontructedByPK, errGetByPK := store.SearchItems(
		ticket.CriteriaPK(
			helpers.PrimaryKey(1),
		),
	)
	require.NoError(t, errGetByPK)
	require.NotEmpty(t, recontructedByPK)
	require.Equal(t,
		t1.PrimaryKey,
		recontructedByPK[0].PrimaryKey,
	)

	require.NoError(t,
		os.Remove(nameFileTest),
	)
}

func TestGenericStore(t *testing.T) {
	nameFileTickets := ".local_test_tickets.json"

	store := NewGenericStoreFile[ticket.Ticket](
		&ParamsNewGenericStoreFile{
			PathStoreFile: nameFileTickets,
		},
	)
	require.NotNil(t, store)

	t1 := ticket.Ticket{
		PrimaryKey: helpers.PrimaryKey(1),

		TicketInfo: ticket.TicketInfo{
			Name: "T1",
		},
	}

	require.NoError(t,
		store.CreateFirstItem(&t1),
	)

	require.Error(t,
		store.CreateItem(&t1, ticket.GetIDTicket),
	)

	criteria := func(item *ticket.Ticket) bool {
		return ticket.GetIDTicket(item) == t1.PrimaryKey
	}

	recontructedItems1, errGet1 := store.SearchItems(criteria)
	require.NoError(t, errGet1)
	require.NotEmpty(t, recontructedItems1)
	require.Len(t, recontructedItems1, 1)
	require.Equal(t,
		t1.Name,
		recontructedItems1[0].Name,
	)

	t1.Name = "T1 updated"

	require.NoError(t,
		store.UpdateItem(
			t1.PrimaryKey,
			&t1,
			ticket.GetIDTicket,
		),
	)

	recontructedItems2, errGet2 := store.SearchItems(criteria)
	require.NoError(t, errGet2)
	require.NotEmpty(t, recontructedItems2)
	require.Len(t, recontructedItems2, 1)
	require.Equal(t,
		t1.Name,
		recontructedItems2[0].Name,
	)

	require.NoError(t,
		store.DeleteItem(t1.PrimaryKey, ticket.GetIDTicket),
	)

	recontructedItems3, errGet3 := store.SearchItems(criteria)
	require.Error(t, errGet3)
	require.Empty(t, recontructedItems3)

	require.NoError(t,
		os.Remove(nameFileTickets),
	)
}
