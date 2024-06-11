package dtos

import (
	"time"

	"github.com/Fernando-hub527/candieiro/internal/entity/consumption"
	"github.com/Fernando-hub527/candieiro/internal/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ConsumptionRecordRequestDTO struct {
	Kw        uint32
	StartTime time.Time
	EndTime   time.Time
	PointId   primitive.ObjectID
}

func (c *ConsumptionRecordRequestDTO) ParseToConsumptionRecord() (*consumption.Consumution, *errors.RequestError) {
	return consumption.FactoryConsumution(consumption.FactoryConsumptionCost(), c.PointId, c.Kw, c.EndTime, c.StartTime)
}

type ConsumptionFluctuationRequestDTO struct {
	Kwh       uint32             `json:"kw_h"`
	CreatedAt time.Time          `json:"time"`
	PointId   primitive.ObjectID `json:"SerialNumber"`
}

func (c *ConsumptionFluctuationRequestDTO) ParseToConsumptionFluctuation() *consumption.ConsumptionFluctuation {
	return consumption.FactoryConsumptionFluctuation(consumption.FactoryConsumptionCost(), c.PointId, c.Kwh, c.CreatedAt)
}

//   return "{\"kw_h\": "+ String(kwh) +", \"time\": "+ time +", \"serialNumber\": "+ serialNumber +"}";
