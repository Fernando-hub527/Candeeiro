package plan

import (
	"context"

	"github.com/Fernando-hub527/candieiro/internal/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PlanRepository struct {
	collPoint *mongo.Collection
}

func New(database *mongo.Database) *PlanRepository {
	return &PlanRepository{
		collPoint: database.Collection("plan"),
	}
}

func (repository *PlanRepository) findPlanById(planId string, ctx context.Context) (*entity.Plan, error) {
	var plan entity.Plan

	filter := bson.M{"plan_id": planId}
	err := repository.collPoint.FindOne(ctx, filter).Decode(&plan)
	if err != nil {
		return nil, err
	}

	return &plan, nil
}

func (repository *PlanRepository) createPlan(plan entity.Plan, ctx context.Context) (interface{}, error) {
	result, err := repository.collPoint.InsertOne(ctx, plan)
	if err != nil {
		return nil, err
	}

	return result.InsertedID, nil
}

func (repository *PlanRepository) removePlan(planId string, ctx context.Context) error {
	filter := bson.M{"_id": planId}
	_, err := repository.collPoint.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}
