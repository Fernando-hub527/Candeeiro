package timetools

import "time"

func SetSeconds(date time.Time, seconds int) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), date.Hour(), date.Minute(), seconds, date.Nanosecond(), date.Location())
}

func SetMinutes(date time.Time, minutes int) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), date.Hour(), minutes, date.Second(), date.Nanosecond(), date.Location())
}

func SetHour(date time.Time, hour int) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), hour, date.Minute(), date.Second(), date.Nanosecond(), date.Location())
}
