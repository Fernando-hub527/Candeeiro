package services

import (
	"github.com/Fernando-hub527/candieiro/internal/entity"
	"github.com/Fernando-hub527/candieiro/internal/repository"
	"github.com/Fernando-hub527/candieiro/pkg/auth"
	"github.com/Fernando-hub527/candieiro/pkg/customerrors"
)

type UserService struct {
	repository repository.IUserRepository
}

func NewUserService(userRepository ...repository.IUserRepository) IUserService {
	var rep repository.IUserRepository = repository.NewUser()
	if len(userRepository) > 0 {
		rep = userRepository[0]
	}

	return &UserService{
		repository: rep,
	}
}

func (u *UserService) CreateUser(userName, password string) (*entity.User, *customerrors.RequestError) {
	newUser, err := entity.NewUser(userName, password)
	if err != nil {
		return nil, err
	}

	userFound, _ := u.repository.FindUserByName(newUser.GetUserName())
	if userFound != nil {
		return nil, customerrors.NewErrorAlreadyRegisteredUser("The user with name" + userName + " is already registered")
	}

	return u.repository.CreateUser(newUser)
}

func (u *UserService) ValidUser(userName, password string) (bool, *customerrors.RequestError) {
	userFound, err := u.repository.FindUserByName(userName)
	if err != nil {
		return false, err
	}

	return userFound.ValidPassword(password)

}

func (u *UserService) GenerateToken(user *entity.User) (string, *customerrors.RequestError) {
	accessToken, errToken := auth.CreateToken(user.GetUserName())
	if errToken != nil {
		return "", customerrors.NewInternalErros("An error occurred while generating session, please try again later")
	}
	return accessToken, nil
}
