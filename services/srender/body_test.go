package srender

import (
	"fmt"
	"os"
	"testing"

	storefile "github.com/TudorHulban/authentication/infra/stores/store-file"
	"github.com/TudorHulban/authentication/services/suser"
	g "github.com/maragudk/gomponents"
	"github.com/stretchr/testify/require"
)

func TestBody(t *testing.T) {
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
		nameFile,
	)
}
