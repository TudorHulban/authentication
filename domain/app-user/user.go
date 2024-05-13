package appuser

type UserInfo struct {
	ID   uint
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
