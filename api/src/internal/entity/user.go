package entity

import (
	"runtime"

	"github.com/Fernando-hub527/candieiro/internal/dtos"
	"github.com/alexedwards/argon2id"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Password  string
	UserName  string
	Email     string
	Telephone uint64
	PlanId    []primitive.ObjectID
}

func FactoryNewUser(newUser dtos.NewUserDTO) (*User, error) {
	newPassword, err := argon2id.CreateHash(newUser.Password, &argon2id.Params{Memory: 28 * 1024, Iterations: 1, Parallelism: uint8(runtime.NumCPU()), SaltLength: 5, KeyLength: 20})
	if err != nil {
		return nil, err
	}

	return &User{
		Password:  newPassword,
		UserName:  newUser.UserName,
		Email:     newUser.Email,
		Telephone: newUser.Telephone,
		PlanId:    []primitive.ObjectID{},
	}, nil
}
