package usecases

import (
	"context"
	"testing"

	"github.com/AllanCordeiro/study-transaction-platform/user-ms/internal/domain/entity"
	"github.com/AllanCordeiro/study-transaction-platform/user-ms/internal/usecases/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestChangePassword(t *testing.T) {
	t.Run("given an user when try to change password it should be alright", func(t *testing.T) {
		expectedName := "John Doe"
		expectedEmail := "j.doe@test.com"
		expectedUserType := "customer"
		expectedPassword := "123456"
		user, _ := entity.NewUser(expectedName, expectedEmail, expectedUserType, "123")
		user.Activate()
		userToUpdate := ChangePasswordUseCaseInput{
			UserId:      user.Id,
			OldPassword: "123",
			NewPassword: expectedPassword,
		}
		gateway := mocks.NewDatabaseMock()
		gateway.On("FindByID", mock.Anything, mock.Anything).Return(user, nil)
		gateway.On("Update", mock.Anything, mock.Anything).Return(user, nil)

		uc := NewChangePasswordUseCase(gateway)
		output, err := uc.Execute(context.TODO(), userToUpdate)

		assert.Nil(t, err)
		assert.True(t, output.Success)
	})

	t.Run("given an user when try to change his password but with an invalid old password it should thrown an error", func(t *testing.T) {
		expectedName := "John Doe"
		expectedEmail := "j.doe@test.com"
		expectedUserType := "customer"
		expectedPassword := "123456"
		user, _ := entity.NewUser(expectedName, expectedEmail, expectedUserType, "123")
		user.Activate()
		userToUpdate := ChangePasswordUseCaseInput{
			UserId:      user.Id,
			OldPassword: "123potato",
			NewPassword: expectedPassword,
		}
		gateway := mocks.NewDatabaseMock()
		gateway.On("FindByID", mock.Anything, mock.Anything).Return(user, nil)
		gateway.On("Update", mock.Anything, mock.Anything).Return(user, nil)

		uc := NewChangePasswordUseCase(gateway)
		output, err := uc.Execute(context.TODO(), userToUpdate)

		assert.NotNil(t, err)
		assert.False(t, output.Success)
		assert.Equal(t, entity.ErrOldPasswordInvalid, err)

	})

}
