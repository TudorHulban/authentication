package suser

import (
	"context"
	"os"
	"testing"

	testuser "github.com/TudorHulban/authentication/fixtures/test-user"
	"github.com/TudorHulban/authentication/helpers"
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

	pkCreatedUser, errCr := service.CreateUser(ctx, &p1)
	require.NoError(t, errCr)
	require.NotEqual(t,
		helpers.PrimaryKeyZero,
		pkCreatedUser,
	)

	_, errCrAgain := service.CreateUser(ctx, &p1)
	require.Error(t, errCrAgain)

	reconstructedUser, errGetByCredentials := service.GetUserByCredentials(
		ctx,
		&ParamsGetUser{
			Email:    testuser.TestUser.Email,
			Password: testuser.TestUser.Password,
		},
	)
	require.NoError(t, errGetByCredentials)
	require.NotZero(t, reconstructedUser)

	reconstructedUserInfo, errGetByID := service.GetUserInfoByID(
		ctx,
		reconstructedUser.PrimaryKey,
	)
	require.NoError(t, errGetByID)
	require.NotZero(t, reconstructedUserInfo)
	require.Equal(t,
		reconstructedUser.Name,
		reconstructedUserInfo.Name,
	)

	require.NoError(t,
		os.Remove(nameFile),
	)
}
