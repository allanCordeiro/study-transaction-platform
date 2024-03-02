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

	mt.Run("given an user not found when find should return error user not found", func(mt *mtest.T) {
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
