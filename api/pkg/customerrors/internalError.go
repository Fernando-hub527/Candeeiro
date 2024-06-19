package customerrors

import "time"

func NewInternalErros(detail string) *RequestError {
	return &RequestError{
		Detail: detail,
		Title:  "Internal Error",
		Status: 500,
		Time:   time.Now(),
	}
}
