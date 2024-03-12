package entity

import "errors"

var (
	ErrInvalidEmail         = errors.New("invalid email")
	ErrPasswordTooLong      = errors.New("password is too long")
	ErrOldPasswordInvalid   = errors.New("password doesnt match")
	ErrInvalidUserProfile   = errors.New("invalid user profile")
	ErrUserNotFound         = errors.New("user not found")
	ErrUserNotChanged       = errors.New("changes were not effective")
	ErrUserAlreadyActivated = errors.New("user is already activated")
	ErrUserAlreadyExists    = errors.New("user already exists")
)
