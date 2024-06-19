package repository

import (
	"github.com/Fernando-hub527/candieiro/internal/entity"
	"github.com/Fernando-hub527/candieiro/pkg/customerrors"
)

type IUserRepository interface {
	CreateUser(user *entity.User) (*entity.User, *customerrors.RequestError)
	FindUserByName(userName string) (*entity.User, *customerrors.RequestError)
}

type IElectricityRepository interface {
	RecordConsumption(consumption *entity.Consumution) *entity.Consumution
	CreateConsumptionPlace(place *entity.ConsumptionPlace) *entity.ConsumptionPlace
	CreatePlan(plan *entity.Plan) *entity.Plan
}
