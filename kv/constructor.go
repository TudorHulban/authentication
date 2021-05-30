package kv

import (
	"auth"

	"github.com/TudorHulban/kv"
)

type Config struct {
	Store kv.KVStore
}

type AuthKV struct {
	Cfg Config
}

var _ auth.IAuthenticator = &AuthKV{}

func NewKVAuth(cfg Config) auth.IAuthenticator {
	return &AuthKV{
		Cfg: cfg,
	}
}

// interface methods to be implemented
func (k *AuthKV) Create(auth.Customer) error {
	return nil
}

func (k *AuthKV) UpdateEmail(custID int64, newEmail string) error           { return nil }
func (k *AuthKV) UpdateName(custID int64, firstName, lastName string) error { return nil }
func (k *AuthKV) UpdatePassword(custID int64, p string) error               { return nil }
func (k *AuthKV) Authenticate(email, password string) error                 { return nil }
func (k *AuthKV) LostPasswordRequest(email string) (string, error)          { return "", nil }
