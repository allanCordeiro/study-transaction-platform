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

func TestDeleteUser(t *testing.T) {
	t.Run("given an existent user when try to delete should be alright", func(t *testing.T) {
		user, _ := entity.NewUser("John Doe", "j.doe@test.com", "customer", "123456")
		user.Deactivate()
		gateway := mocks.NewDatabaseMock()
		gateway.On("FindByID", mock.Anything, mock.Anything).Return(user, nil)
		gateway.On("Update", mock.Anything, mock.Anything).Return(user, nil)

		uc := NewDeleteUserUseCase(gateway)
		err := uc.Execute(context.TODO(), DeleteUserInput{UserId: user.Id})

		assert.Nil(t, err)
	})

	t.Run("given a nonexistent user when try to delete should shown an error", func(t *testing.T) {
		aRandomId := uuid.NewString()
		gateway := mocks.NewDatabaseMock()
		gateway.On("FindByID", mock.Anything, mock.Anything).Return(&entity.User{}, entity.ErrUserNotFound)
		gateway.On("Update", mock.Anything, mock.Anything).Return(&entity.User{}, nil)

		uc := NewDeleteUserUseCase(gateway)
		err := uc.Execute(context.TODO(), DeleteUserInput{UserId: aRandomId})

		assert.Equal(t, entity.ErrUserNotFound, err)
	})
}
