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
	GetTaskByID(ctx context.Context, taskID task.PrimaryKeyTask, result *task.TaskInfo) error
	CloseTask(ctx context.Context, taskID task.PrimaryKeyTask, status task.TaskStatus) error

	AddEvent(ctx context.Context, taskID task.PrimaryKeyTask, event *task.Event) error
	GetEventsForTaskID(ctx context.Context, taskID task.PrimaryKeyTask) ([]*task.Event, error)
}

var _ IStoreTask = &storememory.StoreTask{}
