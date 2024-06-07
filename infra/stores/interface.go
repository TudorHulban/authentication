package stores

import (
	"context"

	appuser "github.com/TudorHulban/authentication/domain/app-user"
	"github.com/TudorHulban/authentication/domain/ticket"
	"github.com/TudorHulban/authentication/helpers"
	storefile "github.com/TudorHulban/authentication/infra/stores/store-file"
)

type IStoreUser interface {
	CreateUser(ctx context.Context, user *appuser.User) error
	GetUserInfo(ctx context.Context, userCredentials *appuser.UserCredentials, result *appuser.UserInfo) error
	UpdateUserInfo(ctx context.Context, userCredentials *appuser.UserCredentials, userInfo *appuser.UserInfo) error
	DeleteUser(ctx context.Context, userCredentials *appuser.UserCredentials) error
}

var _ IStoreUser = &storefile.StoreUsers{}

type IStoreTicket interface {
	CreateTicket(ctx context.Context, task *ticket.Ticket, force ...bool) error
	GetTicketByID(ctx context.Context, taskID helpers.PrimaryKey, result *ticket.TicketInfo) error
	SearchTasks(ctx context.Context, params *ticket.ParamsSearchTasks) (ticket.Tickets, error)
	CloseTask(ctx context.Context, taskID helpers.PrimaryKey, status ticket.TicketStatus) error

	AddEvent(ctx context.Context, taskID helpers.PrimaryKey, event *ticket.Event) error
	GetEventsForTaskID(ctx context.Context, taskID helpers.PrimaryKey) ([]*ticket.Event, error)
}

var _ IStoreTicket = &storefile.StoreTickets{}
