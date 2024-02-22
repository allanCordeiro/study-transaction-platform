package entity

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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

func NewUser(name string, email string, userType string, password string) (*User, error) {

	emailEntity, err := NewEmail(email)
	if err != nil {
		return nil, err
	}

	if userType != Customer.String() && userType != Distributor.String() {
		return nil, ErrInvalidUserProfile
	}

	user := &User{
		Id:        uuid.New().String(),
		Name:      name,
		Email:     *emailEntity,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = user.Validate()
	if err != nil {
		return nil, err
	}
	user.SetProfile(userType)
	user.Password = hashPassword(user.Password)
	return user, nil
}

func (u *User) IsPasswordValid(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) Validate() error {
	//password should have less than 72 bytes
	if len([]byte(u.Password)) > 72 {
		return ErrPasswordTooLong
	}

	return nil
}

func (u *User) SetProfile(userType string) {
	if userType == Customer.String() {
		u.UserType = Customer
	}
	if userType == Distributor.String() {
		u.UserType = Distributor
	}
}

func (u *User) Activate() {
	u.IsActive = true
	u.UpdatedAt = time.Now()
	u.DeletedAt = time.Time{}
}

func (u *User) Deactivate() {
	u.IsActive = false
	u.UpdatedAt = time.Now()
	u.DeletedAt = time.Now()
}

func hashPassword(password string) string {
	//the only exception that this method should throw is whether password is greater than 72 bytes
	//we'll already validating it before
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}
