package appuser

import "github.com/TudorHulban/authentication/helpers"

type UserInfo struct {
	helpers.PrimaryKey

	Name string

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

func GetIDUser(item *User) uint64 {
	return uint64(item.PrimaryKey)
}

func GetIDEmail(item *User) uint64 {
	return helpers.NewWordFrom(item.Email).Hash()
}

var CriteriaCredentials = func(userCredentials *UserCredentials) func(item *User) bool {
	return func(item *User) bool {
		return item.UserCredentials == *userCredentials
	}
}
