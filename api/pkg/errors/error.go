package errors

import (
	"encoding/json"
	"time"
)

type RequestError struct {
	Title  string
	Status int16
	Detail string
	Time   time.Time
}

func (err *RequestError) ToString() string {
	result, errJson := json.Marshal(err)
	if errJson != nil {
		return ""
	}
	return string(result)
}
