package errors

import "time"

func NewErrorAccessDenied(detail string) *RequestError {
	return &RequestError{
		Detail: detail,
		Title:  "Access Denied",
		Status: 403,
		Time:   time.Now(),
	}
}
