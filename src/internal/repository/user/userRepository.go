package user

import (
	"context"

	"github.com/Fernando-hub527/candieiro/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collPoint *mongo.Collection
}

func New(database *mongo.Database) *UserRepository {
	return &UserRepository{
		collPoint: database.Collection("users"),
	}
}

func (repository *UserRepository) createUser(user entity.User, ctx context.Context) error {
	_, err := repository.collPoint.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (repository *UserRepository) findUserByNameOrEmail(name string, email string, ctx context.Context) (*entity.User, error) {
	filter := bson.D{
		{Key: "$and",
			Value: bson.A{
				bson.M{"name": name},
				bson.M{"email": email},
			},
		},
	}

	var userResult entity.User
	err := repository.collPoint.FindOne(ctx, filter).Decode(&userResult)
	if err != nil {
		return nil, err
	}

	return &userResult, nil
}
