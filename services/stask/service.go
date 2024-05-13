package stask

import (
	"context"

	"github.com/TudorHulban/authentication/infra/stores"
)

type Service struct {
	store stores.IStoreTask
}

func NewService(store stores.IStoreTask) *Service {
	return &Service{
		store: store,
	}
}

func (s *Service) CreateTask(ctx context.Context, task *task.Task) error {}
