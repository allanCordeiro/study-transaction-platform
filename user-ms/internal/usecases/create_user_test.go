package usecases

import (
	"context"
	"testing"

	"github.com/AllanCordeiro/study-transaction-platform/user-ms/internal/domain/entity"
	"github.com/AllanCordeiro/study-transaction-platform/user-ms/internal/usecases/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestValidCreateUserUseCase(t *testing.T) {
	t.Run("given a valid user when try to create should be ok", func(t *testing.T) {
		expectedName := "John Doe"
		expectedEmail := "j.doe@test.com"
		expectedPassword := "123456"
		expectedUserType := "customer"
		expectedUser := &CreateUserInput{
			Name:     expectedName,
			Email:    expectedEmail,
			Password: expectedPassword,
			UserType: expectedUserType,
		}
		gateway := mocks.NewDatabaseMock()
		gateway.On("Save", mock.Anything, mock.Anything).Return(nil)

		uc := NewCreateUserUseCase(gateway)
		output, err := uc.Execute(context.TODO(), *expectedUser)

		assert.Nil(t, err)
		assert.NotNil(t, output)
	})
}

func TestInvalidCreateUserUseCase(t *testing.T) {
	tests := []struct {
		name        string
		scenario    CreateUserInput
		expectedErr error
	}{
		{
			name: "given an invalid user email when try to create should receive error",
			scenario: CreateUserInput{
				Name:     "John Doe",
				Email:    "j.doe.test.com",
				Password: "123456",
				UserType: "customer",
			},
			expectedErr: entity.ErrInvalidEmail,
		},
		{
			name: "given an invalid user type when try to create should receive error",
			scenario: CreateUserInput{
				Name:     "John Doe",
				Email:    "j.doe@test.com",
				Password: "123456",
				UserType: "invalid_customer",
			},
			expectedErr: entity.ErrInvalidUserProfile,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gateway := mocks.NewDatabaseMock()
			gateway.On("Save", mock.Anything).Return(test.expectedErr)
			uc := NewCreateUserUseCase(gateway)
			_, err := uc.Execute(context.TODO(), test.scenario)

			assert.Equal(t, test.expectedErr, err)
		})
	}
}
