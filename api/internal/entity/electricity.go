package entity

import "time"

type Plan struct {
	id     int64
	places *ConsumptionPlace
}

type ConsumptionPlace struct {
	deviceSerielNumber uint64
	createdAt          time.Time
}

func (p *ConsumptionPlace) GetDevice() uint64 {
	return p.deviceSerielNumber
}
