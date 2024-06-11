package suser

import (
	"context"
	"os"
	"testing"

	testuser "github.com/TudorHulban/authentication/fixtures/test-user"
	storefile "github.com/TudorHulban/authentication/infra/stores/store-file"
	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	nameFile := ".local_test_users.json"

	service := NewService(
		storefile.NewStoreUser(
			&storefile.ParamsNewStoreUsers{
				PathCacheUsers: nameFile,
			},
		),
	)

	ctx := context.Background()

	p1 := ParamsCreateUser{
		Email:    testuser.TestUser.Email,
		Password: testuser.TestUser.Password,
		Name:     testuser.TestUser.Name,
	}

	require.NoError(t,
		service.CreateUser(ctx, &p1),
	)

	require.Error(t,
		service.CreateUser(ctx, &p1),
	)

	require.NoError(t,
		os.Remove(nameFile),
	)
}
