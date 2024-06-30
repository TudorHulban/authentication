package srender

import (
	"fmt"
	"os"
	"testing"

	storefile "github.com/TudorHulban/authentication/infra/stores/store-file"
	"github.com/TudorHulban/authentication/services/sticket"
	"github.com/TudorHulban/authentication/services/suser"
	g "github.com/maragudk/gomponents"
	"github.com/stretchr/testify/require"
)

func TestBody(t *testing.T) {
	nameFileUsers := fmt.Sprintf(
		"local_users_%s_.json",
		t.Name(),
	)

	nameFileTickets := fmt.Sprintf(
		"local_tickets_%s_.json",
		t.Name(),
	)

	serviceUser := suser.NewService(
		storefile.NewStoreUser(
			&storefile.ParamsNewStoreUsers{
				PathCacheUsers: nameFileUsers,
			},
		),
	)

	serviceTicket := sticket.NewService(
		storefile.NewStoreTicket(
			&storefile.ParamsNewStoreTickets{
				PathCacheTickets: nameFileTickets,
			},
		),
	)

	serviceRender, errCr := NewServiceRender(
		&PiersServiceRender{
			ServiceUser:   serviceUser,
			ServiceTicket: serviceTicket,
		},
	)
	require.NoError(t, errCr)

	body := serviceRender.Body(
		&ParamsBody{
			Main: []g.Node{
				serviceRender.TableItemsHeadForTicketEvents("1"),
			},
		},
	)
	require.NotEmpty(t, body)

	rendered := RenderNodes(body...)

	fmt.Println(
		string(
			rendered,
		),
	)

	os.Remove(
		nameFileUsers,
	)
}
