package consumutionrepository

import (
	"context"
	"time"

	"github.com/Fernando-hub527/candieiro/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ConsumutionRepository struct {
	collConsumution *mongo.Collection
}

func New(database *mongo.Database) *ConsumutionRepository {
	return &ConsumutionRepository{
		collConsumution: database.Collection("ratings"),
	}
}

func (repository *ConsumutionRepository) listConsumutionByIntervalAndPoint(startMoment time.Time, endMoment time.Time, pointId string, ctx context.Context) (*[]entity.Consumution, error) {
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
		return nil, err
	}

	var consumution []entity.Consumution

	if err := cursor.All(ctx, &consumution); err != nil {
		return nil, err
	}

	return &consumution, nil
}

func (repository *ConsumutionRepository) createConsumution(consumution entity.Consumution) {

}
