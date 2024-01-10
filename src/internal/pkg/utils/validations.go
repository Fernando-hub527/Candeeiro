package utils

import (
	"time"

	"github.com/Fernando-hub527/candieiro/internal/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func IsValidObjectId(number string) (primitive.ObjectID, *errors.RequestError) {
	result, errNumber := primitive.ObjectIDFromHex(number)
	if errNumber != nil {
		return result, errors.NewErrorInvalidParamns("Não pode validar id de ")
	}
	return result, nil
}

func validParams(values [][2]string) *errors.RequestError {
	validators := map[string]func(string) *errors.RequestError{
		"time": validTime,
	}

	for index := range values {
		funcValidator := validators[values[index][0]]
		if funcValidator != nil {
			if err := funcValidator(values[index][1]); err != nil {
				return err
			}
		}
	}
	return nil
}

func validParam(expectedType, value string) {

}

func validTime(date string) *errors.RequestError {
	_, err := time.Parse("", date)
	if err != nil {
		return errors.NewErrorInvalidParamns("Unable to deserialize date")
	}
	return nil
}
