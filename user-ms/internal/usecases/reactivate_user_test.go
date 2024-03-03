package usecases

import (
	"context"
	"testing"

	"github.com/AllanCordeiro/study-transaction-platform/user-ms/internal/domain/entity"
	"github.com/AllanCordeiro/study-transaction-platform/user-ms/internal/usecases/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestReactivateUser(t *testing.T) {
	t.Run("given an inactive user when try to reactivate it should be alright", func(t *testing.T) {
		user, _ := entity.NewUser("John Doe", "j.doe@test.com", "customer", "123456")
		user.Activate()
		gateway := mocks.NewDatabaseMock()
		gateway.On("FindByID", mock.Anything, mock.Anything).Return(user, nil)
		gateway.On("Update", mock.Anything, mock.Anything).Return(user, nil)
		user.Deactivate() //in this moment to garantee if the rules are ok (sendind active in the mock and deactivate after this)

		uc := NewReactivateUserUseCase(gateway)
		err := uc.Execute(context.TODO(), ReactivateUserInput{UserId: user.Id})

		assert.Nil(t, err)
	})

	t.Run("given an active user when try to reactivate it should thrown an error", func(t *testing.T) {
		user, _ := entity.NewUser("John Doe", "j.doe@test.com", "customer", "123456")
		user.Activate()
		gateway := mocks.NewDatabaseMock()
		gateway.On("FindByID", mock.Anything, mock.Anything).Return(user, nil)
		gateway.On("Update", mock.Anything, mock.Anything).Return(user, nil)

		uc := NewReactivateUserUseCase(gateway)
		err := uc.Execute(context.TODO(), ReactivateUserInput{UserId: user.Id})

		assert.NotNil(t, entity.ErrUserAlreadyActivated, err)
	})

}
