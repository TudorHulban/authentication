package appuser

import (
	"github.com/TudorHulban/authentication/domain/ticket"
	"github.com/TudorHulban/authentication/helpers"
)

type UserInfo struct {
	helpers.PrimaryKey

	Name  string
	Level ticket.EventType

	helpers.Timestamp
}

type UserCredentials struct {
	Email    string
	Password string
}

type User struct {
	UserCredentials
	UserInfo
}

func GetIDUser(item *User) helpers.PrimaryKey {
	return item.PrimaryKey
}

func GetIDEmail(item *User) helpers.PrimaryKey {
	return helpers.PrimaryKey(
		helpers.NewWordFrom(item.Email).Hash(),
	)
}

var CriteriaCredentials = func(userCredentials *UserCredentials) func(item *User) bool {
	return func(item *User) bool {
		return item.UserCredentials == *userCredentials
	}
}

var CriteriaID = func(pk helpers.PrimaryKey) func(item *User) bool {
	return func(item *User) bool {
		return item.PrimaryKey == pk
	}
}
