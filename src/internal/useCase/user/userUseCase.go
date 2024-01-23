package user

import (
	"context"

	"github.com/Fernando-hub527/candieiro/internal/dtos"
	"github.com/Fernando-hub527/candieiro/internal/entity"
	"github.com/Fernando-hub527/candieiro/internal/pkg/errors"
	userrepository "github.com/Fernando-hub527/candieiro/internal/repository/userRepository"
	"github.com/alexedwards/argon2id"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IUserUseCase interface {
	ValidAccess(userName string, plantId primitive.ObjectID, ctx context.Context) *errors.RequestError
	ValidLogin(userName, password string, ctx context.Context) (*entity.User, *errors.RequestError)
	CreateUser(user dtos.NewUserDTO, ctx context.Context) (*entity.User, *errors.RequestError)
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

func (useCase *UserUseCase) CreateUser(newUser dtos.NewUserDTO, ctx context.Context) (*entity.User, *errors.RequestError) {
	_, err := useCase.repository.FindUserByNameOrEmail(newUser.UserName, newUser.Email, ctx)
	if err == nil {
		return nil, errors.NewErrorAlreadyRegisteredUser("user" + newUser.UserName + "x is already registered")
	}

	user, errNewUser := entity.FactoryNewUser(newUser)
	if errNewUser != nil {
		return nil, errors.NewErrorInvalidParamns(errNewUser.Error())
	}

	if errUser := useCase.repository.CreateUser(*user, ctx); errUser != nil {
		return nil, errUser
	}
	return user, nil
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
