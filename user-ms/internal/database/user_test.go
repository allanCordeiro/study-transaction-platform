package database

import (
	"context"
	"testing"
	"time"

	"github.com/AllanCordeiro/study-transaction-platform/user-ms/internal/domain/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestCreateUser(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("given a valid user when call save should return ok", func(mt *mtest.T) {
		db := NewUserDB(mt.DB)
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		expectedEmail := "j.doe@test.com"
		expectedName := "John Doe"
		expectedPassword := "123456"
		expectedType := entity.Customer
		user, err := entity.NewUser(expectedName, expectedEmail, expectedType.String(), expectedPassword)
		assert.Nil(t, err)

		err = db.Save(context.TODO(), user)
		assert.Nil(t, err)
	})
}

func TestFindUserByID(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("given a valid id when call to find user should return its entity", func(mt *mtest.T) {
		expectedUserId := uuid.NewString()
		expectedEmail := "j.doe@test.com"
		expectedName := "John Doe"
		expectedPassword := "123456"
		expectedType := entity.Customer
		db := NewUserDB(mt.DB)

		firstUserBatch := mtest.CreateCursorResponse(1, "test.user", mtest.FirstBatch,
			bson.D{
				{Key: "id", Value: expectedUserId},
				{Key: "name", Value: expectedName},
				{Key: "email", Value: expectedEmail},
				{Key: "user_type", Value: expectedType.EnumIndex()},
				{Key: "password", Value: expectedPassword},
				{Key: "created_at", Value: time.Now()},
				{Key: "updated_at", Value: time.Now()},
				{Key: "is_active", Value: true},
			})

		mt.AddMockResponses(firstUserBatch)

		foundUser, err := db.FindByID(context.Background(), expectedUserId)
		assert.Nil(t, err)

		assert.Equal(t, expectedUserId, foundUser.Id)
		assert.Equal(t, expectedName, foundUser.Name)
		assert.Equal(t, expectedEmail, foundUser.Email.GetEmail())
		assert.Equal(t, expectedPassword, foundUser.Password)
		assert.Equal(t, expectedType.String(), foundUser.UserType.String())
		assert.NotNil(t, foundUser.CreatedAt)
		assert.NotNil(t, foundUser.UpdatedAt)
		assert.Zero(t, foundUser.DeletedAt)
		assert.True(t, foundUser.IsActive)
	})

	mt.Run("given a nonexistent user when find should return error user not found", func(mt *mtest.T) {
		db := NewUserDB(mt.DB)
		aRandomUserId := "1ccb2e7e-400c-4cdb-85d6-8b3f3311b34a"
		expectedErr := entity.ErrUserNotFound
		//no primitive data on purpose as we're testing an user not found scenario
		noneUserBatch := mtest.CreateCursorResponse(0, "test.user", mtest.FirstBatch)

		mt.AddMockResponses(noneUserBatch)

		_, err := db.FindByID(context.Background(), aRandomUserId)
		assert.NotNil(t, err)
		assert.Equal(t, expectedErr, err)
	})
}

func TestFindUserByEmail(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("given a valid id when call to find by email user should return its entity", func(mt *mtest.T) {
		expectedUserId := uuid.NewString()
		expectedEmail := "j.doe@test.com"
		expectedName := "John Doe"
		expectedPassword := "123456"
		expectedType := entity.Customer
		db := NewUserDB(mt.DB)

		firstUserBatch := mtest.CreateCursorResponse(1, "test.user", mtest.FirstBatch,
			bson.D{
				{Key: "id", Value: expectedUserId},
				{Key: "name", Value: expectedName},
				{Key: "email", Value: expectedEmail},
				{Key: "user_type", Value: expectedType.EnumIndex()},
				{Key: "password", Value: expectedPassword},
				{Key: "created_at", Value: time.Now()},
				{Key: "updated_at", Value: time.Now()},
				{Key: "is_active", Value: true},
			})

		mt.AddMockResponses(firstUserBatch)

		foundUser, err := db.FindByID(context.Background(), expectedEmail)
		assert.Nil(t, err)

		assert.Equal(t, expectedUserId, foundUser.Id)
		assert.Equal(t, expectedName, foundUser.Name)
		assert.Equal(t, expectedEmail, foundUser.Email.GetEmail())
		assert.Equal(t, expectedPassword, foundUser.Password)
		assert.Equal(t, expectedType.String(), foundUser.UserType.String())
		assert.NotNil(t, foundUser.CreatedAt)
		assert.NotNil(t, foundUser.UpdatedAt)
		assert.Zero(t, foundUser.DeletedAt)
		assert.True(t, foundUser.IsActive)
	})

	mt.Run("given a nonexistent user email when find should return error user not found", func(mt *mtest.T) {
		db := NewUserDB(mt.DB)
		aRandomUserEmail := "this_is_not_john_doe@email.com"
		expectedErr := entity.ErrUserNotFound
		//no primitive data on purpose as we're testing an user not found scenario
		noneUserBatch := mtest.CreateCursorResponse(0, "test.user", mtest.FirstBatch)

		mt.AddMockResponses(noneUserBatch)

		_, err := db.FindByID(context.Background(), aRandomUserEmail)
		assert.NotNil(t, err)
		assert.Equal(t, expectedErr, err)
	})
}

func TestUpdateUser(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("given a update user when call to update user should return ok", func(mt *mtest.T) {
		expectedEmail := "j.doe@test.com"
		expectedName := "John Doe"
		expectedPassword := "123456"
		expectedType := entity.Customer
		db := NewUserDB(mt.DB)
		user, _ := entity.NewUser("Joe Doe", expectedEmail, expectedType.String(), "123456")
		user.Activate()

		time.Sleep(time.Second * 1)
		expectedUser := bson.D{
			{Key: "ok", Value: 1},
			{Key: "value", Value: bson.D{
				{Key: "id", Value: user.Id},
				{Key: "name", Value: expectedName}, //see if the name really changes
				{Key: "email", Value: user.Email.GetEmail()},
				{Key: "user_type", Value: user.UserType.EnumIndex()},
				{Key: "password", Value: expectedPassword}, //see if the password really changes
				{Key: "created_at", Value: user.CreatedAt},
				{Key: "updated_at", Value: time.Now()}, //see if the updated at changes
				{Key: "is_active", Value: user.IsActive},
			}},
		}

		mt.AddMockResponses(expectedUser)

		updatedUser, err := db.Update(context.TODO(), user)
		assert.Nil(t, err)

		assert.Equal(t, user.Id, updatedUser.Id)
		assert.Equal(t, expectedName, updatedUser.Name)
		assert.Equal(t, expectedEmail, updatedUser.Email.GetEmail())
		assert.True(t, user.IsPasswordValid(expectedPassword))
		assert.Equal(t, expectedType.String(), updatedUser.UserType.String())
		assert.NotNil(t, updatedUser.CreatedAt)
		assert.Zero(t, updatedUser.DeletedAt)
		assert.True(t, updatedUser.IsActive)
		assert.True(t, updatedUser.CreatedAt.Before(updatedUser.UpdatedAt))
	})

}
