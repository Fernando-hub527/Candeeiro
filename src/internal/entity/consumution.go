package entity

import (
	"time"
)

type Consumution struct {
	PointId     string
	Kw          uint32
	Cost        uint16
	StartMoment time.Time
	EndMoment   time.Time
}
