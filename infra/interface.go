package infra

import (
	"context"

	appuser "github.com/TudorHulban/authentication/domain/app-user"
)

type IStore interface {
	CreateUser(ctx context.Context, user *appuser.User) error
	GetUserInfo(ctx context.Context, userCredentials *appuser.UserCredentials, userInfo *appuser.UserInfo) error
	UpdateUserInfo(ctx context.Context, userCredentials *appuser.UserCredentials, userInfo *appuser.UserInfo) error
	DeleteUser(ctx context.Context, userCredentials *appuser.UserCredentials) error
}
