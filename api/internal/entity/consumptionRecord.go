package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Consumution struct {
	Id                 primitive.ObjectID
	PointId            ConsumptionPlace
	Kw                 uint32
	Cost               float64
	StartOfConsumption time.Time
	EndOfConsumption   time.Time
}
