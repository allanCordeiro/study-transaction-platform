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
	//database, closeConnection := test.OpenConnection()
	db, closeConnection = test.OpenConnection()

	userDb = database.NewUserDB(db, "user")

	defer closeConnection()

	os.Exit(m.Run()) //you need to use it all the time or those tests wont pass
}

func TestCreateUser(t *testing.T) {
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
