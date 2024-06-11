package ssessions

import (
	"time"

	"github.com/TudorHulban/authentication/apperrors"
	appuser "github.com/TudorHulban/authentication/domain/app-user"
	"github.com/TudorHulban/authentication/infra/cache/lru"
)

type Service struct {
	cacheAppUser *lru.CacheLRU[int64, appuser.User]
}

func NewService() *Service {
	return &Service{
		cacheAppUser: lru.NewCacheLRU[int64, appuser.User](
			&lru.ParamsNewCacheLRU{
				TTL:      30 * time.Minute,
				Capacity: 16,
			},
		),
	}
}

func (s *Service) getSessionID() int64 {
	return time.Now().UnixNano()
}

func (s *Service) PutUserTTL(user *appuser.User) (int64, error) {
	if user == nil {
		return 0,
			apperrors.ErrNilInput{
				InputName: "user",
			}
	}

	sessionID := s.getSessionID()

	s.cacheAppUser.PutTTL(
		sessionID,
		*user,
	)

	return sessionID,
		nil
}

func (s *Service) GetUser(sessionID int64) (*appuser.User, error) {
	return s.cacheAppUser.Get(sessionID)
}
