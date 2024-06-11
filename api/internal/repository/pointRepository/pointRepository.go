package pointrepository

import (
	"context"

	"github.com/Fernando-hub527/candieiro/internal/entity"
	"github.com/Fernando-hub527/candieiro/internal/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IPointRepository interface {
	CreatePoint(point entity.Point, ctx context.Context) (interface{}, *errors.RequestError)
	FindPointById(pointId primitive.ObjectID, ctx context.Context) (*entity.Point, *errors.RequestError)
	ListPointsByPlan(planId primitive.ObjectID, ctx context.Context) (*[]entity.Point, *errors.RequestError)
}

type PointRepository struct {
	collPoint *mongo.Collection
}

func New(database *mongo.Database) *PointRepository {
	return &PointRepository{
		collPoint: database.Collection("point"),
	}
}

func (repository *PointRepository) CreatePoint(point entity.Point, ctx context.Context) (interface{}, *errors.RequestError) {
	result, err := repository.collPoint.InsertOne(ctx, point)
	if err != nil {
		return nil, errors.NewInternalErros("Unable create point")
	}
	return result.InsertedID, nil
}

func (repository *PointRepository) FindPointById(pointId primitive.ObjectID, ctx context.Context) (*entity.Point, *errors.RequestError) {
	filter := bson.M{"_id": pointId}

	var result entity.Point
	err := repository.collPoint.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, errors.NewInternalErros("Unable find point")
	}
	return &result, nil
}

func (repository *PointRepository) ListPointsByPlan(planId primitive.ObjectID, ctx context.Context) (*[]entity.Point, *errors.RequestError) {
	filter := bson.M{"plan_id": planId}

	var points []entity.Point
	cursor, errFind := repository.collPoint.Find(ctx, filter)

	if errFind != nil {
		return nil, errors.NewInternalErros("Unable to list points per plant")
	}

	if err := cursor.All(ctx, &points); err != nil {
		return nil, errors.NewInternalErros("Unable to list points per plant")
	}

	return &points, nil

}

// func (repository *PointRepository) updateStatePoint(pointId primitive.ObjectID, state bool, ctx context.Context) errors.RequestError {
// 	update := bson.D{{Key: "$set", Value: bson.D{{Key: "state", Value: state}}}}

// 	_, err := repository.collPoint.UpdateByID(ctx, pointId, update)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (repository *PointRepository) updateAlertPoint(pointId string, alert uint16, ctx context.Context) ererrors.RequestErrorror {
// 	update := bson.D{{Key: "$set", Value: bson.D{{Key: "alert", Value: alert}}}}

// 	_, err := repository.collPoint.UpdateByID(ctx, pointId, update)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (repository *PointRepository) updateShutdownSchedule(pointId string, state bool) (*entity.Point, error) {

// }

// func (repository *PointRepository) listShutdownSchedule(pointId string, state bool) (*entity.Point, error) {

// }

// func (repository *PointRepository) removeShutdownSchedule(pointId string, state bool) (*entity.Point, error) {

// }
