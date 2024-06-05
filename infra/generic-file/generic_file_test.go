package genericfile

import (
	"testing"

	"github.com/TudorHulban/authentication/domain/ticket"
	"github.com/stretchr/testify/require"
)

func TestGenericStore(t *testing.T) {
	store := NewGenericStoreFile[ticket.Ticket](
		&ParamsNewGenericStoreFile{
			PathStoreFile: ".local_ticket.json",
		},
	)
	require.NotNil(t, store)

	t1 := ticket.Ticket{
		PrimaryKeyTicket: ticket.PrimaryKeyTicket(1),

		TicketInfo: ticket.TicketInfo{
			Name: "T1",
		},
	}

	store.createFirstItem(&t1)

	criteria := func(item *ticket.Ticket) bool {
		return ticket.GetID(item) == uint64(t1.PrimaryKeyTicket)
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
			uint64(t1.PrimaryKeyTicket),
			&t1,
			ticket.GetID,
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
		store.DeleteItem(uint64(t1.PrimaryKeyTicket), ticket.GetID),
	)

	recontructedItems3, errGet3 := store.SearchItems(criteria)
	require.NoError(t, errGet3)
	require.Empty(t, recontructedItems3)
}
