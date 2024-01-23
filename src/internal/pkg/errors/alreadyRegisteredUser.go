package errors

import "time"

func NewErrorAlreadyRegisteredUser(detail string) *RequestError {
	return &RequestError{
		Detail: detail,
		Title:  "Already registered user",
		Status: 400,
		Time:   time.Now(),
	}
}
