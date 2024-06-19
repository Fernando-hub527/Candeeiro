package customerrors

import "time"

func NewErrorNotFound(detail string) *RequestError {
	return &RequestError{
		Detail: detail,
		Title:  "Not found",
		Status: 404,
		Time:   time.Now(),
	}
}
