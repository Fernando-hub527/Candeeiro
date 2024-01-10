package electricity

import (
	"time"

	"github.com/Fernando-hub527/candieiro/internal/entity"
	"github.com/Fernando-hub527/candieiro/internal/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IElectricityUseCase interface {
	FindPointById(pointId primitive.ObjectID) (entity.Point, *errors.RequestError)
	ListPointsByPlant(plantId primitive.ObjectID) (*[]entity.Point, *errors.RequestError)
	ListConsumptionByIntervalAndPoint(pointId primitive.ObjectID, startMoment time.Time, endMoment time.Time) (*[]entity.Point, *errors.RequestError)
}
