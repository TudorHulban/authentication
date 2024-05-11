package fixtures

import appuser "github.com/TudorHulban/authentication/domain/app-user"

var testUser = appuser.User{
	UserCredentials: appuser.UserCredentials{
		Email:    "x@x.co",
		Password: "123",
	},
	UserInfo: appuser.UserInfo{
		Name: "John Doe",
	},
}
