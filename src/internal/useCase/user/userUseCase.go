package user

import (
	"github.com/Fernando-hub527/candieiro/internal/entity"
	"github.com/Fernando-hub527/candieiro/internal/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IUserUseCase interface {
	ValidAccess(userName string, plantId primitive.ObjectID) *errors.RequestError
	CreateUser(entity.User) error
	liberarPlantaParaUsuario(userName string, plantId primitive.ObjectID)
}
