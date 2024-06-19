package customerrors

import "time"

func NewErrorInvalidParamns(detail string) *RequestError {
	return &RequestError{
		Detail: detail,
		Title:  "Invalid paramn",
		Status: 400,
		Time:   time.Now(),
	}
}
