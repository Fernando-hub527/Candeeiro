package dtos

import (
	"encoding/json"

	"github.com/Fernando-hub527/candieiro/internal/pkg/errors"
)

func FactoryDTO[T interface{}](jsonDTO string) (*T, *errors.RequestError) {
	var result T
	if err := json.Unmarshal([]byte(jsonDTO), &result); err != nil {
		return nil, errors.NewErrorInvalidParamns("Unable to deserialize message from point:\n" + jsonDTO)
	}
	return &result, nil
}
