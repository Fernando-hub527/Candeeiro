package consumption

import (
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ConsumptionFluctuation struct {
	PointId primitive.ObjectID
	KwH     uint32
	CostH   float64
	Time    time.Time
}

func FactoryConsumptionFluctuation(consumptionCost IConsumptionCost, pointId primitive.ObjectID, kwh uint32, time time.Time) *ConsumptionFluctuation {
	return &ConsumptionFluctuation{
		PointId: pointId,
		KwH:     kwh,
		Time:    time,
		CostH:   consumptionCost.getConsumptionCost(kwh),
	}
}

func (c *ConsumptionFluctuation) ToJson() ([]byte, error) {
	return json.Marshal(c)
}
