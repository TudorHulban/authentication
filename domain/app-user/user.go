package appuser

type UserInfo struct {
	Name string
}

type UserCredentials struct {
	Email    string
	Password string
}

type User struct {
	UserCredentials
	UserInfo
}
