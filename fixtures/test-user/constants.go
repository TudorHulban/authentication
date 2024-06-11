package testuser

import appuser "github.com/TudorHulban/authentication/domain/app-user"

var TestUser = appuser.User{
	UserCredentials: appuser.UserCredentials{
		Email:    "x@x.co",
		Password: "123",
	},
	UserInfo: appuser.UserInfo{
		Name: "John Doe",
	},
}
