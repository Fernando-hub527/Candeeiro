package services

import "github.com/Fernando-hub527/candieiro/internal/repository"

type UserService struct {
	repository *repository.UserRepository
}

func NewUserService(userRepository ...*repository.UserRepository) *UserService {
	repository := repository.NewUser()

	if len(userRepository) > 0 {
		repository = userRepository[0]
	}

	return &UserService{
		repository: repository,
	}
}
