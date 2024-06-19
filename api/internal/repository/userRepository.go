package repository

import (
	"github.com/Fernando-hub527/candieiro/internal/entity"
	"github.com/Fernando-hub527/candieiro/pkg/customerrors"
)

type UserRepository struct {
}

func NewUser() *UserRepository {
	return &UserRepository{}
}

func (u *UserRepository) CreateUser(user *entity.User) (*entity.User, *customerrors.RequestError) {
	return &entity.User{}, nil
}

func (u *UserRepository) FindUserByName(userName string) (*entity.User, *customerrors.RequestError) {
	return &entity.User{}, nil
}
