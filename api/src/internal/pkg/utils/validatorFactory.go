package utils

import (
	"time"

	"github.com/Fernando-hub527/candieiro/internal/pkg/errors"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// func MakeStandardValidator() map[string]func(string) *errors.RequestError {
// 	return map[string]func(string) *errors.RequestError{
// 		"time":     validTime,
// 		"objectId": isValidObjectId,
// 	}
// }

func ValidObjectId(number string, sendError func(ctx echo.Context, err errors.RequestError) error, ctx echo.Context) (*primitive.ObjectID, *errors.RequestError) {
	id, errNumber := primitive.ObjectIDFromHex(number)
	if errNumber != nil {
		err := errors.NewErrorInvalidParamns("Invalid id")
		sendError(ctx, *err)
		return nil, err
	}
	return &id, nil
}

func ValidTime(date string, sendError func(ctx echo.Context, err errors.RequestError) error, ctx echo.Context) (*time.Time, *errors.RequestError) {
	newDate, err := time.Parse(time.DateTime, date)
	if err != nil {
		err := errors.NewErrorInvalidParamns("cannot deserialize date, remember to use the format " + time.DateTime)
		sendError(ctx, *err)
		return nil, err
	}
	return &newDate, nil
}
