package suser

import (
	"context"

	"github.com/TudorHulban/authentication/apperrors"
	appuser "github.com/TudorHulban/authentication/domain/app-user"
	"github.com/TudorHulban/authentication/helpers"
	"github.com/TudorHulban/authentication/infra/stores"
	"github.com/TudorHulban/epochid"
	"github.com/asaskevich/govalidator"
)

type Service struct {
	store stores.IStoreUser
}

func NewService(store stores.IStoreUser) *Service {
	return &Service{
		store: store,
	}
}

type ParamsCreateUser struct {
	Email    string `valid:"required"`
	Password string `valid:"required"`

	Name string `valid:"required"`
}

func (s *Service) CreateUser(ctx context.Context, params *ParamsCreateUser) (helpers.PrimaryKey, error) {
	if _, errValidation := govalidator.ValidateStruct(params); errValidation != nil {
		return helpers.PrimaryKeyZero,
			apperrors.ErrValidation{
				Caller: "CreateUser",
				Issue:  errValidation,
			}
	}

	result := helpers.PrimaryKey(
		epochid.NewIDIncremental10KWCoCorrection(),
	)

	if errCr := s.store.CreateUser(
		ctx,
		&appuser.User{
			UserCredentials: appuser.UserCredentials{
				Email:    params.Email,
				Password: params.Password,
			},
			UserInfo: appuser.UserInfo{
				Name: params.Name,

				PrimaryKey: result,
			},
		},
	); errCr != nil {
		return helpers.PrimaryKeyZero,
			errCr
	}

	return result,
		nil
}

type ParamsGetUser struct {
	Email    string `valid:"required"`
	Password string `valid:"required"`
}

func (s *Service) GetUserByCredentials(ctx context.Context, params *ParamsGetUser) (*appuser.User, error) {
	if _, errValidation := govalidator.ValidateStruct(params); errValidation != nil {
		return nil,
			apperrors.ErrValidation{
				Caller: "GetUser",
				Issue:  errValidation,
			}
	}

	var result appuser.UserInfo

	if errGetUserInfo := s.store.GetUserInfoByCredentials(
		ctx,
		&appuser.UserCredentials{
			Email:    params.Email,
			Password: params.Password,
		},
		&result,
	); errGetUserInfo != nil {
		return nil, errGetUserInfo
	}

	return &appuser.User{
			UserCredentials: appuser.UserCredentials{
				Email: params.Email,
			},
			UserInfo: result,
		},
		nil
}

func (s *Service) GetUserInfoByID(ctx context.Context, pk helpers.PrimaryKey) (*appuser.UserInfo, error) {
	if pk == helpers.PrimaryKeyZero {
		return nil,
			apperrors.ErrNilInput{
				InputName: "pk",
			}
	}

	var result appuser.UserInfo

	if errGetUserInfo := s.store.GetUserInfoByID(
		ctx,
		pk,
		&result,
	); errGetUserInfo != nil {
		return nil, errGetUserInfo
	}

	return &result,
		nil
}

type ParamsUpdateUserInfo struct {
	Email    string `valid:"required"`
	Password string `valid:"required"`

	Name string `valid:"required"`
}

func (s *Service) UpdateUserInfo(ctx context.Context, params *ParamsUpdateUserInfo) error {
	if _, errValidation := govalidator.ValidateStruct(params); errValidation != nil {
		return apperrors.ErrValidation{
			Caller: "UpdateUserInfo",
			Issue:  errValidation,
		}
	}

	return s.store.UpdateUserInfo(
		ctx,
		&appuser.UserCredentials{
			Email:    params.Email,
			Password: params.Password,
		},
		&appuser.UserInfo{
			Name: params.Name,
		},
	)
}

type ParamsDeleteUser struct {
	Email    string `valid:"required"`
	Password string `valid:"required"`
}

func (s *Service) DeleteUser(ctx context.Context, params *ParamsDeleteUser) error {
	if _, errValidation := govalidator.ValidateStruct(params); errValidation != nil {
		return apperrors.ErrValidation{
			Caller: "DeleteUser",
			Issue:  errValidation,
		}
	}

	return s.store.DeleteUser(
		ctx,
		&appuser.UserCredentials{
			Email:    params.Email,
			Password: params.Password,
		},
	)
}
