package auth

import "github.com/pkg/errors"

var (
	// ErrUnknownCredentials To be used when bad credentials.
	ErrUnknownCredentials = errors.New("bad credentials")

	ErrEmailExists = errors.New("email already exists")
)
