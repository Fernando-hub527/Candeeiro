package repository

import (
	"context"

	"github.com/Fernando-hub527/candieiro/internal/entity"
	"github.com/Fernando-hub527/candieiro/internal/pkg/errors"
	"github.com/stretchr/testify/mock"
)

type mockUserRepository struct {
	mock.Mock
}

func NewMockUserRepository() *mockUserRepository {
	return &mockUserRepository{}
}

func (mock *mockUserRepository) FindUserByNameOrEmail(name string, email string, ctx context.Context) (*entity.User, *errors.RequestError) {
	args := mock.Called(name, email, ctx)
	return parseType[entity.User](args.Get(0)), parseType[errors.RequestError](args.Get(1))
}

func (mock *mockUserRepository) CreateUser(user entity.User, ctx context.Context) *errors.RequestError {
	args := mock.Called(user, ctx)
	return args.Get(0).(*errors.RequestError)
}

func parseType[T interface{}](object interface{}) *T {
	obj, _ := object.(*T)
	return obj
}
