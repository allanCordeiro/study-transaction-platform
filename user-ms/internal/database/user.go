package database

import (
	"context"
	"errors"
	"time"

	"github.com/AllanCordeiro/study-transaction-platform/user-ms/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userEntity struct {
	Id        string    `bson:"id"`
	Name      string    `bson:"name"`
	Email     string    `bson:"email"`
	UserType  uint8     `bson:"user_type"`
	Password  string    `bson:"password"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
	DeletedAt time.Time `bson:"deleted_at"`
	IsActive  bool      `bson:"is_active"`
}

type UserDB struct {
	DB *mongo.Database
}

func NewUserDB(db *mongo.Database) *UserDB {
	return &UserDB{DB: db}
}

func (u *UserDB) Save(ctx context.Context, user *entity.User) error {
	newUser := userEntity{
		Id:        user.Id,
		Name:      user.Name,
		Email:     user.Email.GetEmail(),
		UserType:  user.UserType.EnumIndex(),
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}

	_, err := u.DB.Collection("user").InsertOne(ctx, newUser)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserDB) FindByMail(ctx context.Context, email string) (*entity.User, error) {
	filter := bson.D{{Key: "email", Value: email}}
	var retrievedUser userEntity
	err := u.DB.Collection("user").FindOne(ctx, filter).Decode(&retrievedUser)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, entity.ErrUserNotFound
		}
		return nil, err
	}
	userMail, _ := entity.NewEmail(retrievedUser.Email)
	user := &entity.User{
		Id:        retrievedUser.Id,
		Name:      retrievedUser.Name,
		Email:     &userMail,
		UserType:  entity.UserType(retrievedUser.UserType),
		Password:  retrievedUser.Password,
		CreatedAt: retrievedUser.CreatedAt,
		UpdatedAt: retrievedUser.UpdatedAt,
		DeletedAt: retrievedUser.DeletedAt,
		IsActive:  retrievedUser.IsActive,
	}
	return user, nil
}

func (u *UserDB) FindByID(ctx context.Context, id string) (*entity.User, error) {
	filter := bson.M{"id": id}
	var retrievedUser userEntity
	err := u.DB.Collection("user").FindOne(ctx, filter).Decode(&retrievedUser)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, entity.ErrUserNotFound
		}
		return nil, err
	}
	userMail, _ := entity.NewEmail(retrievedUser.Email)
	user := &entity.User{
		Id:        retrievedUser.Id,
		Name:      retrievedUser.Name,
		Email:     &userMail,
		UserType:  entity.UserType(retrievedUser.UserType),
		Password:  retrievedUser.Password,
		CreatedAt: retrievedUser.CreatedAt,
		UpdatedAt: retrievedUser.UpdatedAt,
		DeletedAt: retrievedUser.DeletedAt,
		IsActive:  retrievedUser.IsActive,
	}
	return user, nil
}

func (u *UserDB) Update(ctx context.Context, user *entity.User) (*entity.User, error) {
	filter := bson.D{{Key: "id", Value: user.Id}}
	update := bson.D{{Key: "$set", Value: bson.D{
		{Key: "name", Value: user.Name},
		{Key: "password", Value: user.Password},
		{Key: "user_type", Value: user.UserType.EnumIndex()},
		{Key: "updated_at", Value: user.UpdatedAt},
		{Key: "deleted_at", Value: user.DeletedAt},
		{Key: "is_active", Value: user.IsActive},
	}}}

	var updatedUser userEntity

	err := u.DB.Collection("user").FindOneAndUpdate(
		ctx,
		filter,
		update,
		options.FindOneAndUpdate().SetReturnDocument(1),
	).Decode(&updatedUser)

	if err != nil {
		return nil, err
	}

	mail, _ := entity.NewEmail(updatedUser.Email)
	return &entity.User{
		Id:        updatedUser.Id,
		Name:      updatedUser.Name,
		Email:     &mail,
		UserType:  entity.UserType(updatedUser.UserType),
		CreatedAt: updatedUser.CreatedAt,
		UpdatedAt: updatedUser.UpdatedAt,
		DeletedAt: updatedUser.DeletedAt,
		IsActive:  updatedUser.IsActive,
	}, nil

}
