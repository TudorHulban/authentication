package auth

// Customer concentrates customer information.
// http://woocommerce.github.io/woocommerce-rest-api-docs/#customer-properties
type Customer struct {
	ID            int64
	CreatedUNIX   int64
	LastLoginUNIX int64
	EMail         string
	FirstName     string
	LastName      string
	Password      string
	Role          string
}

type IAuthenticator interface {
	Create(Customer) error
	UpdateEmail(custID int64, newEmail string) error
	UpdateName(custID int64, firstName, lastName string) error
	UpdatePassword(custID int64, p string) error
	Authenticate(email, password string) error
	LostPasswordRequest(email string) (string, error)
}
