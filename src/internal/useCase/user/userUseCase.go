package user

import (
	"context"

	"github.com/Fernando-hub527/candieiro/internal/pkg/errors"
	userrepository "github.com/Fernando-hub527/candieiro/internal/repository/userRepository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IUserUseCase interface {
	ValidAccess(userName string, plantId primitive.ObjectID, ctx context.Context) *errors.RequestError
	// CreateUser(entity.User) error
	// liberarPlantaParaUsuario(userName string, plantId primitive.ObjectID)
}

type UserUseCase struct {
	repository userrepository.IUserRepository
}

func NewUserCase(repository userrepository.IUserRepository) *UserUseCase {
	return &UserUseCase{
		repository: repository,
	}
}

func (useCase *UserUseCase) ValidAccess(userName string, plantId primitive.ObjectID, ctx context.Context) *errors.RequestError {
	user, err := useCase.repository.FindUserByNameOrEmail(userName, "", ctx)
	if err != nil {
		return errors.NewErrorAccessDenied("Unregistered user")
	}

	for _, plan := range user.PlanId {
		if plan == plantId {
			return nil
		}
	}

	return errors.NewErrorAccessDenied("Unregistered user")
}
