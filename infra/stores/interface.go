package stores

import (
	"context"

	appuser "github.com/TudorHulban/authentication/domain/app-user"
	"github.com/TudorHulban/authentication/domain/ticket"
	storefile "github.com/TudorHulban/authentication/infra/stores/store-file"
	storememory "github.com/TudorHulban/authentication/infra/stores/store-memory"
)

type IStoreUser interface {
	CreateUser(ctx context.Context, user *appuser.User) error
	GetUserInfo(ctx context.Context, userCredentials *appuser.UserCredentials, result *appuser.UserInfo) error
	UpdateUserInfo(ctx context.Context, userCredentials *appuser.UserCredentials, userInfo *appuser.UserInfo) error
	DeleteUser(ctx context.Context, userCredentials *appuser.UserCredentials) error
}

var _ IStoreUser = &storememory.StoreUser{}

type IStoreTicket interface {
	CreateTicket(ctx context.Context, task *ticket.Ticket) error
	GetTicketByID(ctx context.Context, taskID ticket.PrimaryKeyTicket, result *ticket.TicketInfo) error
	SearchTasks(ctx context.Context, params *ticket.ParamsSearchTasks) (ticket.Tickets, error)
	CloseTask(ctx context.Context, taskID ticket.PrimaryKeyTicket, status ticket.TicketStatus) error

	AddEvent(ctx context.Context, taskID ticket.PrimaryKeyTicket, event *ticket.Event) error
	GetEventsForTaskID(ctx context.Context, taskID ticket.PrimaryKeyTicket) ([]*ticket.Event, error)
}

var _ IStoreTicket = &storememory.StoreTicket{}
var _ IStoreTicket = &storefile.StoreTickets{}
