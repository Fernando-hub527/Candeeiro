package ports

import (
	"context"
	"time"

	"github.com/Fernando-hub527/candieiro/hexa/internal/core/consumer/domain"
)

type ConsumerService interface {
	ListPlantConsumptionByInterval(ctx context.Context, startTime, endTime time.Time, consumerId int) (*[]domain.Consumer, error)
	ListConsumptionByInterval(ctx context.Context, startTime, endTime time.Time, consumerId int) (*domain.Consumer, error)
	UpdateConsumerDevice(ctx context.Context, serialDevice uint64, consumerId uint32) (*domain.Consumer, error)
	ListenConsumption(ctx context.Context, consumerId int64) (*domain.Record, error)
}

type ConsumerRepository interface {
	CreateConsumptionRecord(ctx context.Context, record *domain.Record) (*domain.Record, error)
	ListConsumptionByIntervalAndConsumer(ctx context.Context, startTime, endTime time.Time, consumer uint32) (*[]domain.Record, error)
	ListConsumptionByIntervalAndPlant(ctx context.Context, startTime, endTime time.Time, plant uint32) (*[]domain.Record, error)

	CreateConsumer(ctx context.Context, consumer *domain.Consumer) (*domain.Consumer, error)
	ListConsumersByPlant(ctx context.Context, plant uint32) (*[]domain.Consumer, error)
	findConsumerById(ctx context.Context, id uint32) (*domain.Consumer, error)
}
