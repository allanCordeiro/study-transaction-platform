package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        string
	Name      string
	Email     Email
	UserType  UserType
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	IsActive  bool
}

func NewUser(name string, email Email, userType UserType, password string) *User {
	return &User{
		Id:       uuid.New().String(),
		Name:     name,
		Email:    email,
		UserType: userType,
		Password: password,
	}
}
