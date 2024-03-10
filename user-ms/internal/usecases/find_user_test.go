package usecases

import (
	"context"
	"testing"

	"github.com/AllanCordeiro/study-transaction-platform/user-ms/internal/domain/entity"
	"github.com/AllanCordeiro/study-transaction-platform/user-ms/internal/usecases/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFindUser(t *testing.T) {
	user, _ := entity.NewUser("John Doe", "j.doe@test.com", "customer", "123456")
	user.Activate()
	gateway := mocks.NewDatabaseMock()
	gateway.On("FindByID", mock.Anything, mock.Anything).Return(user, nil)
	gateway.On("FindByMail", mock.Anything, mock.Anything).Return(user, nil)

	t.Run("given a valid email when find user should return ok", func(t *testing.T) {
		uc := NewFindUserUseCase(gateway)
		output, err := uc.Execute(context.TODO(), FindUserInput{Id: user.Id})

		assert.Nil(t, err)
		assert.Equal(t, "John Doe", output.Name)
		assert.Equal(t, "j.doe@test.com", output.Email)
		assert.Equal(t, "customer", output.UserType)
		assert.NotNil(t, user.Password)
		assert.True(t, user.IsActive)
		assert.NotNil(t, user.Id)
	})

	t.Run("given a valid id when find user should return ok", func(t *testing.T) {
		uc := NewFindUserUseCase(gateway)
		output, err := uc.Execute(context.TODO(), FindUserInput{Email: user.Email.GetEmail()})

		assert.Nil(t, err)
		assert.Equal(t, "John Doe", output.Name)
		assert.Equal(t, "j.doe@test.com", output.Email)
		assert.Equal(t, "customer", output.UserType)
		assert.NotNil(t, user.Password)
		assert.True(t, user.IsActive)
		assert.NotNil(t, user.Id)
	})

}
