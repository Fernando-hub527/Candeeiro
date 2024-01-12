package pointrepository

import (
	"context"

	"github.com/Fernando-hub527/candieiro/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PointRepository struct {
	collPoint *mongo.Collection
}

func New(database *mongo.Database) *PointRepository {
	return &PointRepository{
		collPoint: database.Collection("point"),
	}
}

func (repository *PointRepository) createPoint(point entity.Point, ctx context.Context) (interface{}, error) {
	result, err := repository.collPoint.InsertOne(ctx, point)
	if err != nil {
		return nil, err
	}
	return result.InsertedID, nil
}

func (repository *PointRepository) findPointById(pointId primitive.ObjectID, ctx context.Context) (*entity.Point, error) {
	filter := bson.M{"_id": pointId}

	var result entity.Point
	err := repository.collPoint.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (repository *PointRepository) listPointsByPlan(planId primitive.ObjectID, ctx context.Context) (*[]entity.Point, error) {
	filter := bson.M{"plan_id": planId}

	var points []entity.Point
	cursor, errFind := repository.collPoint.Find(ctx, filter)

	if errFind != nil {
		return nil, errFind
	}

	if err := cursor.All(ctx, &points); err != nil {
		return nil, err
	}

	return &points, nil

}

func (repository *PointRepository) updateStatePoint(pointId primitive.ObjectID, state bool, ctx context.Context) error {
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "state", Value: state}}}}

	_, err := repository.collPoint.UpdateByID(ctx, pointId, update)
	if err != nil {
		return err
	}
	return nil
}

func (repository *PointRepository) updateAlertPoint(pointId string, alert uint16, ctx context.Context) error {
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "alert", Value: alert}}}}

	_, err := repository.collPoint.UpdateByID(ctx, pointId, update)
	if err != nil {
		return err
	}
	return nil
}

// func (repository *PointRepository) updateShutdownSchedule(pointId string, state bool) (*entity.Point, error) {

// }

// func (repository *PointRepository) listShutdownSchedule(pointId string, state bool) (*entity.Point, error) {

// }

// func (repository *PointRepository) removeShutdownSchedule(pointId string, state bool) (*entity.Point, error) {

// }
