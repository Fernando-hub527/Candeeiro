package user_test

import (
	"context"
	"testing"

	"github.com/Fernando-hub527/candieiro/internal/dtos"
	"github.com/Fernando-hub527/candieiro/internal/entity"
	"github.com/Fernando-hub527/candieiro/internal/infra/mocks/repository"
	"github.com/Fernando-hub527/candieiro/internal/pkg/errors"
	"github.com/Fernando-hub527/candieiro/internal/useCase/user"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestValidAccess(t *testing.T) {
	mock := repository.NewMockUserRepository()
	useCase := user.NewUserCase(mock)
	t.Run("If user has access, nil is returned", func(t *testing.T) {
		planId := primitive.NewObjectID()
		mock.On("FindUserByNameOrEmail", "fernando", "", context.TODO()).Return(&entity.User{Password: "", UserName: "fernando", Email: "", Telephone: 77998574669, PlanId: []primitive.ObjectID{planId}}, nil)
		err := useCase.ValidAccess("fernando", planId, context.TODO())

		assert.Nil(t, err)
	})

	t.Run("If user not found, access denied is returned", func(t *testing.T) {
		mock.On("FindUserByNameOrEmail", "fernando", "", context.TODO()).Return(nil, errors.NewErrorNotFound("unregistered user"))
		err := useCase.ValidAccess("fernando", primitive.NewObjectID(), context.TODO())

		expectedError := errors.NewErrorAccessDenied("Unable to validate credentials")
		expectedError.Time = err.Time

		assert.Equal(t, expectedError, err)
	})

	t.Run("if user does not have access, access denied is returned", func(t *testing.T) {
		mock.On("FindUserByNameOrEmail", "fernando", "", context.TODO()).Return(&entity.User{Password: "", UserName: "fernando", Email: "", Telephone: 77998574669, PlanId: []primitive.ObjectID{}}, nil)
		err := useCase.ValidAccess("fernando", primitive.NewObjectID(), context.TODO())

		expectedError := errors.NewErrorAccessDenied("Unable to validate credentials")
		expectedError.Time = err.Time

		assert.Equal(t, expectedError, err)
	})
}

func TestValidLogin(t *testing.T) {
	mock := repository.NewMockUserRepository()
	useCase := user.NewUserCase(mock)

	t.Run("If login is valid, user is returned", func(t *testing.T) {
		user, _ := entity.FactoryNewUser(dtos.NewUserDTO{UserName: "fernando", Password: "123456", Email: "fernando.saraiva@gmail.com", Telephone: 77998574669})

		mock.On("FindUserByNameOrEmail", "fernando", "", context.TODO()).Return(user, nil)
		foundUser, err := useCase.ValidLogin(user.UserName, "123456", context.TODO())

		assert.Nil(t, err)
		assert.Equal(t, user, foundUser)
	})

	t.Run("If user not found, error is returned", func(t *testing.T) {
		mock.On("FindUserByNameOrEmail", "fernando", "", context.TODO()).Return(nil, errors.NewErrorNotFound("unregistered user"))

		user, err := useCase.ValidLogin("fernando", "1234", context.TODO())

		expectedError := errors.NewErrorAccessDenied("invalid username or password")
		expectedError.Time = err.Time

		assert.Equal(t, expectedError, err)
		assert.Nil(t, user)
	})

	t.Run("If password is invalid, error é retornado", func(t *testing.T) {
		user, _ := entity.FactoryNewUser(dtos.NewUserDTO{UserName: "fernando", Password: "123456", Email: "fernando.saraiva@gmail.com", Telephone: 77998574669})

		mock.On("FindUserByNameOrEmail", "fernando", "", context.TODO()).Return(user, nil)

		foundUser, err := useCase.ValidLogin(user.UserName, "1234", context.TODO())

		expectedError := errors.NewErrorAccessDenied("invalid username or password")
		expectedError.Time = err.Time

		assert.Nil(t, foundUser)
		assert.Equal(t, err, expectedError)
	})

	t.Run("if user does not have access, access denied is returned", func(t *testing.T) {
		mock.On("FindUserByNameOrEmail", "fernando", "", context.TODO()).Return(&entity.User{Password: "", UserName: "fernando", Email: "", Telephone: 77998574669, PlanId: []primitive.ObjectID{}}, nil)
		err := useCase.ValidAccess("fernando", primitive.NewObjectID(), context.TODO())

		expectedError := errors.NewErrorAccessDenied("Unable to validate credentials")
		expectedError.Time = err.Time

		assert.Equal(t, expectedError, err)
	})
}
