package entity

import (
	"runtime"

	"github.com/Fernando-hub527/candieiro/pkg/customerrors"
	"github.com/alexedwards/argon2id"
)

type User struct {
	userName string
	password string
	plans    []Plan
}

func validUser(userName, password string) *customerrors.RequestError {
	if len(userName) < 5 || len(userName) > 50 {
		return customerrors.NewErrorInvalidParamns("username with invalid length")
	}
	if len(password) < 5 || len(userName) > 100 {
		return customerrors.NewErrorInvalidParamns("username with invalid length")
	}
	return nil
}

func NewUser(name, password string, plans ...Plan) (*User, *customerrors.RequestError) {
	err := validUser(name, password)
	if err != nil {
		return nil, err
	}
	newPassword, passErr := argon2id.CreateHash(password, &argon2id.Params{Memory: 28 * 1024, Iterations: 1, Parallelism: uint8(runtime.NumCPU()), SaltLength: 5, KeyLength: 20})
	if passErr != nil {
		return nil, customerrors.NewInternalErros("An error occurred while creating user, please try again later")
	}

	return &User{
		userName: name,
		password: newPassword,
		plans:    plans,
	}, nil
}

func (u *User) ValidPassword(hash string) (match bool, err *customerrors.RequestError) {
	result, errPass := argon2id.ComparePasswordAndHash(u.password, hash)
	if errPass != nil {
		return false, customerrors.NewInternalErros("An error occurred while validating the password, please try again later")
	}
	return result, nil
}

func (u *User) GetUserName() string {
	return u.userName
}

func (u *User) GetPlans() []Plan {
	return u.plans
}
