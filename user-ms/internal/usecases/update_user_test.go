package usecases

import (
	"context"
	"testing"

	"github.com/AllanCordeiro/study-transaction-platform/user-ms/internal/domain/entity"
	"github.com/AllanCordeiro/study-transaction-platform/user-ms/internal/usecases/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateUser(t *testing.T) {
	t.Run("given an user when try to update it should be alright", func(t *testing.T) {
		expectedName := "John Doe"
		expectedEmail := "j.doe@test.com"
		expectedUserType := "customer"
		expectedPassword := "123456"
		user, _ := entity.NewUser(expectedName, expectedEmail, expectedUserType, expectedPassword)
		user.Activate()
		userToUpdate := UpdateUserUseCaseInput{
			UserId:   user.Id,
			Name:     expectedName,
			UserType: expectedUserType,
			IsActive: true,
		}
		gateway := mocks.NewDatabaseMock()
		gateway.On("FindByID", mock.Anything, mock.Anything).Return(user, nil)
		gateway.On("Update", mock.Anything, mock.Anything).Return(user, nil)

		uc := NewUpdateUserUseCase(gateway)
		err := uc.Execute(context.TODO(), userToUpdate)

		assert.Nil(t, err)
	})

	t.Run("given a nonexistent user when try to update it should thrown an error", func(t *testing.T) {
		userToUpdate := UpdateUserUseCaseInput{
			UserId: uuid.NewString(),
		}
		gateway := mocks.NewDatabaseMock()
		gateway.On("FindByID", mock.Anything, mock.Anything).Return(&entity.User{}, entity.ErrUserNotFound)
		gateway.On("Update", mock.Anything, mock.Anything).Return(&entity.User{}, nil)

		uc := NewUpdateUserUseCase(gateway)
		err := uc.Execute(context.TODO(), userToUpdate)

		assert.NotNil(t, err)
		assert.Equal(t, entity.ErrUserNotFound, err)
	})

}
