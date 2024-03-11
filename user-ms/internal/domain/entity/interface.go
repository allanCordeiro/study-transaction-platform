package entity

import "context"

type UserInterface interface {
	Save(ctx context.Context, user *User) error
	FindByMail(ctx context.Context, email string) (*User, error)
	FindByID(ctx context.Context, id string) (*User, error)
	Update(ctx context.Context, user *User) (*User, error)
}
