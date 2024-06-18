package repository

import "github.com/Fernando-hub527/candieiro/internal/entity"

type IUserRepository interface {
	createUser(user *entity.User) *entity.User
	findUserByName(userName string) *entity.User
}

type IElectricityRepository interface {
	recordConsumption(consumption *entity.Consumution) *entity.Consumution
	createConsumptionPlace(place *entity.ConsumptionPlace) *entity.ConsumptionPlace
	createPlan(plan *entity.Plan) *entity.Plan
}
