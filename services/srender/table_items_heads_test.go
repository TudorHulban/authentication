package srender

import (
	"fmt"
	"os"
	"testing"

	storefile "github.com/TudorHulban/authentication/infra/stores/store-file"
	"github.com/TudorHulban/authentication/services/suser"
	"github.com/stretchr/testify/require"
)

func TestTableTicketsHead(t *testing.T) {
	nameFile := fmt.Sprintf(
		"local_%s_.json",
		t.Name(),
	)

	serviceUser := suser.NewService(
		storefile.NewStoreUser(
			&storefile.ParamsNewStoreUsers{
				PathCacheUsers: nameFile,
			},
		),
	)

	serviceRender, errCr := NewServiceRender(
		&PiersServiceRender{
			ServiceUser: serviceUser,
		},
	)
	require.NoError(t, errCr)

	serviceRender.
		TableItemsHeadForTickets("1").
		Render(os.Stdout)

	os.Remove(
		nameFile,
	)
}
