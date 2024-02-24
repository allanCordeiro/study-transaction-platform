package database

import (
	"context"
	"errors"
	"time"

	"github.com/AllanCordeiro/study-transaction-platform/user-ms/internal/domain/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userEntity struct {
	//objectId  primitive.ObjectID `bson:"_id"`
	id        string    `bson:"id"`
	name      string    `bson:"name"`
	email     string    `bson:"email"`
	userType  uint8     `bson:"user_type"`
	password  string    `bson:"password"`
	createdAt time.Time `bson:"created_at"`
	updatedAt time.Time `bson:"updated_at"`
	deletedAt time.Time `bson:"deleted_at"`
	isActive  bool      `bson:"is_active"`
}

type UserDB struct {
	Coll *mongo.Collection
}

func NewUserDB(coll *mongo.Collection) *UserDB {
	return &UserDB{Coll: coll}
}

func (u *UserDB) Save(ctx context.Context, user *entity.User) error {
	newUser := userEntity{
		id:        user.Id,
		name:      user.Name,
		email:     user.Email.GetEmail(),
		userType:  user.UserType.EnumIndex(),
		password:  user.Password,
		createdAt: user.CreatedAt,
		updatedAt: user.UpdatedAt,
		deletedAt: user.DeletedAt,
	}

	_, err := u.Coll.InsertOne(ctx, newUser)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserDB) FindByMail(ctx context.Context, email string) (*entity.User, error) {
	filter := bson.D{{Key: "email", Value: email}}
	var retrievedUser userEntity
	err := u.Coll.FindOne(ctx, filter).Decode(retrievedUser)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, entity.ErrUserNotFound
		}
		return nil, err
	}
	userMail, _ := entity.NewEmail(retrievedUser.email)
	user := &entity.User{
		Id:        retrievedUser.id,
		Name:      retrievedUser.name,
		Email:     *userMail,
		UserType:  entity.UserType(retrievedUser.userType),
		Password:  retrievedUser.password,
		CreatedAt: retrievedUser.createdAt,
		UpdatedAt: retrievedUser.updatedAt,
		DeletedAt: retrievedUser.deletedAt,
		IsActive:  retrievedUser.isActive,
	}
	return user, nil
}

func (u *UserDB) FindByID(ctx context.Context, id string) (*entity.User, error) {
	filter := bson.D{{Key: "id", Value: id}}
	var retrievedUser userEntity
	err := u.Coll.FindOne(ctx, filter).Decode(retrievedUser)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, entity.ErrUserNotFound
		}
		return nil, err
	}
	userMail, _ := entity.NewEmail(retrievedUser.email)
	user := &entity.User{
		Id:        retrievedUser.id,
		Name:      retrievedUser.name,
		Email:     *userMail,
		UserType:  entity.UserType(retrievedUser.userType),
		Password:  retrievedUser.password,
		CreatedAt: retrievedUser.createdAt,
		UpdatedAt: retrievedUser.updatedAt,
		DeletedAt: retrievedUser.deletedAt,
		IsActive:  retrievedUser.isActive,
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

	result, err := u.Coll.UpdateOne(ctx, filter, update)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, entity.ErrUserNotFound
		}
		return nil, err
	}

	if result.ModifiedCount < 1 {
		return nil, entity.ErrUserNotChanged
	}
	return user, nil
}
