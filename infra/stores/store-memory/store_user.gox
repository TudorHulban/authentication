package storememory

import (
	"context"
	"fmt"
	"sync"

	appuser "github.com/TudorHulban/authentication/domain/app-user"
)

type StoreUser struct {
	cache map[appuser.UserCredentials]appuser.UserInfo

	mu sync.RWMutex
}

func NewStoreMemory() *StoreUser {
	return &StoreUser{
		cache: make(
			map[appuser.UserCredentials]appuser.UserInfo,
		),
	}
}

func (s *StoreUser) CreateUser(ctx context.Context, user *appuser.User) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.cache[user.UserCredentials]; exists {
		return fmt.Errorf(
			"user with email %s already exists",
			user.Email,
		)
	}

	s.cache[user.UserCredentials] = user.UserInfo

	return nil
}

func (s *StoreUser) GetUserInfo(ctx context.Context, userCredentials *appuser.UserCredentials, userInfo *appuser.UserInfo) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	result, exists := s.cache[*userCredentials]
	if !exists {
		return fmt.Errorf(
			"user with email %s not found",
			userCredentials.Email,
		)
	}

	*userInfo = result

	return nil
}

func (s *StoreUser) UpdateUserInfo(ctx context.Context, userCredentials *appuser.UserCredentials, userInfo *appuser.UserInfo) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.cache[*userCredentials]; !exists {
		return fmt.Errorf(
			"user with email %s not found",
			userCredentials.Email,
		)
	}

	s.cache[*userCredentials] = *userInfo

	return nil
}

func (s *StoreUser) DeleteUser(ctx context.Context, userCredentials *appuser.UserCredentials) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.cache[*userCredentials]; !exists {
		return fmt.Errorf(
			"user with email %s not found",
			userCredentials.Email,
		)
	}

	delete(
		s.cache,
		*userCredentials,
	)

	return nil
}
