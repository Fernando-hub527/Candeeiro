package repository

import (
	"context"

	"github.com/Fernando-hub527/candieiro/internal/entity"
	"github.com/Fernando-hub527/candieiro/internal/pkg/errors"
	"github.com/Fernando-hub527/candieiro/internal/pkg/utils"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type mockPointRepository struct {
	mock.Mock
}

func NewMockPointRepository() *mockPointRepository {
	return &mockPointRepository{}
}

func (mock *mockPointRepository) CreatePoint(point entity.Point, ctx context.Context) (interface{}, *errors.RequestError) {
	args := mock.Called(point, ctx)
	return args.Get(0), utils.ParseType[errors.RequestError](args.Get(1))
}

func (mock *mockPointRepository) FindPointById(pointId primitive.ObjectID, ctx context.Context) (*entity.Point, *errors.RequestError) {
	args := mock.Called(pointId, ctx)
	return utils.ParseType[entity.Point](args.Get(0)), utils.ParseType[errors.RequestError](args.Get(1))
}

func (mock *mockPointRepository) ListPointsByPlan(planId primitive.ObjectID, ctx context.Context) (*[]entity.Point, *errors.RequestError) {
	args := mock.Called(planId, ctx)
	return utils.ParseType[[]entity.Point](args.Get(0)), utils.ParseType[errors.RequestError](args.Get(1))
}
