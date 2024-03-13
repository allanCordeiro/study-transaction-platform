package usecases

import (
	"context"
	"os"
	"testing"

	"github.com/AllanCordeiro/study-transaction-platform/user-ms/internal/database"
	"github.com/AllanCordeiro/study-transaction-platform/user-ms/internal/domain/entity"
	"github.com/AllanCordeiro/study-transaction-platform/user-ms/test"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	db              *mongo.Database
	closeConnection func()
	userDb          *database.UserDB
)

func TestMain(m *testing.M) {
	db, closeConnection = test.OpenConnection()

	userDb = database.NewUserDB(db, "user")
	defer closeConnection()
	runCode := m.Run()

	os.Exit(runCode) //you need to use it all the time or those tests wont pass
}

func TestCreateUser(t *testing.T) {
	t.Run("given an user when call create user use case should be ok", func(t *testing.T) {
		expectedName := "John Doe"
		expectedEmail := "j.doe@test.com"
		expectedType := "customer"
		expectedPassword := "1233456"
		expectedInput := CreateUserInput{
			Name:     expectedName,
			Email:    expectedEmail,
			Password: expectedPassword,
			UserType: expectedType,
		}

		usecase := NewCreateUserUseCase(userDb)
		output, err := usecase.Execute(context.TODO(), expectedInput)

		assert.Nil(t, err)
		assert.NotNil(t, output)
	})

	t.Run("given an user when create a duplicated user should throw an error", func(t *testing.T) {
		expectedName := "John Doe"
		expectedEmail := "j.doe.duplicated@test.com"
		expectedType := "customer"
		expectedPassword := "1233456"
		expectedInput := CreateUserInput{
			Name:     expectedName,
			Email:    expectedEmail,
			Password: expectedPassword,
			UserType: expectedType,
		}
		expectedError := entity.ErrUserAlreadyExists
		usecase := NewCreateUserUseCase(userDb)
		_, err := usecase.Execute(context.TODO(), expectedInput)
		assert.Nil(t, err)

		//usecase = NewCreateUserUseCase(userDb)
		output, err := usecase.Execute(context.TODO(), expectedInput)

		assert.NotNil(t, err)
		assert.Nil(t, output)
		assert.Equal(t, expectedError, err)

	})
}
