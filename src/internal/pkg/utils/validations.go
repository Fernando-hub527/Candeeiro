package utils

import (
	"github.com/Fernando-hub527/candieiro/internal/pkg/errors"
)

func ValidParams(values [][2]string, validator map[string]func(string) *errors.RequestError) *errors.RequestError {

	for index := range values {
		funcValidator := validator[values[index][0]]
		if funcValidator != nil {
			if err := funcValidator(values[index][1]); err != nil {
				return err
			}
		} else {
			return errors.NewInternalErros("unknown parameter type")
		}
	}
	return nil
}
