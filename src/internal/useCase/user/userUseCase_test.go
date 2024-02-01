package user_test

import (
	"context"
	"testing"

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
