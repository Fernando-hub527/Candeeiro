package userrepository

import (
	"context"

	"github.com/Fernando-hub527/candieiro/internal/entity"
	"github.com/Fernando-hub527/candieiro/internal/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserRepository interface {
	CreateUser(user entity.User, ctx context.Context) *errors.RequestError
	FindUserByNameOrEmail(name string, email string, ctx context.Context) (*entity.User, *errors.RequestError)
}

type userRepository struct {
	collPoint *mongo.Collection
}

func New(database *mongo.Database) *userRepository {
	return &userRepository{
		collPoint: database.Collection("user"),
	}
}

func (repository *userRepository) CreateUser(user entity.User, ctx context.Context) *errors.RequestError {
	_, err := repository.collPoint.InsertOne(ctx, user)
	if err != nil {
		return errors.NewInternalErros("Unable to create user")
	}
	return nil
}

func (repository *userRepository) FindUserByNameOrEmail(name string, email string, ctx context.Context) (*entity.User, *errors.RequestError) {
	filter := bson.D{
		{Key: "$or",
			Value: bson.A{
				bson.M{"UserName": "Fernando"},
				bson.M{"Email": email},
			},
		},
	}

	var userResult entity.User
	err := repository.collPoint.FindOne(context.TODO(), filter).Decode(&userResult)
	if err != nil {
		return nil, errors.NewErrorNotFound("User not found")
	}

	return &userResult, nil
}
