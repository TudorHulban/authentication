package srender

import (
	"fmt"
	"os"
	"testing"

	storefile "github.com/TudorHulban/authentication/infra/stores/store-file"
	"github.com/TudorHulban/authentication/services/suser"
	"github.com/stretchr/testify/require"
)

func TestNewFormSearchTickets(t *testing.T) {
	p := ParamsNewFormSearchTickets{
		ActionButtonSearch: "/tickets",
		ActionButtonCreate: "/ticket",

		LabelButtonCreate: "Submit",
	}

	nameFile := fmt.Sprintf(
		"local_cache_user_%s_.json",
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

	serviceRender.NewFormSearchCreateTickets(&p).Render(os.Stdout)
}
