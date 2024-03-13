package usecases

import (
	"os"
	"testing"

	"github.com/AllanCordeiro/study-transaction-platform/user-ms/internal/database"
	"github.com/AllanCordeiro/study-transaction-platform/user-ms/test"
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
	// t.Run("given an user when create a duplicated user should throw an error", func(t *testing.T) {
	// 	expectedName := "John Doe"
	// 	expectedEmail := "j.doe@test.com"
	// 	expectedType := "customer"
	// 	expectedPassword := "1233456"
	// 	newUser, err := entity.NewUser(expectedName, expectedEmail, expectedType, expectedPassword)

	// 	assert.Nil(t, err)
	// 	newUser.Activate()
	// 	// err = userDb.Save(context.TODO(), newUser)
	// 	// assert.Nil(t, err)
	// 	result, err := userDb.DB.Collection("user").
	// 		InsertOne(context.TODO(), bson.M{"id": newUser.Id,
	// 			"name":       newUser.Name,
	// 			"email":      newUser.Email.GetEmail(),
	// 			"user_type":  newUser.UserType.EnumIndex(),
	// 			"password":   newUser.Password,
	// 			"created_at": newUser.CreatedAt,
	// 			"updated_at": newUser.UpdatedAt,
	// 			"deleted_at": newUser.DeletedAt,
	// 			"is_active":  newUser.IsActive,
	// 		})
	// 	assert.Nil(t, err)

	// 	log.Println(result.InsertedID)

	// 	err = userDb.Save(context.TODO(), newUser)
	// 	assert.NotNil(t, err)
	// 	assert.Equal(t, entity.ErrUserAlreadyExists, err)

	// })
}
