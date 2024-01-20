package user

import (
	"context"

	"github.com/Fernando-hub527/candieiro/internal/entity"
	"github.com/Fernando-hub527/candieiro/internal/pkg/errors"
	userrepository "github.com/Fernando-hub527/candieiro/internal/repository/userRepository"
	"github.com/alexedwards/argon2id"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IUserUseCase interface {
	ValidAccess(userName string, plantId primitive.ObjectID, ctx context.Context) *errors.RequestError
	validLogin(userName, password string) (*entity.User, *errors.RequestError)
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

func (useCase *UserUseCase) ValidLogin(userName, password string, ctx context.Context) (*entity.User, *errors.RequestError) {
	user, err := useCase.repository.FindUserByNameOrEmail(userName, "", ctx)

	if err != nil {
		return nil, errors.NewErrorAccessDenied("invalid username or password")
	}

	math, _ := argon2id.ComparePasswordAndHash(password, user.Password)
	if !math {
		return nil, errors.NewErrorAccessDenied("invalid username or password")
	}

	return user, nil
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
