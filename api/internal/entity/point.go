package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Point struct {
	Id     primitive.ObjectID
	PlanId primitive.ObjectID
	State  bool
	Rele   bool
	Alert  uint16
}
