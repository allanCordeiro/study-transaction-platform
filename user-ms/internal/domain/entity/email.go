package entity

import (
	"net/mail"
)

type Email struct {
	value string
}

func NewEmail(email string) (Email, error) {
	mailAddress := Email{value: email}
	err := mailAddress.validate()
	if err != nil {
		return Email{}, err
	}
	return mailAddress, nil

}

func (e *Email) GetEmail() string {
	return e.value
}

func (e *Email) validate() error {
	_, err := mail.ParseAddress(e.value)
	if err != nil {
		return ErrInvalidEmail
	}
	return nil
}
