package consumutionrepository

import (
	"context"
	"time"

	"github.com/Fernando-hub527/candieiro/internal/entity"
	"github.com/Fernando-hub527/candieiro/internal/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IConsumutionRepository interface {
	ListConsumutionByIntervalAndPoint(startMoment time.Time, endMoment time.Time, pointId primitive.ObjectID, ctx context.Context) (*[]entity.Consumution, *errors.RequestError)
	CreateConsumution(consumution entity.Consumution)
}

type ConsumutionRepository struct {
	collConsumution *mongo.Collection
}

func New(database *mongo.Database) *ConsumutionRepository {
	return &ConsumutionRepository{
		collConsumution: database.Collection("ratings"),
	}
}

func (repository *ConsumutionRepository) ListConsumutionByIntervalAndPoint(startMoment time.Time, endMoment time.Time, pointId primitive.ObjectID, ctx context.Context) (*[]entity.Consumution, *errors.RequestError) {
	filter := bson.D{
		{Key: "$and",
			Value: bson.A{
				bson.D{{Key: "start_moment", Value: bson.D{{Key: "$gt", Value: startMoment}}}},
				bson.D{{Key: "end_moment", Value: bson.D{{Key: "$lt", Value: endMoment}}}},
				bson.D{{Key: "point_id", Value: pointId}},
			},
		},
	}

	cursor, err := repository.collConsumution.Find(ctx, filter)
	if err != nil {
		return nil, errors.NewInternalErros("Unable to list consumption")
	}

	var consumution []entity.Consumution

	if err := cursor.All(ctx, &consumution); err != nil {
		return nil, errors.NewInternalErros("Unable to list consumption")
	}

	return &consumution, nil
}

func (repository *ConsumutionRepository) CreateConsumution(consumution entity.Consumution) {

}
