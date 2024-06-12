package storefile

import (
	"context"

	appuser "github.com/TudorHulban/authentication/domain/app-user"
	genericfile "github.com/TudorHulban/authentication/infra/generic-file"
)

type StoreUsers struct {
	storeUsers *genericfile.GenericStoreFile[appuser.User]
}

type ParamsNewStoreUsers struct {
	PathCacheUsers string
}

func NewStoreUser(params *ParamsNewStoreUsers) *StoreUsers {
	return &StoreUsers{
		storeUsers: genericfile.
			NewGenericStoreFile[appuser.User](
			&genericfile.ParamsNewGenericStoreFile{
				PathStoreFile: params.PathCacheUsers,
			},
		),
	}
}

func (s *StoreUsers) CreateUser(ctx context.Context, item *appuser.User) error {
	return s.storeUsers.CreateItem(item, appuser.GetIDEmail)
}

func (s *StoreUsers) GetUserInfo(ctx context.Context, userCredentials *appuser.UserCredentials, result *appuser.UserInfo) error {
	reconstructedItem, errGet := s.storeUsers.SearchItem(appuser.CriteriaCredentials(userCredentials))
	if errGet != nil {
		return errGet
	}

	*result = reconstructedItem.UserInfo

	return nil
}

func (s *StoreUsers) UpdateUserInfo(ctx context.Context, userCredentials *appuser.UserCredentials, userInfo *appuser.UserInfo) error {
	var reconstructedItem appuser.UserInfo

	if errGet := s.GetUserInfo(ctx, userCredentials, &reconstructedItem); errGet != nil {
		return errGet
	}

	return s.storeUsers.UpdateItem(
		uint64(reconstructedItem.PrimaryKey),
		&appuser.User{
			UserCredentials: *userCredentials,
			UserInfo:        *userInfo,
		},
		appuser.GetIDUser,
	)
}

func (s *StoreUsers) DeleteUser(ctx context.Context, userCredentials *appuser.UserCredentials) error {
	var reconstructedItem appuser.UserInfo

	if errGet := s.GetUserInfo(ctx, userCredentials, &reconstructedItem); errGet != nil {
		return errGet
	}

	return s.storeUsers.UpdateItem(
		uint64(reconstructedItem.PrimaryKey),
		&appuser.User{
			UserCredentials: *userCredentials,
			UserInfo: appuser.UserInfo{
				PrimaryKey: reconstructedItem.PrimaryKey,
				Name:       reconstructedItem.Name,
				Timestamp:  reconstructedItem.Timestamp.WithDeleteNow(),
			},
		},
		appuser.GetIDUser,
	)
}
