package dtos

import (
	"encoding/json"
	"time"

	"github.com/Fernando-hub527/candieiro/internal/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type NewConsumutionDTO struct {
	Kw        int64
	StartTime time.Time
	EndTime   time.Time
	PointId   primitive.ObjectID
}

func FactoryDTO[T interface{}](jsonDTO string) (*T, *errors.RequestError) {
	var result T
	if err := json.Unmarshal([]byte(jsonDTO), &result); err != nil {
		return nil, errors.NewErrorInvalidParamns("Unable to deserialize message from point:\n" + jsonDTO)
	}
	return &result, nil
}
