package entity

import "errors"

var (
	ErrInvalidEmail       = errors.New("email invalido")
	ErrPasswordTooLong    = errors.New("password is too long")
	ErrInvalidUserProfile = errors.New("invalid user profile")
	ErrUserNotFound       = errors.New("user not found")
	ErrUserNotChanged     = errors.New("changes were not effective")
)
