package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Consumution struct {
	Id          primitive.ObjectID `bson:"_id"`
	PointId     string
	Kw          uint32
	Cost        uint16
	StartMoment time.Time
	EndMoment   time.Time
}
