package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateUser(t *testing.T) {
	scenarios := []struct {
		name            string
		userName        string
		email           string
		password        string
		userType        string
		isValidScenario bool
		err             error
	}{
		{
			name:            "given a valid user when call new user should return user entity",
			userName:        "Allan Cordeiro",
			email:           "allan@email.com",
			password:        "123456",
			userType:        "customer",
			isValidScenario: true,
			err:             nil,
		},
		{
			name:            "given an invalid user email when call new user should return invalid email message",
			userName:        "Allan Cordeiro",
			email:           "allan.email.com",
			password:        "123456",
			userType:        "customer",
			isValidScenario: false,
			err:             ErrInvalidEmail,
		},
		{
			name:            "given an invalid user password when call new user should return invalid password",
			userName:        "Allan Cordeiro",
			email:           "allan@email.com",
			password:        "123456123456123456123456123456123456123456123456123456123456123456123456123456123456123456123456123456123456",
			userType:        "customer",
			isValidScenario: false,
			err:             ErrPasswordTooLong,
		},
		{
			name:            "given an invalid user profile when call new user should return invalid profile",
			userName:        "Allan Cordeiro",
			email:           "allan@email.com",
			password:        "123456",
			userType:        "potatoes",
			isValidScenario: false,
			err:             ErrInvalidUserProfile,
		},
	}

	for _, test := range scenarios {
		t.Run(test.name, func(t *testing.T) {
			user, err := NewUser(test.name, test.email, test.userType, test.password)
			if test.isValidScenario == true {
				assert.Nil(t, err)
				assert.Equal(t, test.name, user.Name)
				assert.Equal(t, test.email, user.Email.GetEmail())
				assert.Equal(t, test.userType, user.UserType.String())
				assert.True(t, user.IsPasswordValid(test.password))
			}
			if test.isValidScenario == false {
				assert.Equal(t, test.err, err)
			}
		})
	}
}

func TestActivateUser(t *testing.T) {
	userName := "Allan Cordeiro"
	email := "allan@email.com"
	password := "123456"
	userType := "customer"

	user, err := NewUser(userName, email, userType, password)
	assert.Nil(t, err)
	user.Activate()
	assert.NotNil(t, user.CreatedAt)
	assert.NotNil(t, user.UpdatedAt)
	assert.Zero(t, user.DeletedAt)
	assert.True(t, user.IsActive)
}

func TestDeactivateUser(t *testing.T) {
	userName := "Allan Cordeiro"
	email := "allan@email.com"
	password := "123456"
	userType := "customer"

	user, err := NewUser(userName, email, userType, password)
	assert.Nil(t, err)
	user.Activate()
	assert.True(t, user.IsActive)
	assert.Zero(t, user.DeletedAt)
	user.Deactivate()
	assert.False(t, user.IsActive)
	assert.NotZero(t, user.DeletedAt)
}

func TestChangePassword(t *testing.T) {
	userName := "Allan Cordeiro"
	email := "allan@email.com"
	password := "123456"
	userType := "customer"
	expectedPassword := "nova_senha"

	user, err := NewUser(userName, email, userType, password)
	assert.Nil(t, err)
	user.Activate()
	user.NewPassword(expectedPassword)

	assert.True(t, user.IsPasswordValid(expectedPassword))

}
