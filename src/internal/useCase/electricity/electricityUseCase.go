package electricity

import (
	"context"
	"time"

	"github.com/Fernando-hub527/candieiro/internal/entity"
	"github.com/Fernando-hub527/candieiro/internal/pkg/errors"
	consumutionrepository "github.com/Fernando-hub527/candieiro/internal/repository/consumutionRepository"
	pointrepository "github.com/Fernando-hub527/candieiro/internal/repository/pointRepository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IElectricityUseCase interface {
	FindPointById(pointId primitive.ObjectID, ctx context.Context) (*entity.Point, *errors.RequestError)
	ListPointsByPlant(plantId primitive.ObjectID, ctx context.Context) (*[]entity.Point, *errors.RequestError)
	ListConsumptionByIntervalAndPoint(pointId primitive.ObjectID, startMoment time.Time, endMoment time.Time, ctx context.Context) (*[]entity.Consumution, *errors.RequestError)
}

type ElectricityUseCase struct {
	repository            pointrepository.IPointRepository
	repositoryConsumution consumutionrepository.IConsumutionRepository
}

func NewElectricityUseCase(repository pointrepository.IPointRepository, repConsumution consumutionrepository.IConsumutionRepository) *ElectricityUseCase {
	return &ElectricityUseCase{
		repository:            repository,
		repositoryConsumution: repConsumution,
	}
}

func (elc *ElectricityUseCase) FindPointById(pointId primitive.ObjectID, ctx context.Context) (*entity.Point, *errors.RequestError) {
	return elc.repository.FindPointById(pointId, ctx)
}

func (elc *ElectricityUseCase) ListPointsByPlant(plantId primitive.ObjectID, ctx context.Context) (*[]entity.Point, *errors.RequestError) {
	return elc.repository.ListPointsByPlan(plantId, ctx)
}

func (elc *ElectricityUseCase) ListConsumptionByIntervalAndPoint(pointId primitive.ObjectID, startMoment time.Time, endMoment time.Time, ctx context.Context) (*[]entity.Consumution, *errors.RequestError) {
	dif := endMoment.Sub(startMoment)
	if dif > 360 || dif < 0 {
		return nil, errors.NewErrorInvalidParamns("Dates have invalid values. The end date must be greater than the start date and the interval must be less than 360 days")
	}

	return elc.repositoryConsumution.ListConsumutionByIntervalAndPoint(startMoment, endMoment, pointId, ctx)
}

func (elc *ElectricityUseCase) CreateConsumutionRecord() {

}
