package auth

// Customer concentrates customer information.
// http://woocommerce.github.io/woocommerce-rest-api-docs/#customer-properties
type Customer struct {
	EMail string

	CreatedUNIX   int64
	LastLoginUNIX int64

	FirstName    string
	LastName     string
	PasswordSalt string
	PasswordHash string
	Role         string
}

type IAuthenticator interface {
	Create(Customer) error
	CustomerDetails(email string) (*Customer, error)
	UpdateName(email, firstName, lastName string) error
	UpdatePassword(email, newPassword string) error
	Authenticate(email, password string) error
	LostPasswordRequest(email string) (string, error)
	Delete(email string) error
}
