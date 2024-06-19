package services

import (
	"time"

	"github.com/Fernando-hub527/candieiro/internal/entity"
	"github.com/Fernando-hub527/candieiro/pkg/customerrors"
)

type IUserService interface {
	CreateUser(userName, password string) (*entity.User, *customerrors.RequestError)
	ValidUser(userName, password string) (bool, *customerrors.RequestError)
	GenerateToken(user *entity.User) (string, *customerrors.RequestError)
}

type IElectricityService interface {
	CreatePlan(plan *entity.Plan)
	CreateConsumptionPlace(place *entity.ConsumptionPlace)
	listConsumptionByIntervalAndPlace(start, end time.Time, placeId uint64)
	listConsumptionByIntervalAndPlan(start, end time.Time, planId uint64)
	recordConsumption(record *entity.Consumution)
}
