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
	mTestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mTestDb.Run("given a valid user when call save should return ok", func(mt *mtest.T) {
		userCollection := mt.Coll
		db := NewUserDB(userCollection)
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
	mTestDb := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mTestDb.Run("given a valid id when call to find user should return its entity", func(mt *mtest.T) {
		expectedUserId := uuid.NewString()
		expectedEmail := "j.doe@test.com"
		expectedName := "John Doe"
		expectedPassword := "123456"
		expectedType := entity.Customer
		mt.AddMockResponses(mtest.CreateCursorResponse(
			1,
			"foo.bar",
			mtest.FirstBatch,
			bson.D{
				{Key: "id", Value: expectedUserId},
				{Key: "name", Value: expectedName},
				{Key: "email", Value: expectedEmail},
				{Key: "password", Value: expectedPassword},
				{Key: "user_type", Value: expectedType.EnumIndex()},
				{Key: "created_at", Value: time.Now()},
				{Key: "updated_at", Value: time.Now()},
				{Key: "is_active", Value: true},
			},
		))
		userCollection := mt.Coll
		db := NewUserDB(userCollection)

		foundUser, err := db.FindByID(context.TODO(), expectedUserId)
		assert.Nil(t, err)

		assert.Equal(t, expectedUserId, foundUser.Id)
		assert.Equal(t, expectedName, foundUser.Name)
		assert.Equal(t, expectedEmail, foundUser.Email.GetEmail())
		assert.Equal(t, expectedPassword, foundUser.Password)
		assert.Equal(t, expectedType.EnumIndex(), foundUser.UserType)
		assert.NotNil(t, foundUser.CreatedAt)
		assert.NotNil(t, foundUser.UpdatedAt)
		assert.Nil(t, foundUser.DeletedAt)
		assert.True(t, foundUser.IsActive)
	})
}
