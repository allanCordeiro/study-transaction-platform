package database

import (
	"context"
	"os"
	"testing"

	"github.com/AllanCordeiro/study-transaction-platform/user-ms/internal/domain/entity"
	"github.com/AllanCordeiro/study-transaction-platform/user-ms/test"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	db              *mongo.Database
	closeConnection func()
	userDb          *UserDB
)

func TestMain(m *testing.M) {
	db, closeConnection = test.OpenConnection()

	userDb = NewUserDB(db, "user")
	defer closeConnection()
	runCode := m.Run()

	os.Exit(runCode) //you need to use it all the time or those tests wont pass
}

func TestITCreateUser(t *testing.T) {
	t.Run("given an user when create user usecase should return ok", func(t *testing.T) {
		expectedName := "John Doe"
		expectedEmail := "j.doe@test.com"
		expectedType := "customer"
		expectedPassword := "1233456"
		expectedIsActive := true
		newUser, err := entity.NewUser(expectedName, expectedEmail, expectedType, expectedPassword)

		assert.Nil(t, err)
		newUser.Activate()
		err = userDb.Save(context.TODO(), newUser)
		assert.Nil(t, err)

		foundUser, err := userDb.FindByID(context.TODO(), newUser.Id)
		assert.Nil(t, err)
		assert.Equal(t, expectedName, foundUser.Name)
		assert.Equal(t, expectedEmail, foundUser.Email.GetEmail())
		assert.Equal(t, expectedType, foundUser.UserType.String())
		assert.True(t, foundUser.IsPasswordValid(expectedPassword))
		assert.Equal(t, expectedIsActive, foundUser.IsActive)

	})

	t.Run("given a distributor when create user usecase should return ok", func(t *testing.T) {
		expectedName := "John Doe"
		expectedEmail := "j.doe.distributor@test.com"
		expectedType := "distributor"
		expectedPassword := "1233456"
		expectedIsActive := true
		newUser, err := entity.NewUser(expectedName, expectedEmail, expectedType, expectedPassword)

		assert.Nil(t, err)
		newUser.Activate()
		err = userDb.Save(context.TODO(), newUser)
		assert.Nil(t, err)

		foundUser, err := userDb.FindByID(context.TODO(), newUser.Id)
		assert.Nil(t, err)
		assert.Equal(t, expectedName, foundUser.Name)
		assert.Equal(t, expectedEmail, foundUser.Email.GetEmail())
		assert.Equal(t, expectedType, foundUser.UserType.String())
		assert.True(t, foundUser.IsPasswordValid(expectedPassword))
		assert.Equal(t, expectedIsActive, foundUser.IsActive)

	})
}

func TestITFindUserByEmail(t *testing.T) {
	t.Run("given an email when find by email should return user data", func(t *testing.T) {
		expectedName := "John Doe"
		expectedEmail := "j.doe.distributor@test.com"
		expectedType := "distributor"
		expectedPassword := "1233456"
		expectedIsActive := true
		newUser, err := entity.NewUser(expectedName, expectedEmail, expectedType, expectedPassword)

		assert.Nil(t, err)
		newUser.Activate()
		err = userDb.Save(context.TODO(), newUser)
		assert.Nil(t, err)

		founduser, err := userDb.FindByMail(context.TODO(), expectedEmail)

		assert.Nil(t, err)
		assert.Equal(t, expectedName, founduser.Name)
		assert.Equal(t, expectedEmail, founduser.Email.GetEmail())
		assert.Equal(t, expectedType, founduser.UserType.String())
		assert.True(t, founduser.IsPasswordValid(expectedPassword))
		assert.Equal(t, expectedIsActive, founduser.IsActive)
		assert.NotNil(t, founduser.Id)

	})

	t.Run("given a nonexistent email when find by email should throw error", func(t *testing.T) {
		expectedEmail := "dummy_not-found@test.com"
		expectedError := entity.ErrUserNotFound

		founduser, err := userDb.FindByMail(context.TODO(), expectedEmail)

		assert.NotNil(t, err)
		assert.Nil(t, founduser)
		assert.Equal(t, expectedError, err)

	})
}

func TestITFindUserById(t *testing.T) {
	t.Run("given an id when find by id should return user data", func(t *testing.T) {
		expectedName := "John Doe"
		expectedEmail := "j.doe@test.com"
		expectedType := "customer"
		expectedPassword := "1233456"
		expectedIsActive := true
		newUser, err := entity.NewUser(expectedName, expectedEmail, expectedType, expectedPassword)

		assert.Nil(t, err)
		newUser.Activate()
		err = userDb.Save(context.TODO(), newUser)
		assert.Nil(t, err)

		founduser, err := userDb.FindByID(context.TODO(), newUser.Id)

		assert.Nil(t, err)
		assert.Equal(t, expectedName, founduser.Name)
		assert.Equal(t, expectedEmail, founduser.Email.GetEmail())
		assert.Equal(t, expectedType, founduser.UserType.String())
		assert.True(t, founduser.IsPasswordValid(expectedPassword))
		assert.Equal(t, expectedIsActive, founduser.IsActive)
		assert.NotNil(t, founduser.Id)

	})

	t.Run("given a nonexistent id when find by id should throw error", func(t *testing.T) {
		expectedId := "dummy_not-found@test.com"
		expectedError := entity.ErrUserNotFound

		founduser, err := userDb.FindByID(context.TODO(), expectedId)

		assert.NotNil(t, err)
		assert.Nil(t, founduser)
		assert.Equal(t, expectedError, err)

	})
}

func TestITUpdateUser(t *testing.T) {
	t.Run("given an user when create user usecase should return ok", func(t *testing.T) {
		expectedName := "John Doe"
		expectedEmail := "j.doe@test.com"
		expectedType := "customer"
		expectedPassword := "1233456"
		expectedIsActive := true
		newUser, err := entity.NewUser("Wrong Joao", expectedEmail, "distributor", "123")
		assert.Nil(t, err)
		newUser.Activate()
		err = userDb.Save(context.TODO(), newUser)
		assert.Nil(t, err)

		newUser.Name = expectedName
		newUser.SetProfile("customer")
		newUser.NewPassword(expectedPassword)
		_, err = userDb.Update(context.TODO(), newUser)
		assert.Nil(t, err)
		foundUser, err := userDb.FindByID(context.TODO(), newUser.Id)
		assert.Nil(t, err)

		assert.Equal(t, expectedName, foundUser.Name)
		assert.Equal(t, expectedEmail, foundUser.Email.GetEmail())
		assert.Equal(t, expectedType, foundUser.UserType.String())
		assert.True(t, foundUser.IsPasswordValid(expectedPassword))
		assert.Equal(t, expectedIsActive, foundUser.IsActive)
		assert.True(t, foundUser.CreatedAt.Before(foundUser.UpdatedAt))

	})
}
