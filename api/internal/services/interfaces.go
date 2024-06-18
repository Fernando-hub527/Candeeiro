package services

import "github.com/Fernando-hub527/candieiro/internal/entity"

type IUserService interface {
	CreateUser(userName, password string) *entity.User
	ValidUser(userName, password string) bool
}
