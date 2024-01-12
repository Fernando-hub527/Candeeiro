package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	UserName string
	Email    string
	Whatsapp uint64
	PlanId   []primitive.ObjectID
}
