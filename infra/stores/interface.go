package stores

import (
	"context"

	appuser "github.com/TudorHulban/authentication/domain/app-user"
	"github.com/TudorHulban/authentication/domain/task"
	storememory "github.com/TudorHulban/authentication/infra/stores/store-memory"
)

type IStoreUser interface {
	CreateUser(ctx context.Context, user *appuser.User) error
	GetUserInfo(ctx context.Context, userCredentials *appuser.UserCredentials, result *appuser.UserInfo) error
	UpdateUserInfo(ctx context.Context, userCredentials *appuser.UserCredentials, userInfo *appuser.UserInfo) error
	DeleteUser(ctx context.Context, userCredentials *appuser.UserCredentials) error
}

var _ IStoreUser = &storememory.StoreUser{}

type IStoreTask interface {
	CreateTask(ctx context.Context, task *task.Task) error
	GetTaskByID(ctx context.Context, taskID uint, result *task.Task) error
	UpdateTask(ctx context.Context, task *task.Task) error
	CloseTask(ctx context.Context, taskID, status uint) error

	CreateEvent(ctx context.Context, taskID uint, event *task.Event) error
	GetEventsForTaskID(ctx context.Context, taskID uint) ([]*task.Event, error)
}

var _ IStoreTask = &storememory.StoreTask{}
