package ssessions

import (
	"time"

	cachelruspecialised "github.com/TudorHulban/LRU/cache-lruspecialised"
	appuser "github.com/TudorHulban/authentication/domain/app-user"
)

type Service struct {
	cacheAppUser *cachelruspecialised.CacheLRU[int64, appuser.User]
}

func NewService() *Service {
	return &Service{
		cacheAppUser: cachelruspecialised.NewCacheLRU[int64, appuser.User](
			&cachelruspecialised.ParamsNewCacheLRU{
				TTL:      5 * time.Minute,
				Capacity: 16,
			},
		),
	}
}

func (s *Service) getSessionID() int64 {
	return time.Now().UnixNano()
}

func (s *Service) PutUserTTL(user *appuser.User) int64 {
	sessionID := s.getSessionID()

	s.cacheAppUser.PutTTL(
		sessionID,
		*user,
	)

	return sessionID
}

func (s *Service) GetUser(sessionID int64) (*appuser.User, error) {
	return s.cacheAppUser.Get(sessionID)
}
