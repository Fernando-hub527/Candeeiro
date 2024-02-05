package electricity_test

import (
	"context"
	"testing"
	"time"

	"github.com/Fernando-hub527/candieiro/internal/dtos"
	"github.com/Fernando-hub527/candieiro/internal/entity"
	"github.com/Fernando-hub527/candieiro/internal/infra/mocks/repository"
	"github.com/Fernando-hub527/candieiro/internal/pkg/errors"
	"github.com/Fernando-hub527/candieiro/internal/useCase/electricity"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestListConsumptionByIntervalAndPoint(t *testing.T) {
	pointRepository := repository.NewMockPointRepository()
	consumutionRepository := repository.NewMockConsumutionRepository()
	electricityUseCase := electricity.NewElectricityUseCase(pointRepository, consumutionRepository)

	t.Run("If point is not found, error is returned", func(t *testing.T) {
		pointId := primitive.NewObjectID()
		pointRepository.On("FindPointById", pointId, context.TODO()).Return(nil, errors.NewErrorNotFound("Point not found"))

		consumution, err := electricityUseCase.ListConsumptionByIntervalAndPoint(pointId, time.Now(), time.Now(), context.TODO())
		expectedError := errors.NewErrorNotFound("Point not found")
		expectedError.Time = err.Time

		assert.Nil(t, consumution)
		assert.Equal(t, expectedError, err)
	})

	t.Run("If final moment is less than initial moment, return error", func(t *testing.T) {
		pointId := primitive.NewObjectID()
		pointRepository.On("FindPointById", pointId, context.TODO()).Return(entity.Point{PlanId: primitive.NewObjectID(), Id: pointId, State: false, Rele: false, Alert: 0}, nil)

		consumution, err := electricityUseCase.ListConsumptionByIntervalAndPoint(pointId, time.Now().Add(5*time.Hour), time.Now(), context.TODO())
		expectedError := errors.NewErrorInvalidParamns("Dates have invalid values. The end date must be greater than the start date and the interval must be less than 360 days")
		expectedError.Time = err.Time

		assert.Nil(t, consumution)
		assert.Equal(t, expectedError, err)
	})

	t.Run("If interval is greater than 360 days return error", func(t *testing.T) {
		pointId := primitive.NewObjectID()
		pointRepository.On("FindPointById", pointId, context.TODO()).Return(entity.Point{PlanId: primitive.NewObjectID(), Id: pointId, State: false, Rele: false, Alert: 0}, nil)

		consumution, err := electricityUseCase.ListConsumptionByIntervalAndPoint(pointId, time.Now(), time.Now().Add(8664*time.Hour), context.TODO())
		expectedError := errors.NewErrorInvalidParamns("Dates have invalid values. The end date must be greater than the start date and the interval must be less than 360 days")
		expectedError.Time = err.Time

		assert.Nil(t, consumution)
		assert.Equal(t, expectedError, err)
	})

	t.Run("if parameter is valid, consumption is returned", func(t *testing.T) {
		pointId := primitive.NewObjectID()
		var expectedConsumution = []entity.Consumution{{Id: primitive.NewObjectID(), PointId: pointId, Kw: 150, Cost: 10, StartOfConsumption: time.Now(), EndOfConsumption: time.Now()}}

		pointRepository.On("FindPointById", pointId, context.TODO()).Return(entity.Point{PlanId: primitive.NewObjectID(), Id: pointId, State: false, Rele: false, Alert: 0}, nil)
		consumutionRepository.On("ListConsumutionByIntervalAndPoint", expectedConsumution[0].StartOfConsumption, expectedConsumution[0].EndOfConsumption, pointId, context.TODO()).Return(&expectedConsumution, nil)

		consumution, _ := electricityUseCase.ListConsumptionByIntervalAndPoint(pointId, expectedConsumution[0].StartOfConsumption, expectedConsumution[0].EndOfConsumption, context.TODO())

		assert.Equal(t, &expectedConsumution, consumution)
	})
}

func TestCreateConsumution(t *testing.T) {
	pointRepository := repository.NewMockPointRepository()
	consumutionRepository := repository.NewMockConsumutionRepository()
	electricityUseCase := electricity.NewElectricityUseCase(pointRepository, consumutionRepository)

	t.Run("If point is not found, error is returned", func(t *testing.T) {
		pointId := primitive.NewObjectID()
		pointRepository.On("FindPointById", pointId, context.TODO()).Return(nil, errors.NewErrorNotFound("Point not found"))
		err := electricityUseCase.CreateConsumutionRecord(dtos.NewConsumutionDTO{Kw: 55, StartTime: time.Now(), EndTime: time.Now(), PointId: pointId})
		expectedError := errors.NewErrorNotFound("Point not found")
		expectedError.Time = err.Time

		assert.Equal(t, expectedError, err)
	})

	t.Run("If interval invalid, error is returned", func(t *testing.T) {
		pointId := primitive.NewObjectID()
		pointRepository.On("FindPointById", pointId, context.TODO()).Return(entity.Point{PlanId: primitive.NewObjectID(), Id: pointId, State: false, Rele: false, Alert: 0}, nil)

		err := electricityUseCase.CreateConsumutionRecord(dtos.NewConsumutionDTO{Kw: 55, StartTime: time.Now().Add(5 * time.Hour), EndTime: time.Now(), PointId: pointId})
		expectedError := errors.NewErrorInvalidParamns("Consumption with a greater interval than allowed")
		expectedError.Time = err.Time

		assert.Equal(t, expectedError, err)
	})
}
