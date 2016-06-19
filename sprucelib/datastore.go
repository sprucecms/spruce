package sprucelib

import (
	"errors"
)

type DataStore interface {
	Close()
	// GetUserById(int) (User, error)
	// GetAllUsers() ([]User, error)
	GetUserByUsernameAndPassword(string, string) (User, error)
	GetUserByToken(string) (User, error)
	NewTokenForUser(User) (string, error)
	// StoreUser(*User) error
	// DeleteUser(*User) error
}

var ErrUnknownUser = errors.New("Unknown user")
var ErrInvalidToken = errors.New("Invalid token")
var ErrIncorrectPassword = errors.New("Incorrect password")
