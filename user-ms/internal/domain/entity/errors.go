package entity

import "errors"

var (
	ErrInvalidEmail         = errors.New("invalid email")
	ErrPasswordTooLong      = errors.New("password is too long")
	ErrInvalidUserProfile   = errors.New("invalid user profile")
	ErrUserNotFound         = errors.New("user not found")
	ErrUserNotChanged       = errors.New("changes were not effective")
	ErrUserAlreadyActivated = errors.New("user is already activated")
)
