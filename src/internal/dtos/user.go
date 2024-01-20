package dtos

import (
	"github.com/Fernando-hub527/candieiro/internal/pkg/errors"
	"github.com/go-playground/validator/v10"
)

type userLoginDTO struct {
	Password string `validate:"required, gte=10, lte=100"`
	UserName string `validate:"required, gte=10, lte=50"`
}

func NewUserLogin(password, userName string) (*userLoginDTO, *errors.RequestError) {
	user := &userLoginDTO{Password: password, UserName: userName}
	err := validator.New(validator.WithRequiredStructEnabled()).Struct(user)

	if err != nil {
		return nil, errors.NewErrorInvalidParamns(err.(validator.ValidationErrors)[0].Error())
	}
	return user, nil
}

type ResponseTokenDto struct {
	AccessToken  string
	RefreshToken string
	TokenId      string
}
