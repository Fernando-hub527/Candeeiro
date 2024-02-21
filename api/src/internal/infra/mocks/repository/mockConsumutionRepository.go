package repository

import (
	"context"
	"time"

	"github.com/Fernando-hub527/candieiro/internal/entity/consumption"
	"github.com/Fernando-hub527/candieiro/internal/pkg/errors"
	"github.com/Fernando-hub527/candieiro/internal/pkg/utils"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type mockConsumutionRepository struct {
	mock.Mock
}

func NewMockConsumutionRepository() *mockConsumutionRepository {
	return &mockConsumutionRepository{}
}

func (mock *mockConsumutionRepository) ListConsumutionByIntervalAndPoint(startMoment, endMoment time.Time, pointId primitive.ObjectID, ctx context.Context) (*[]consumption.Consumution, *errors.RequestError) {
	args := mock.Called(startMoment, endMoment, pointId, ctx)
	return utils.ParseType[[]consumption.Consumution](args.Get(0)), utils.ParseType[errors.RequestError](args.Get(1))
}

func (mock *mockConsumutionRepository) UpdateConsumutionOrCreate(consumution consumption.Consumution) *errors.RequestError {
	args := mock.Called(consumution)
	return utils.ParseType[errors.RequestError](args.Get(0))
}

func (mock *mockConsumutionRepository) FindConsumutionRecordByIntervalAndPoint(startMoment, endMoment time.Time, pointId primitive.ObjectID, ctx context.Context) (*consumption.Consumution, *errors.RequestError) {
	args := mock.Called(startMoment, endMoment, pointId, ctx)
	return utils.ParseType[consumption.Consumution](args.Get(0)), utils.ParseType[errors.RequestError](args.Get(1))
}
