package dtos

import (
	"encoding/json"
	"io"

	"github.com/Fernando-hub527/candieiro/internal/pkg/errors"
	"github.com/go-playground/validator/v10"
)

type UserLoginDTO struct {
	Password string `validate:"required, min=10, max=100"`
	UserName string `validate:"required, min=10, max=50"`
}

type ResponseTokenDto struct {
	AccessToken  string
	RefreshToken string
	TokenId      string
}

type NewUserDTO struct {
	UserName  string `validate:"required, min=5, max=60"`
	Password  string `validate:"required, min=12, max=120"`
	Email     string `validate:"required, email"`
	Telephone uint64 `validate:"len=11"`
}

type ResponseUserDTO struct {
	UserName  string
	Email     string
	Telephone uint64
}

func FactoryNewDTO[T interface{}](newUser io.ReadCloser, dtoTarget T) (*T, *errors.RequestError) {

	json.NewDecoder(newUser).Decode(&dtoTarget)

	err := validator.New(validator.WithRequiredStructEnabled()).Struct(dtoTarget)

	if err != nil {
		return nil, errors.NewErrorInvalidParamns(err.(validator.ValidationErrors)[0].Error())
	}
	return &dtoTarget, nil
}
