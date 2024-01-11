package electricity

import (
	"fmt"
	"time"

	"github.com/Fernando-hub527/candieiro/internal/entity"
	"github.com/Fernando-hub527/candieiro/internal/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IElectricityUseCase interface {
	FindPointById(pointId primitive.ObjectID) (*entity.Point, *errors.RequestError)
	ListPointsByPlant(plantId primitive.ObjectID) (*[]entity.Point, *errors.RequestError)
	ListConsumptionByIntervalAndPoint(pointId primitive.ObjectID, startMoment time.Time, endMoment time.Time) (*[]entity.Point, *errors.RequestError)
}

type ElectricityUseCase struct {
}

func (*ElectricityUseCase) FindPointById(pointId primitive.ObjectID) (*entity.Point, *errors.RequestError) {
	fmt.Println("encontrando ponto por id", pointId)
	return nil, errors.NewInternalErros("not implemeted")
}

func (*ElectricityUseCase) ListPointsByPlant(plantId primitive.ObjectID) (*[]entity.Point, *errors.RequestError) {
	fmt.Println("listando pontos por planta", plantId)
	return nil, errors.NewInternalErros("not implemeted")
}

func (*ElectricityUseCase) ListConsumptionByIntervalAndPoint(pointId primitive.ObjectID, startMoment time.Time, endMoment time.Time) (*[]entity.Point, *errors.RequestError) {
	fmt.Println("listando consumo por intervalo e ponto", pointId, startMoment, endMoment)
	return nil, errors.NewInternalErros("not implemeted")
}
