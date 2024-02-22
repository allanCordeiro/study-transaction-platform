package entity

import "errors"

var (
	ErrInvalidEmail       = errors.New("email invalido")
	ErrPasswordTooLong    = errors.New("password is too long")
	ErrInvalidUserProfile = errors.New("invalid user profile")
)
