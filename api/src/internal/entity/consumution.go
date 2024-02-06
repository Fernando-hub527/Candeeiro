package entity

import (
	"time"

	"github.com/Fernando-hub527/candieiro/internal/dtos"
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

func FactoryConsumution(newConsumution dtos.NewConsumutionDTO) (*Consumution, *errors.RequestError) {
	if newConsumution.EndTime.Hour() != newConsumution.StartTime.Hour() {
		return nil, errors.NewErrorInvalidParamns("Consumption with a greater interval than allowed")
	}
	return &Consumution{
		PointId:            newConsumution.PointId,
		Kw:                 newConsumution.Kw,
		StartOfConsumption: newConsumution.StartTime,
		EndOfConsumption:   newConsumution.EndTime,
		Cost:               GetCostByKw(newConsumution.Kw),
	}, nil
}

func GetCostByKw(kw uint32) float64 {
	const costKW = 0.30
	return costKW * float64(kw)
}
