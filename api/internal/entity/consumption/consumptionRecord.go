package consumption

import (
	"time"

	"github.com/Fernando-hub527/candieiro/internal/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Consumution struct {
	Id                 primitive.ObjectID `bson:"_id"`
	PointId            primitive.ObjectID
	Kw                 uint32
	Cost               float64
	StartOfConsumption time.Time
	EndOfConsumption   time.Time
}

func FactoryConsumution(consumptionCost IConsumptionCost, pointId primitive.ObjectID, kw uint32, startTime, endTime time.Time) (*Consumution, *errors.RequestError) {
	if endTime.Hour() != startTime.Hour() {
		return nil, errors.NewErrorInvalidParamns("Consumption with a greater interval than allowed")
	}

	return &Consumution{
		PointId:            pointId,
		Kw:                 kw,
		StartOfConsumption: startTime,
		EndOfConsumption:   endTime,
		Cost:               consumptionCost.getConsumptionCost(kw),
	}, nil
}
