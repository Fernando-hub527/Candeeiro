package domain

import (
	"errors"
	"time"
)

type Record struct {
	startTime  time.Time
	endTime    time.Time
	totalKv    float64
	totalCost  float64
	consumerId uint32
}

func NewRecord(consumerId uint32, startTime, endTime time.Time, totalKv, totalCost float64) (*Record, error) {
	if startTime.Unix() > endTime.Unix() {
		return nil, errors.New("invalid period, start record must be before end record")
	}

	return &Record{
		consumerId: consumerId,
		startTime:  startTime,
		endTime:    endTime,
		totalKv:    totalKv,
		totalCost:  totalCost,
	}, nil
}

type Consumer struct {
	id           uint32
	placeId      uint32
	records      []*Record
	name         string
	description  string
	serialDevice uint64
}

func NewConsumer(id, placeId uint32, serialDevice uint64, name, description string, records ...*Record) *Consumer {
	return &Consumer{
		id:           id,
		placeId:      placeId,
		records:      records,
		name:         name,
		description:  description,
		serialDevice: serialDevice,
	}
}
